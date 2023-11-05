package Token_Handler

import (
	"testing"
)

func TestGenerate_Signature(t *testing.T) {
	tests := []struct {
		header    string
		payload   string
		algorithm string
		want      string
	}{
		{
			header:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			payload:   "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9",
			algorithm: "HS256",
			want:      "gMK3RcrUXBFWCavXxSOOA3sHy_Q3cWyGAbQ4zvuWJF8",
		},
		{
			header:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			payload:   "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9",
			algorithm: "HS384",
			want:      "gMK3RcrUXBFWCavXxSOOA3sHy_Q3cWyGAbQ4zvuWJF8",
		},
		{
			header:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			payload:   "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9",
			algorithm: "HS512",
			want:      "JUu57vre_2-fIwQyEIfW_aXiNYqZtM3J2pZFNvn-9IhoLo7hsOhK4o4CmUroAjuhscvvd5-L0kGHdRTLG7rb2Q",
		},
	}

	for _, test := range tests {
		t.Run(test.algorithm, func(t *testing.T) {
			got, _ := generate_signature(test.header, test.payload, test.algorithm)
			if got != test.want {
				t.Errorf("generate_signature() = %v, want %v", got, test.want)
			}
		})
	}
}
