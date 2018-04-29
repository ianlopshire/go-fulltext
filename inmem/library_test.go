package inmem_test

import (
	"reflect"
	"testing"

	"github.com/ianlopshire/go-fulltext"
	"github.com/ianlopshire/go-fulltext/inmem"
)

func TestLibrary_List(t *testing.T) {
	for _, tt := range []struct {
		name   string
		docs   []fulltext.Document
		count  int
		offset int
		want   []fulltext.Document
	}{
		{
			name: "no offset",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
			count:  1,
			offset: 0,
			want: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
			},
		},
		{
			name: "with offset case",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
			count:  1,
			offset: 1,
			want: []fulltext.Document{
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
		},
		{
			name: "count > len(docs)",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
			count:  50,
			offset: 0,
			want: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
		},
		{
			name: "count = len(docs)",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
			count:  3,
			offset: 0,
			want: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
		},
		{
			name: "ofsett > len(docs)",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
			count:  50,
			offset: 50,
			want:   nil,
		},
		{
			name: "ofsett = len(docs)",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc3", []byte(`content 3`)),
			},
			count:  3,
			offset: 50,
			want:   nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lib := inmem.NewLibrary(tt.docs)
			have, _ := lib.List(tt.count, tt.offset)
			if !reflect.DeepEqual(tt.want, have) {
				t.Errorf("List(%v, %v) = %+v, want %+v", tt.count, tt.offset, have, tt.want)
			}
		})
	}
}

func TestLibrary_Find(t *testing.T) {
	for _, tt := range []struct {
		name      string
		docs      []fulltext.Document
		token     fulltext.DocumentToken
		want      fulltext.Document
		shouldErr bool
	}{
		{
			name: "general case",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
			token:     "doc1",
			want:      inmem.NewDocument("doc1", []byte(`content 1`)),
			shouldErr: false,
		},
		{
			name: "duplicate token in constructor",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
				inmem.NewDocument("doc2", []byte(`content 2-2`)),
			},
			token:     "doc2",
			want:      inmem.NewDocument("doc2", []byte(`content 2-2`)),
			shouldErr: false,
		},
		{
			name: "err case",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
			token:     "doc3",
			want:      nil,
			shouldErr: true,
		},
		{
			name: "err case empty token",
			docs: []fulltext.Document{
				inmem.NewDocument("doc1", []byte(`content 1`)),
				inmem.NewDocument("doc2", []byte(`content 2`)),
			},
			token:     "",
			want:      nil,
			shouldErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lib := inmem.NewLibrary(tt.docs)
			have, err := lib.Find(tt.token)
			if tt.shouldErr != (err != nil) {
				t.Errorf("Fine(%v) err = %v, want %v", tt.token, err != nil, tt.shouldErr)
			}
			if !tt.shouldErr && !reflect.DeepEqual(have, tt.want) {
				t.Errorf("Fine(%v) = %v, want %v", tt.token, have, tt.want)
			}
		})
	}
}

func TestLibrary_Count(t *testing.T) {
	const want = 2
	have, _ := inmem.NewLibrary([]fulltext.Document{
		inmem.NewDocument("doc1", []byte(`content 1`)),
		inmem.NewDocument("doc2", []byte(`content 2`)),
	}).Count()
	if have != want {
		t.Errorf("count() = %v, want %v", have, want)
	}
}
