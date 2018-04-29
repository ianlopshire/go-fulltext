package inmem_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/ianlopshire/go-fulltext/inmem"
)

func TestDocument_Token(t *testing.T) {
	const want = "token1"
	have := inmem.NewDocument(want, nil).Token()
	if want != have {
		t.Errorf("Token() = %v, want %v", have, want)
	}
}

func TestDocument_Reader(t *testing.T) {
	want := []byte(`content`)
	r, _ := inmem.NewDocument("", want).Reader()
	have, _ := ioutil.ReadAll(r)
	if !bytes.Equal(want, have) {
		t.Errorf("Reader() = %s, want %s", have, want)
	}
}
