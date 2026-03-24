package nyt

import (
	"fmt"
	"net/url"
	"slices"
)

const TOP_STORIES_BASE = API_BASE + "svc/topstories/v2/"

type topStoriesResponse struct {
	Results []Article `json:"results"`
}

func (c *Client) TopStories(section string) ([]Article, error) {
	if !slices.Contains(ValidTopStoriesSections, section) {
		return nil, ErrInvalidSection
	}
	target, err := url.JoinPath(TOP_STORIES_BASE, fmt.Sprintf("%s.json", section))
	if err != nil {
		return nil, err
	}
	var result topStoriesResponse
	if err := c.Get(target, &result); err != nil {
		return nil, err
	}
	return result.Results, nil
}
