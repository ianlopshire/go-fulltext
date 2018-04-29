package inmem

import (
	"github.com/ianlopshire/go-fulltext"
)

// Library defines a set of documents.
type Library struct {
	docs []fulltext.Document
	m    map[fulltext.DocumentToken]int
}

// NewLibrary creates a new library containing the given documents.
//
// If two docs have the same token, the document with the greater index
// will be used.
func NewLibrary(docs []fulltext.Document) *Library {
	l := &Library{
		docs: make([]fulltext.Document, 0, len(docs)),
		m:    map[fulltext.DocumentToken]int{},
	}
	for i, doc := range docs {
		l.m[doc.Token()] = i
		l.docs = append(l.docs, doc)
	}
	return l
}

// Find returns the document with the given token.
//
// If a document is not found, Find returns fulltext.ErrDocNotFound.
func (l *Library) Find(token fulltext.DocumentToken) (fulltext.Document, error) {
	i, ok := l.m[token]
	if !ok {
		return nil, fulltext.ErrDocNotFound
	}
	return l.docs[i], nil
}

// List returns a list of documents.
//
// The count param controlls the maximum number of documents returned.
// The offset param controlls the index where the returned list starts.
//
// List will never return an error.
func (l *Library) List(count, offset int) ([]fulltext.Document, error) {
	if count == 0 || offset > len(l.docs) {
		return nil, nil
	}

	if offset+count > len(l.docs) {
		return l.docs[offset:], nil
	}

	return l.docs[offset : offset+count], nil
}

// Count returns the number of documents in the library.
//
// Count will never return an error.
func (l *Library) Count() (int, error) {
	return len(l.docs), nil
}
