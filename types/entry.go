package types

import (
	"encoding/json"
)

var (
	EntryFactoryInstance EntryFactory
    HostEntryFactoryInstance EntryFactory
)

func init() {
	EntryFactoryInstance = EntryFactoryFunc(createEntry)
    HostEntryFactoryInstance = EntryFactoryFunc(createHostEntry)
}

type Entries []Entry

type Entry interface {
	Marshaller
	Unmarshaller
}

type Marshaller interface {
	Marshal() ([]byte, error)
}

type Unmarshaller interface {
	Unmarshal(data []byte) error
}

type EntryFactory interface {
	CreateEntry() Entry
}

type EntryFactoryFunc func() Entry

func (eff EntryFactoryFunc) CreateEntry() Entry {
	return eff()
}

type entry string

func (e *entry) Marshal() ([]byte, error) {
	return []byte(*e),nil
}

func (e *entry) Unmarshal(data []byte) error {
	*e = entry(string(data))
    return nil
}

func createEntry() Entry {
    te := entry("")
	return &te
}


type HostEntry struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (e *HostEntry) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *HostEntry) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

func createHostEntry() Entry{
       return &HostEntry{}
}
