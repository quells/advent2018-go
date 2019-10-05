package file

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	tests := []struct {
		name      string
		original  string
		wantLines []string
		wantErr   bool
	}{
		{"empty", "", []string{}, false},
		{"one line", "abc", []string{"abc"}, false},
		{"two lines", "abc\n123", []string{"abc", "123"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewBufferString(tt.original)
			gotLines, err := ReadLines(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (len(gotLines) == len(tt.wantLines)) && (len(gotLines) == 0) {
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("ReadLines() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
