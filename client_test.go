package nyt

import (
	"os"
	"testing"
)

func newTestClient(t *testing.T) *Client {
	t.Helper()
	apiKey := os.Getenv("NYT_API_KEY")
	if apiKey == "" {
		t.Skip("NYT_API_KEY not set")
	}
	return NewClient(apiKey)
}

func TestTopStories_Home(t *testing.T) {
	c := newTestClient(t)
	articles, err := c.TopStories("home")
	if err != nil {
		t.Fatalf("TopStories(home): %v", err)
	}
	if len(articles) == 0 {
		t.Fatal("TopStories(home): expected articles, got none")
	}
}

func TestTopStories_US(t *testing.T) {
	c := newTestClient(t)
	articles, err := c.TopStories("us")
	if err != nil {
		t.Fatalf("TopStories(us): %v", err)
	}
	if len(articles) == 0 {
		t.Fatal("TopStories(us): expected articles, got none")
	}
}


