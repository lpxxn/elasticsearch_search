package goe_demo

import (
	"os"
	"testing"

	"github.com/lpxxn/elasticsearch_search/es_lib/entity"
)
var ES7ClientT entity.Es7ClientType
var myTest2Index = "mytest_geo1"

func TestLocations1(t *testing.T) {
	ES7ClientT.
}

func TestMain(m *testing.M) {
	ES7ClientT = entity.NewEsClient()

	os.Exit(m.Run())
}
