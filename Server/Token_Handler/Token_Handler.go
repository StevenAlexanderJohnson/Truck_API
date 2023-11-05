package Token_Handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"time"

	struct_def "Maria_Demo/Structs"
	"Maria_Demo/convert_handler"
	"strings"
)

/*
Given a header, payload, and algorithm, generate a signature
The signature is used to verify the token is valid upon request.
This is done by taking the header and payload value from the passed token and generating
a signature with the passed values. Then compare the signature from the passed token with
the newly generated one. If the signatures match then the token is valid.

If any of the data in the header or payload has been changed, the token signature will not match
and we can reject the request.
*/
func generate_signature(header []byte, payload []byte, algorithm string) (string, error) {
	// This is a secret value and should not be shared publicly.
	// It should be stored in some environment variable or a config file but for demo purposes,
	// I am storing it in a string literal.
	key := []byte("6v9y$B&E)H@MbQeThWmZq4t7w!z%C*F-JaNdRfUjXn2r5u8x/A?D(G+KbPeShVkY")

	var h hash.Hash
	if algorithm == "HS256" {
		h = hmac.New(sha256.New, key)
	} else if algorithm == "HS384" {
		h = hmac.New(sha512.New384, key)
	} else {
		return "", &struct_def.Invalid_Token_Error{}
	}

	// Base64-encode the header and payload before signing
	encodedHeader := base64.RawURLEncoding.EncodeToString(header)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payload)

	// Concatenate the encoded header and payload with a "." in between
	data := []byte(encodedHeader + "." + encodedPayload)

	// Sign the data using the selected algorithm
	h.Write(data)

	// Return the base64-encoded signature
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil)), nil
}

/*
Given a header and payload, generate a token.
Convert the header and payload to a string, base64 encode them, then concatenate them with the signature.
*/
func Sign_Token(header struct_def.Jwt_Header, payload struct_def.Jwt_Payload) (string, error) {
	header_string, err := convert_handler.Struct_ToBytes(header)
	if err != nil {
		return "", err
	}
	payload_string, err := convert_handler.Struct_ToBytes(payload)
	if err != nil {
		return "", err
	}
	signature, err := generate_signature(header_string, payload_string, "HS256")
	if err != nil {
		fmt.Println("[Token_Handler] Signature Error:", err)
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(header_string) + "." +
			base64.RawURLEncoding.EncodeToString(payload_string) + "." +
			signature,
		nil
}

/*
Using the logic from the signature comment, we verify the passed token is valid.
*/
func Verify_Token(token string) (bool, error) {
	// Split the token into header, payload, and signature
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, nil
	}
	// Convert the header and payload to base64 strings
	header, err1 := base64.RawURLEncoding.DecodeString(parts[0])
	payload, err2 := base64.RawURLEncoding.DecodeString(parts[1])
	if err1 != nil || err2 != nil {
		return false, &struct_def.Invalid_Token_Error{}
	}
	// Generate a signature using the passed values
	generated_signature, err := generate_signature(header, payload, "HS256")
	if err != nil {
		return false, err
	}
	// Compare the passed signature with the generated one. If they match it returns true.
	return parts[2] == generated_signature, nil
}

func Generate_Token(roles []string, email string) (string, error) {
	header := struct_def.Jwt_Header{Algorithm: "HS256"}
	// The token expires in one hour from the time of generation.
	payload := struct_def.Jwt_Payload{
		User:     email,
		Roles:    roles,
		Expires:  time.Now().Add(time.Hour + time.Duration(1)).UTC().UnixNano(),
		Assigned: time.Now().UTC().UnixNano(),
	}
	return Sign_Token(header, payload)
}

// Parse the token payload into a string to be read.
func Read_Token_Data(token string) ([]byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, &struct_def.Invalid_Token_Error{}
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	return payload, nil
}
