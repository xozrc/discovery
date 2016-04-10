package types

import (
	"encoding/json"
)

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

type entry struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (e *entry) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *entry) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

func CreateEntry() Entry {
	return &entry{}
}

var (
	EntryFactoryInstance EntryFactory
)

func init() {
	EntryFactoryInstance = EntryFactoryFunc(CreateEntry)
}
