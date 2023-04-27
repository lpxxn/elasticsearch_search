package open_search

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init([]string{"https://vpc-cafe-cache-yax5i6n5md2r2blnct5ypdiyja.cn-northwest-1.es.amazonaws.com.cn"}, "", "")
	os.Exit(m.Run())
}
