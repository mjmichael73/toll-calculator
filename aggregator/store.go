package main

import "github.com/mjmichael73/toll-calculator/types"

type MemoryStore struct {}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}
func (m *MemoryStore) Insert(d types.Distance) error {
	return nil
}