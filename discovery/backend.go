package discovery

import (
	"fmt"
	"log"
	"path"
	"sync"
	"time"
)
import (
	"github.com/docker/docker/pkg/tlsconfig"
	"github.com/docker/libkv"
	store "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"

	types "github.com/xozrc/discovery/types"
)

func init() {
	etcd.Register()
}

// discovery backend
type DiscoveryBackend interface {
	Register(entry types.Entry, ttl time.Duration) error
	Deregister(entry types.Entry) error
	Watch() (chan types.Entries, chan error, error)
	Unwatch() error
}

func NewBackend(s store.Store, ef types.EntryFactory, p string, serv string, hb time.Duration) DiscoveryBackend {
	edb := &discoveryBackend{
		store:        s,
		entryFactory: ef,
		prefix:       p,
		service:      serv,
		heartbeat:    hb,
	}
	return edb
}

type discoveryBackend struct {
	//kv store base on docker/libkv/store
	store store.Store
	//entry factory
	entryFactory types.EntryFactory
	//custom prefix
	prefix string
	//service name
	service string
	//heartbeat
	heartbeat time.Duration
	//lock for watch and unwatch
	mu sync.RWMutex
	//path
	path string
	//watch stop channel
	stopCh chan struct{}
}

//register
func (edb *discoveryBackend) Register(entry types.Entry, ttl time.Duration) (err error) {
	opts := &store.WriteOptions{TTL: ttl}
	tb, err := entry.Marshal()
	if err != nil {
		return
	}
	key := path.Join(edb.Path(), string(tb))
	err = edb.store.Put(key, tb, opts)
	return
}

//deregister
func (edb *discoveryBackend) Deregister(entry types.Entry) (err error) {
	tb, err := entry.Marshal()
	if err != nil {
		return
	}
	key := path.Join(edb.Path(), string(tb))
	err = edb.store.Delete(key)
	return
}

//service name path
func (edb *discoveryBackend) Path() string {
	if edb.path == "" {
		edb.path = path.Join(edb.prefix, edb.service)
	}
	return edb.path
}

//watch
func (edb *discoveryBackend) Watch() (entriesCh chan types.Entries, errCh chan error, err error) {

	edb.mu.Lock()
	defer edb.mu.Unlock()

	//already watch
	if edb.isWatch() {
		err = fmt.Errorf("discovery backend already watch")
		return
	}

	//setup stop watch channel
	edb.stopCh = make(chan struct{})
	//setup entries channel
	entriesCh = make(chan types.Entries)
	//setup error channel
	errCh = make(chan error)

	//go coroutine for async watch
	go func() {
		for {

			select {
			//stop watch
			case <-edb.stopCh:
				break
			default:
			}

			var watchCh <-chan []*store.KVPair
			//if exists
			exists, err := edb.store.Exists(edb.Path())
			if err != nil {
				errCh <- err
				goto Err
			}

			//no exist
			if !exists {
				if err := edb.store.Put(edb.Path(), []byte(""), &store.WriteOptions{IsDir: true}); err != nil {
					errCh <- err
					goto Err
				}
			}

			//watch tree
			watchCh, err = edb.store.WatchTree(edb.Path(), edb.stopCh)

			if err != nil {
				errCh <- err
				goto Err
			}

			//wait watch channel
			for {
				select {
				case pairs := <-watchCh:
					{

						entries := make([]types.Entry, len(pairs))

						for i, kv := range pairs {

							te := edb.entryFactory.CreateEntry()
							err = te.Unmarshal([]byte(kv.Value))

							if err != nil {
								errCh <- err
								continue
							}
							entries[i] = te
						}

						entriesCh <- types.Entries(entries)
					}
				case <-edb.stopCh:
					break
				}
			}
		Err:
			time.Sleep(edb.heartbeat)
		}
	}()

	return
}

//stop watch
func (edb *discoveryBackend) Unwatch() error {
	edb.mu.Lock()
	defer edb.mu.Unlock()

	if !edb.isWatch() {
		return fmt.Errorf("discovery backend not watch yet")
	}

	close(edb.stopCh)
	edb.stopCh = nil
	return nil
}

//if watch
func (edb *discoveryBackend) isWatch() bool {
	return edb.stopCh != nil
}

func NewEtcdStore(cafile, certfile, keyfile string, addrs []string) (store.Store, error) {

	var config *store.Config
	if cafile != "" && certfile != "" && keyfile != "" {
		log.Println("Initializing discovery with TLS")
		tlsConfig, err := tlsconfig.Client(tlsconfig.Options{
			CAFile:   cafile,
			CertFile: certfile,
			KeyFile:  keyfile,
		})
		if err != nil {
			return nil, err
		}
		config = &store.Config{
			// Set ClientTLS to trigger https (bug in libkv/etcd)
			ClientTLS: &store.ClientTLSConfig{
				CACertFile: cafile,
				CertFile:   certfile,
				KeyFile:    keyfile,
			},
			// The actual TLS config that will be used
			TLS: tlsConfig,
		}
	} else {
		log.Println("Initializing discovery without TLS")

	}

	s, err := libkv.NewStore(store.ETCD, addrs, config)
	return s, err
}
