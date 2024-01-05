package test

import (
	"testing"

	"github.com/Courtcircuits/BitTorrent.go/bencode"
	"github.com/stretchr/testify/assert"
)

func TestParsingIntegers(t *testing.T) {
	bencode_string := "i123e"
	actual_value := 123

	parsed_value, err := bencode.Parse([]byte(bencode_string))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, parsed_value, actual_value)
}
