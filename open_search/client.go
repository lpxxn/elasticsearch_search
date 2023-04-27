package open_search

import "github.com/opensearch-project/opensearch-go/v2"

var Client *opensearch.Client

func Init(addr []string, userName, pwd string) error {
	var err error
	Client, err = opensearch.NewClient(opensearch.Config{
		Addresses: addr,
		Username:  userName,
		Password:  pwd,
	})
	return err
}
