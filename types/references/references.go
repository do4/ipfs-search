package references

import (
	"context"
	"github.com/ipfs-search/ipfs-search/index"
)

// References represents a list of references
type References []Reference

// Contains returns true of a given reference exists, false when it doesn't
func (references References) Contains(newRef *Reference) bool {
	for _, r := range references {
		if r.ParentHash == newRef.ParentHash {
			return true
		}
	}

	return false
}

// Get first set of references out of a set of indexes, propagating 404's.
// TODO: Use interface Index
func Get(ctx context.Context, indexes []index.Index, id string) (*References, error) {
	// Query All the Indexes
	// TODO: Replace by single MultiGet
	// Ref: https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html

	// TODO: Test cases where:
	// 1. No document is found -> nil, 404 error
	// 2. Document is found, with references -> references, nil
	// 3. Document is found, no references -> nil, nil

	// Anonymous struct for references
	src := struct {
		references References
	}{}

	var err error

	for _, i := range indexes {
		err = i.GetFields(ctx, id, src, "references")

		switch {
		case err == nil:
			// Document found
			return &src.references, nil
		case index.IsNotFound(err):
			// 404, continue to next index
			continue
		default:
			// Unexpected error, return immediately
			return nil, err
		}
	}

	// 404 on all indexes, propagate upwards
	return nil, err
}
