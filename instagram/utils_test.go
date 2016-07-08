package instagram_test

import (
	"testing"

	"github.com/xboston/go-instagram/instagram"
)

func TestGetProfileByURL(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err bool
	}{
		{in: "", err: true},
		{in: "/", err: true},
		{in: "//", err: true},

		{in: "http://no-instagram.com", err: true},
		{in: "http://instagram.com", err: true},
		{in: "HTTP://x.instagram.com", err: true},
		{in: "HTTP://x.instagram.com/username", err: true},
		{in: "HTTP://instagram.com/username", out: "username"},
		{in: "HTTPs://instagram.com/username", out: "username"},
		{in: "instagram.com/username", out: "username"},
		{in: "www.instagram.com/username", out: "username"},
		{in: "//instagram.com/username", out: "username"},
		{in: "https://www.instagram.com/username/?test=test", out: "username"},
	}

	for _, tt := range tests {
		profile, err := instagram.GetProfileByURL(tt.in)
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
		if *profile != tt.out {
			t.Errorf(`"%s": got "%s", want "%v"`, tt.in, *profile, tt.out)
		}
	}
}
