package test

import (
	"testing"

	"github.com/Courtcircuits/BitTorrent.go/bencode"
	"github.com/stretchr/testify/assert"
)

func TestParsingIntegers(t *testing.T) {
	bencode_strings := []string{"i123e", "i-123e"}
	actual_values := []int{123, -123}

	for i, bencode_string := range bencode_strings {
		parsed_value, err := bencode.Parse([]byte(bencode_string))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, parsed_value, actual_values[i])
	}
}
