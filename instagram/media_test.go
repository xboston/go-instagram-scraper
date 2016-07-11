package instagram_test

import (
	"testing"

	"github.com/xboston/go-instagram/instagram"
)

func TestMediaExist(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err bool
	}{
		{in: "", err: true},

		{in: "BFmQqm6Bxdk", err: false},
		{in: "BFmQqm6Bx11", err: true},
	}

	client := instagram.NewClient(nil)

	for _, tt := range tests {

		_, err := client.Media.Exist(tt.in)
		if err != nil {
			if !tt.err {
				t.Errorf(`"%s": unexpected error "%v"`, tt.in, err)
			}
			continue
		}
		if tt.err && err == nil {
			t.Errorf(`"%s": expected error`, tt.in)
			continue
		}
	}
}
