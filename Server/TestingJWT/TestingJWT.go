package main

import (
	struct_def "Maria_Demo/Structs"
	"Maria_Demo/Token_Handler"
	"fmt"
)

func main() {
	header := struct_def.Jwt_Header{Algorithm: "HS256", Type: "JWT"}
	payload := struct_def.Jwt_Payload{User: "Testing", Roles: []string{"admin", "user"}}
	fmt.Println("Testing JWT")
	token, err := Token_Handler.Sign_Token(header, payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	fmt.Println(Token_Handler.Verify_Token(token))
}
