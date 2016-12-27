package main

// Retriever - Bridge for doing http requests
type Retriever interface {
	retrieve(url string) []byte
}

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}
