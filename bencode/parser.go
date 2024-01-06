package bencode

import (
	"fmt"
	"strconv"
)

func Parse(data []byte) (interface{}, error) {
	pointer := 0
	return parseAux(data, &pointer)
}

func parseAux(data []byte, pointer *int) (interface{}, error) {
	if data[*pointer] == 'l' {
		next(data, pointer)
		return parseLists(data, pointer)
	}
	if data[*pointer] == 'i' {
		next(data, pointer)
		return parseInteger(data, pointer)
	}
	return parseStrings(data, pointer)
}

func parseInteger(data []byte, pointer *int) (int, error) {
	stringed_integer := ""
	for data[*pointer] != 'e' {
		stringed_integer += string(data[*pointer])
		next(data, pointer)
	}
	value, err := strconv.Atoi(stringed_integer)
	next(data, pointer) // skip 'e'
	return value, err
}

func parseStrings(data []byte, pointer *int) (string, error) {
	stringed_length := ""
	for data[*pointer] != ':' {
		stringed_length += string(data[*pointer])
		next(data, pointer)
	}
	length, err := strconv.Atoi(stringed_length)
	if err != nil {
		return "", err
	}
	next(data, pointer)
	value := string(data[*pointer : *pointer+length])
	for i := 0; i < length	; i++ {
		next(data, pointer)
	}
	return value, nil
}

func parseLists(data []byte, pointer *int) ([]interface{}, error) {
	var list []interface{}
	for data[*pointer] != 'e' {
		value, err := parseAux(data, pointer)
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}
	return list, nil
}

// for debugging purposes
func next(data []byte, pointer *int) {
	fmt.Println(string(data[*pointer]))
	*pointer++
}
