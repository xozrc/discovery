package discovery_test

import (
	"fmt"
	"testing"
	"time"
)
import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/discovery/discovery"
	"github.com/xozrc/discovery/discovery/mock_store"
	"github.com/xozrc/discovery/types"
)

var print = fmt.Print

var (
	prefix      = "xozrc_test"
	serviceName = "test"
	ttl         = time.Duration(time.Second * 10)
	heartbeat   = time.Duration(time.Second * 10)
)

func TestWatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := mock_store.NewMockStore(ctrl)
	bke := discovery.NewBackend(ms, types.EntryFactoryInstance, prefix, serviceName, heartbeat)
	testWatch(t, bke)
	testWatchError(t, bke)
}

func testWatchError(t *testing.T, bke discovery.DiscoveryBackend) {
	fmt.Println("test watch error")
	bke.Watch()
	_, _, err := bke.Watch()
	defer bke.Unwatch()
	assert.Error(t, err, "should watch error")

}

func testWatch(t *testing.T, bke discovery.DiscoveryBackend) {
	fmt.Println("test watch correct")
	_, _, err := bke.Watch()
	defer bke.Unwatch()
	assert.NoError(t, err, "should watch correct")
}

func TestUnwatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := mock_store.NewMockStore(ctrl)
	bke := discovery.NewBackend(ms, types.EntryFactoryInstance, prefix, serviceName, heartbeat)
	testUnwatch(t, bke)
	testUnwatchError(t, bke)
}

func testUnwatch(t *testing.T, bke discovery.DiscoveryBackend) {
	fmt.Println("test unwatch correct")
	_, _, err := bke.Watch()
	assert.NoError(t, err, "should watch correct")
	err = bke.Unwatch()
	assert.NoError(t, err, "should unwatch correct")
}

func testUnwatchError(t *testing.T, bke discovery.DiscoveryBackend) {
	fmt.Println("test unwatch error")
	_, _, err := bke.Watch()
	assert.NoError(t, err, "should watch correct")
	err = bke.Unwatch()
	assert.NoError(t, err, "should unwatch correct")
	err = bke.Unwatch()
	assert.Error(t, err, "should unwatch error")
}

func TestRegister(t *testing.T) {

}

func TestUnregister(t *testing.T) {

}
