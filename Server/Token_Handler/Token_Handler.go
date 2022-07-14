package Token_Handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"hash"

	struct_def "Maria_Demo/Structs"
	"Maria_Demo/convert_handler"
	"strings"
)

func generate_signature(header string, payload string, algorithm string) (string, error) {
	key := []byte("6v9y$B&E)H@MbQeThWmZq4t7w!z%C*F-JaNdRfUjXn2r5u8x/A?D(G+KbPeShVkY")

	var h hash.Hash

	if algorithm == "HS256" {
		h = hmac.New(sha256.New, key)
	} else if algorithm == "HS384" {
		h = hmac.New(sha512.New384, key)
	} else if algorithm == "HS512" {
		h = hmac.New(sha512.New384, key)
	} else {
		return "", &struct_def.Invalid_Token_Error{}
	}
	return base64.RawURLEncoding.EncodeToString(h.Sum([]byte(header + "." + payload))), nil
}

func Sign_Token(header struct_def.Jwt_Header, payload struct_def.Jwt_Payload) (string, error) {
	header_string, err := convert_handler.Struct_ToString(header)
	payload_string, err := convert_handler.Struct_ToString(payload)
	signature, err := generate_signature(header_string, payload_string, "HS256")
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString([]byte(header_string)) + "." +
			base64.RawURLEncoding.EncodeToString([]byte(payload_string)) + "." +
			signature,
		nil
}

func Verify_Token(token string) (bool, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, nil
	}
	header, err1 := base64.RawURLEncoding.DecodeString(parts[0])
	payload, err2 := base64.RawURLEncoding.DecodeString(parts[1])
	if err1 != nil || err2 != nil {
		return false, &struct_def.Invalid_Token_Error{}
	}
	generated_signature, err := generate_signature(string(header), string(payload), "HS256")
	if err != nil {
		return false, err
	}
	return parts[2] == generated_signature, nil
}

func Generate_Token(roles []string, email string) (string, error) {
	header := struct_def.Jwt_Header{Algorithm: "HS256"}
	payload := struct_def.Jwt_Payload{User: email, Roles: roles}
	return Sign_Token(header, payload)
}

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
