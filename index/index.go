package index

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

// Index wraps an Elasticsearch index to store documents
type Index struct {
	Client *elastic.Client
	Name   string
}

// Upsert a document properties and given id
func (i *Index) Upsert(ctx context.Context, id string, properties map[string]interface{}) error {
	_, err := i.Client.Update().
		Index(i.Name).
		Id(id).
		Doc(properties).
		DocAsUpsert(true).
		Do(ctx)

	if err != nil {
		// Handle error
		return err
	}

	return nil

}

// GetFields retreives `fields` from document with `id` from the index.
func (i *Index) GetFields(ctx context.Context, id string, dst interface{}, fields ...string) error {
	fsc := elastic.NewFetchSourceContext(true)
	fsc.Include(fields...)

	result, err := i.Client.
		Get().
		Index(i.Name).
		FetchSourceContext(fsc).
		Id(id).
		Do(ctx)

	if err != nil {
		return err
	}

	// Decode resulting field json into `dst`
	return json.Unmarshal(*result.Source, dst)
}

// IsNotFound return true if the error is due to the document not being found.
func IsNotFound(err error) bool {
	return elastic.IsNotFound(err)
}
