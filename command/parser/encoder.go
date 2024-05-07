package parser

import "fmt"

func Encode(data string) []byte {

	// if empty return RESP empty bulk string with the length of -1
	if len(data) == 0 {
		return []byte("$-1\r\n")	
	}

	// Check if response is an simples string or an error
	firstChar := data[0]
	switch firstChar {
	case '+',':': 
		return []byte(data + "\r\n")
	case '-':
		return []byte(data + "\r\n")
	default: // Assume everything else is a bulk string I guess 
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(data), data))
	}
}
