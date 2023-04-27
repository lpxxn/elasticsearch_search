package open_search

import "testing"

func TestSearchAfter1(t *testing.T) {
	body := struct {
		Name            string `json:"name"`
		NextSnowflakeID int32  `json:"snowflake_id"`
	}{}
	body.Name = ""
}
