package metacpan

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// The AutocompleteFields defines fields in search result of autocomplete API.
type AutocompleteFields struct {
	Author        string `json:"author"`
	Release       string `json:"release"`
	Documentation string `json:"documentation"`
	Distribution  string `json:"distribution"`
}

// The AutocompleteSuggestFields defines fields in search result of autocomplete suggest API.
type AutocompleteSuggestFields struct {
	Author       string `json:"author"`
	Release      string `json:"release"`
	Date         string `json:"date"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
}

// The AutocompleteHit defines one of search result of autocomplete API.
type AutocompleteHit struct {
	Fields AutocompleteFields `json:"fields"`
}

// URL returns url on metacpan.
func (a AutocompleteHit) URL() string {
	return fmt.Sprintf("https://%s/pod/%s", htmlHost, a.Fields.Documentation)
}

// URL returns url on metacpan.
func (a *AutocompleteSuggestFields) URL() string {
	return fmt.Sprintf("https://%s/pod/%s", htmlHost, a.Name)
}

// The AutocompleteHits defines hits in search result of autocomplete API.
type AutocompleteHits struct {
	Hits []AutocompleteHit `json:"hits"`
}

// The AutocompleteResult defines search result of autocomplete API.
type AutocompleteResult struct {
	Hits AutocompleteHits `json:"hits"`
}

type AutocompleteSuggestResult struct {
	Suggestions []AutocompleteSuggestFields `json:"suggestions"`
}

// SearchAutocomplete search autocomplete by query and returns hits.
func SearchAutocomplete(q string) ([]AutocompleteHit, error) {
	body, err := request(fmt.Sprintf("%s?q=%s", APISearchAutocomplete, url.QueryEscape(q)))

	if err != nil {
		return nil, err
	}

	var result AutocompleteResult
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result.Hits.Hits, nil
}

// SearchAutocomplete search autocomplete by query and returns hits.
func SearchAutocompleteSuggest(q string) ([]AutocompleteSuggestFields, error) {
	body, err := request(fmt.Sprintf(
		"%s?q=%s",
		APISearchAutocompleteSuggest,
		url.QueryEscape(q),
	))

	if err != nil {
		return nil, err
	}

	var result AutocompleteSuggestResult
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}
