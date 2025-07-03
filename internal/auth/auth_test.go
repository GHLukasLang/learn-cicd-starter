package auth

import (
	//"reflect"
	"net/http"
	"testing"
	//"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr string
	}{
		"empty header": {input: http.Header{}, want: "", wantErr: "no authorization header included"},
		"len<2":        {input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", wantErr: "malformed authorization header"},
		"no Auth key":  {input: http.Header{"Authorization": []string{"ABC", "EFG", "HIJ"}}, want: "", wantErr: "malformed authorization header"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
			if err == nil || err.Error() != tc.wantErr {
				t.Errorf("got error %v, want %q", err, tc.wantErr)
			}

		})
	}
}
