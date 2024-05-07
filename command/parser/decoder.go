package parser

import (
	"bytes"
	"errors"
	"strconv"
)

// Decode by Redis Serialization Protocol datas
// return: command, []argument, error
func Decode(data []byte) (string, []string, error) {

	if len(data) == 0 {
		return "", nil, errors.New("empty data given")
	}

	// Split the data
	parts := bytes.Split(data, []byte("\r\n"))
	if len(parts) < 3 {
		return "", nil, errors.New("invalid data format")
	}

	// The first part should start with '*' and tell us how many element to expect
	if parts[0][0] != '*' {
		return "", nil, errors.New("unexpected array format")
	}

	numElements, err := strconv.Atoi(string(parts[0][1:]))
	if err != nil {
		return "", nil, errors.New("invalid number of elements")
	}

	if numElements < 1 {
		return "", nil, errors.New("number of elements must be positive")
	}

	// Make a Slice to hold commmands and arguments
	results := make([]string, 0, numElements)
	index := 1

	for len(results) < numElements {
		if index+2 >= len(parts) {
			return "", nil, errors.New("data format error: not enough data")
		}

		elements := parts[index+1]
		results = append(results, string(elements))
		index += 2
	}

	if len(results) == 0 {
		return "", nil, errors.New("no command found")
	}

	command := results[0]
	arguments := results[1:]

	return command, arguments, nil
}
