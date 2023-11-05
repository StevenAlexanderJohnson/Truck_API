package convert_handler

import (
	"encoding/base64"
	"testing"

	struct_def "Maria_Demo/Structs"
)

func TestStruct_ToString(t *testing.T) {
	// arrange
	truckData := struct_def.Truck_Data{
		Truck_Id:     1,
		Truck_Number: "TRUCK001",
		Truck_Type:   2,
		Truck_Price:  500,
	}
	expectedOutput := `{"Truck_Id":1,"Truck_Number":"TRUCK001","Truck_Type":2,"Truck_Price":500}`

	// act
	output, err := Struct_ToString(truckData)

	// assert
	if err != nil {
		t.Errorf("Failed to convert truck data to string: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output %s, but got %s", expectedOutput, output)
	}
}

func TestStruct_ToBytes(t *testing.T) {
	tests := []struct {
		name           string
		input          interface{}
		expectedOutput string
		willFail       bool
	}{
		{
			name: "Sample JWT Header: 1",
			input: struct_def.Jwt_Header{
				Algorithm: "HS256",
				Type:      "JWT",
			},
			expectedOutput: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			willFail:       false,
		},
		{
			name: "Sample JWT Header: 2",
			input: struct_def.Jwt_Header{
				Algorithm: "HS384",
				Type:      "JWT",
			},
			expectedOutput: "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9",
			willFail:       false,
		},
		{
			name: "Sample JWT Header: 3",
			input: struct_def.Jwt_Header{
				Algorithm: "HS256",
				Type:      "",
			},
			expectedOutput: "eyJhbGciOiJIUzI1NiIsInR5cCI6IiJ9",
			willFail:       false,
		},
		{
			name: "Sample JWT Payload: 1",
			input: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin", "user"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			expectedOutput: "eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0",
			willFail:       false,
		},
		{
			name: "Sample JWT Payload: 2",
			input: struct_def.Jwt_Payload{
				User:     "test@email.com",
				Roles:    []string{"admin"},
				Expires:  1699220445493665300,
				Assigned: 1699216845493665300,
			},
			expectedOutput: "eyJ1c2VyIjoidGVzdEBlbWFpbC5jb20iLCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXSwiZXhwIjoxNjk5MjIwNDQ1NDkzNjY1MzAwLCJhc3NpZ25lZCI6MTY5OTIxNjg0NTQ5MzY2NTMwMH0",
			willFail:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := Struct_ToBytes(test.input)
			if err != nil && !test.willFail {
				t.Errorf("Test %v failed: %v", test.name, err)
				t.Fail()
				return
			}
			if base64.RawURLEncoding.EncodeToString(output) != test.expectedOutput {
				if test.willFail {
					return
				}
				t.Errorf("Test %v failed: Expected output %v, but got %v", test.name, test.expectedOutput, base64.RawURLEncoding.EncodeToString(output))
				t.Fail()
				return
			}
		})
	}
}
