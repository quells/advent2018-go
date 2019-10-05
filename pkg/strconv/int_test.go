package strconv

import "testing"

func TestParseSignedInt(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantI   int
		wantErr bool
	}{
		{"empty", "", 0, true},
		{"no sign", "123", 123, false},
		{"plus", "+123", 123, false},
		{"minus", "-123", -123, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := ParseSignedInt(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSignedInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("ParseSignedInt() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
