package bencode

import (
	"strconv"
)

func Parse(data []byte) (interface{}, error) {
	return parseInteger(data)
}

func parseString(data []byte) ([]byte, error) {
	panic("Not implemented")
}

func parseInteger(data []byte) (int, error) {
	return strconv.Atoi(string(data[1 : len(data)-1]))
}

func parseList(data []byte) ([]byte, error) {
	panic("Not implemented")
}

func parseDictionary(data []byte) ([]byte, error) {
	panic("Not implemented")
}
