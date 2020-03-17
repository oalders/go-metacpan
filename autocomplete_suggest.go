package metacpan

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// AutocompleteSuggestion describes a single search result of the autocomplete suggest API.
type AutocompleteSuggestion struct {
	Author       string `json:"author"`
	Release      string `json:"release"`
	Date         string `json:"date"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
}

type AutocompleteSuggestResults struct {
	Suggestions []AutocompleteSuggestion `json:"suggestions"`
}

// URL returns url on metacpan.
func (s *AutocompleteSuggestion) URL() string {
	return fmt.Sprintf("https://%s/pod/%s", htmlHost, s.Name)
}

// SearchAutocomplete search autocomplete by query and returns hits.
func SearchAutocompleteSuggest(q string) ([]AutocompleteSuggestion, error) {
	body, err := request(fmt.Sprintf(
		"%s?q=%s",
		APISearchAutocompleteSuggest,
		url.QueryEscape(q),
	))

	if err != nil {
		return nil, err
	}

	var result AutocompleteSuggestResults
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}
