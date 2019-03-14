package gocb

import (
	"testing"
	"time"

	"github.com/couchbase/gocb/cbft"

	"github.com/couchbaselabs/gocbcore8alpha"
)

func TestSearchQueryServiceNotFound(t *testing.T) {
	doHTTP := func(req *gocbcore.HttpRequest) (*gocbcore.HttpResponse, error) {
		return nil, gocbcore.ErrNoFtsService
	}

	provider := &mockHTTPProvider{
		doFn: doHTTP,
	}

	q := SearchQuery{
		Name:  "test",
		Query: cbft.NewMatchQuery("test"),
	}
	timeout := 60 * time.Second

	cluster := testGetClusterForHTTP(provider, timeout, 0, 0)

	res, err := cluster.SearchQuery(q, nil)
	if err == nil {
		t.Fatal("Expected query to return error")
	}

	if res != nil {
		t.Fatalf("Expected result to be nil but was %v", res)
	}

	if !IsServiceNotFoundError(err) {
		t.Fatalf("Expected error to be ServiceNotFoundError but was %v", err)
	}
}
