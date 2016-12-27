package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TestRetriever struct{}

func (retriever TestRetriever) retrieve(URL string) []byte {
	strippedURL := strings.Replace(strings.Replace(URL[21:], "?", "_", -1), "&", "_", -1)
	blah, err := os.Open("testdata/" + strippedURL)
	if err != nil {
		log.Printf("Error opening test file %v", err)
	}
	body, _ := ioutil.ReadAll(blah)
	return body
}

// Gets a test server that'll pull from local files rather than reading out
func getTestServer() Server {
	s := Server{}
	s.retr = TestRetriever{}
	s.marshaller = prodUnmarshaller{}
	return s
}
