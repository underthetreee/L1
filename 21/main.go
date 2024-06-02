package main

import "fmt"

// Interface for read and write data
type StoreReadWriter interface {
	Read() string
	Write(data string)
}

// Represents file storage
type FileStore struct {
}

// Method to store data in file store
func (s *FileStore) Read() string {
	return "read data from file store"
}

// Method to write data from file store
func (s *FileStore) Write(data string) {
	fmt.Println("write data in fs store")
}

// Represents db storage
type DBStore struct {
}

// Method to read data from db store
func (s *DBStore) Read() string {
	return "read data from db store"
}

// Method to store data in db store
func (s *DBStore) Write(data string) {
	fmt.Println("write data in db store")
}

// Adapter for StoreReadWriter interface
type DataAdapter struct {
	rw StoreReadWriter
}

// Creates new adapter
func NewDataAdapter(rw StoreReadWriter) *DataAdapter {
	return &DataAdapter{
		rw: rw,
	}
}

// Method to call read method of StoreReadWriter interface
func (a *DataAdapter) Read() string {
	return a.rw.Read()
}

// Method to call write method of StoreReadWriter interface
func (a *DataAdapter) Write(data string) {
	a.rw.Write(data)
}

func main() {
	// Creates file store
	fs := &FileStore{}
	// Creates adapter with file store because its implements StoreReadWriter interface
	adapter := NewDataAdapter(fs)
	// Read data using adapter
	fmt.Println("read data", adapter.Read())
	// Write data using adapter
	adapter.Write("new data")

	// Same with db store
	db := &DBStore{}
	adapter = NewDataAdapter(db)
	fmt.Println("read data", adapter.Read())
	adapter.Write("new data")
}
