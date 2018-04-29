package inmem

import (
	"bytes"
	"io"

	"github.com/ianlopshire/go-fulltext"
)

// Document is a completley in-memory document.
type Document struct {
	token fulltext.DocumentToken
	data  []byte
}

// NewDocument creates a new document.
func NewDocument(token fulltext.DocumentToken, data []byte) *Document {
	return &Document{
		token: token,
		data:  data,
	}
}

// Token returns the document token.
func (d *Document) Token() fulltext.DocumentToken {
	return d.token
}

// Reader returns an io.Reader that accesses the document data.
//
// Reader never returns an error.
func (d *Document) Reader() (io.Reader, error) {
	return bytes.NewReader(d.data), nil
}
