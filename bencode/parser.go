package bencode

import (
	"bytes"
	"strconv"
)

func Parse(data []byte) (interface{}, error) {
	if data[0] == 'i' {
		return parseInteger(data)
	}
	return parseString(data)
}

func parseString(data []byte) (string, error) {
	splitted := bytes.SplitN(data, []byte{':'}, 2)
	length, err := strconv.Atoi(string(splitted[0]))
	if err != nil {
		return "", err
	}
	string_data := string(splitted[1])
	return string_data[:length], nil
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
