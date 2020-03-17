package metacpan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAutocompleteSuggest(t *testing.T) {
	hits, err := SearchAutocompleteSuggest("HTML")
	assert.NoError(t, err)
	assert.Greater(t, len(hits), 5)
	assert.Contains(t, hits[0].URL(), "https://metacpan.org/pod/")
	assert.NotEmpty(t, hits[0].Date)
}
