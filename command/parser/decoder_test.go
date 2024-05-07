package parser

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name       string
		data       []byte
		wantCmd    string
		wantArgs   []string
		wantErr    bool
	}{
		{
			name:       "simple set command",
			data:       []byte("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"),
			wantCmd:    "SET",
			wantArgs:   []string{"key", "value"},
			wantErr:    false,
		},
		{
			name:       "invalid format missing parts",
			data:       []byte("*2\r\n$3\r\nGET\r\n"),
			wantCmd:    "",
			wantArgs:   nil,
			wantErr:    true,
		},
		{
			name:       "empty input",
			data:       []byte(""),
			wantCmd:    "",
			wantArgs:   nil,
			wantErr:    true,
		},
		{
			name:       "incorrect header",
			data:       []byte("3\r\n$3\r\nGET\r\n$3\r\nkey\r\n"),
			wantCmd:    "",
			wantArgs:   nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, args, err := Decode(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if cmd != tt.wantCmd {
					t.Errorf("Decode() gotCmd = %v, want %v", cmd, tt.wantCmd)
				}
				if !reflect.DeepEqual(args, tt.wantArgs) {
					t.Errorf("Decode() gotArgs = %v, want %v", args, tt.wantArgs)
				}
			}
		})
	}
}
