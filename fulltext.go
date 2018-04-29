// Package fulltext aims to provide a set of interfaces and common implementations for performing
// full-text search across a set of documents.
package fulltext

import (
	"errors"
	"io"
)

// DocumentToken identifies a document.
//
// A DocumentToken should be unique within a library.
type DocumentToken string

// Document represents a single document.
type Document interface {
	// Token returns the documents token.
	Token() DocumentToken

	// Reader returns an io.Reader that accesses the document data or an error.
	Reader() (io.Reader, error)
}

var (
	ErrDocNotFound = errors.New("document not found")
)

// Library defines a set of documents.
type Library interface {
	// Find returns the document with the given token or an error.
	//
	// If a document is not found, Find should return ErrDocNotFound.
	Find(DocumentToken) (Document, error)

	// List returns a list of documents or an error.
	//
	// The count param controlls the maximum number of documents returned.
	// The offset param controlls the index where the returned list starts.
	//
	// Calling List should return consistent results given a specific count
	// and offset.
	List(count, offset int) ([]Document, error)

	// Count should return the number of documents in a library or an error.
	Count() (int, error)
}
