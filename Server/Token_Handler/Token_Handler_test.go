package Token_Handler

import (
	struct_def "Maria_Demo/Structs"
	"Maria_Demo/convert_handler"
	"testing"
)

func TestGenerate_Signature(t *testing.T) {
	tests := []struct {
		testName  string
		header    struct_def.Jwt_Header
		payload   struct_def.Jwt_Payload
		algorithm string
		want      string
		willFail  bool
	}{
		{
			testName: "Valid Header w/ HS384 Algorithm",
			header: struct_def.Jwt_Header{
				Algorithm: "HS384",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			algorithm: "Invalid",
			want:      "EavvOXgGprci0P2A02Ow8iBbLpZdiAiAhmG9B72ZBPQYSKh4HvvlQ3dbY97329_j",
			willFail:  true,
		},
		{
			testName: "Valid Header w/ HS256 Algorithm",
			header: struct_def.Jwt_Header{
				Algorithm: "HS256",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			algorithm: "HS256",
			want:      "CDY0cWzz9D1zvABhYgLRyUJjTGzJyzCzERMWcFoQOEs",
			willFail:  false,
		},
		{
			testName: "Invalid Payload w/ HS384 Algorithm",
			header: struct_def.Jwt_Header{
				Algorithm: "HS384",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			algorithm: "HS384",
			want:      "6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.algorithm, func(t *testing.T) {
			t.Logf("Test: %v", test.testName)
			// Ignoring the error value because convert_handler is tested separately.
			header, _ := convert_handler.Struct_ToBytes(test.header)
			payload, _ := convert_handler.Struct_ToBytes(test.payload)
			got, err := generate_signature(header, payload, test.algorithm)

			if err != nil && test.willFail {
				return
			}
			if err != nil && !test.willFail {
				t.Errorf("generate_signature() error = %v, wantErr %v", err, test.willFail)
				t.Fail()
				return
			}
			if got != test.want {
				if test.willFail {
					return
				}
				t.Errorf("generate_signature() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSign_Token(t *testing.T) {
	test := []struct {
		testName string
		header   struct_def.Jwt_Header
		payload  struct_def.Jwt_Payload
		want     string
		willFail bool
	}{
		{
			testName: "Valid JWT",
			header: struct_def.Jwt_Header{
				Algorithm: "HS256",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			want:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.CDY0cWzz9D1zvABhYgLRyUJjTGzJyzCzERMWcFoQOEs",
			willFail: false,
		},
		{
			testName: "Invalid JWT",
			header: struct_def.Jwt_Header{
				Algorithm: "HS256",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			want:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail: true,
		},
		{
			testName: "Empty Header",
			header:   struct_def.Jwt_Header{},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			want:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail: true,
		},
		{
			testName: "Tampered Payload",
			header: struct_def.Jwt_Header{
				Algorithm: "HS256",
			},
			payload: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user", "superuser"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			want:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.CDY0cWzz9D1zvABhYgLRyUJjTGzJyzCzERMWcFoQOEs",
			willFail: true,
		},
	}

	for _, test := range test {
		t.Run(test.testName, func(t *testing.T) {
			t.Logf("Test: %v", test.testName)
			got, err := Sign_Token(test.header, test.payload)
			if err != nil && test.willFail {
				return
			}
			if err != nil && !test.willFail {
				t.Errorf("Sign_Token() error = %v, wantErr %v", err, test.willFail)
				t.Fail()
				return
			}
			if got != test.want {
				if test.willFail {
					return
				}
				t.Errorf("Sign_Token() = %v, want %v", got, test.want)
			}
			if test.willFail {
				t.Errorf("Sign_Token(): %v should have failed but didn't", test.testName)
				t.Fail()
			}
		})
	}
}

func TestVerify_Token(t *testing.T) {
	tests := []struct {
		testName string
		token    string
		want     bool
		willFail bool
	}{
		{
			testName: "Valid JWT",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			want:     true,
			willFail: false,
		},
		{
			testName: "Invalid JWT",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iLCJzdXBlcmFkbWluIl0sImV4cCI6MTY5OTIyMDQ0NTQ5MzY2NTMwMCwiYXNzaWduZWQiOjE2OTkyMTY4NDU0OTM2NjUzMDB9.PQdoNHYHtq2v3jz0iKVOZiDJQd9qpTsJrry78Y1M_Ns",
			want:     false,
			willFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			output, err := Verify_Token(test.token)

			if err != nil && test.willFail {
				return
			} else if err != nil {
				t.Errorf("Verify_Token() error = %v, wantErr %v", err, test.willFail)
				t.Fail()
				return
			}

			if output != test.want {
				if test.willFail {
					return
				}
				t.Errorf("Verify_Token() = %v, want %v", output, test.want)
			}
			if test.willFail {
				t.Errorf("Verify_Token(): %v should have failed but didn't", test.testName)
				t.Fail()
			}
		})
	}
}

/*
	No point in testing Generate_Token() because it's just a wrapper for Sign_Token()
*/

func TestRead_Token_Data(t *testing.T) {
	tests := []struct {
		testName string
		token    string
		willFail bool
	}{
		{
			testName: "Valid JWT",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail: false,
		},
		{
			testName: "No Payload",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail: true,
		},
		{
			testName: "Too Many Sections",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0",
			willFail: true,
		},
		{
			testName: "Not Valid Base64",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9.eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0=.6yzJKYydpdes8WHqYADx5eIB3INhguNk1X6OEyNip-8",
			willFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			output, err := Read_Token_Data(test.token)

			if err != nil && test.willFail {
				return
			} else if err != nil {
				t.Errorf("Read_Token_Data() error = %v, wantErr %v", err, test.willFail)
				t.Fail()
				return
			}

			if output == nil {
				t.Errorf("Read_Token_Data() = %v, want %v", output, test.willFail)
			}
			if test.willFail {
				t.Log(output)
				t.Errorf("Read_Token_Data(): %v should have failed but didn't", test.testName)
				t.Fail()
			}
		})
	}
}
