package klaviyo_test

import (
	"context"
	"testing"

	"github.com/bluettipower/klaviyo-go"
)

func TestListBrowse(t *testing.T) {
	client := setupClient()
	t.Log(client.APIKey)
	resp, err := client.List.Browse(context.TODO(), klaviyo.BrowseListRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, v := range resp.Data {
		t.Logf("name: %s, id: %s,", *v.Attributes.Name, *v.ID)
	}
}
