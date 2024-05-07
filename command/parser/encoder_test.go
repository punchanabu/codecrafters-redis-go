package parser

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []byte
	}{
		{
			name: "Simple string",
			data: "+OK",
			want: []byte("+OK\r\n"),
		},
		{
			name: "Error message",
			data: "-Error",
			want: []byte("-Error\r\n"),
		},
		{
			name: "Bulk string",
			data: "hello world",
			want: []byte("$11\r\nhello world\r\n"),
		},
		{
			name: "Empty string",
			data: "",
			want: []byte("$-1\r\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode(%q) = %q, want %q", tt.data, got, tt.want)
			}
		})
	}
}
