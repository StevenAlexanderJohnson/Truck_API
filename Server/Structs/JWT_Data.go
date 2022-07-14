package struct_def

type Jwt_Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Jwt_Payload struct {
	User  string   `json:"user"`
	Roles []string `json:"roles"`
}

type Invalid_Algorithm_Error struct {
	Message string `default:"Invalid algorithm"`
}

func (m *Invalid_Algorithm_Error) Error() string {
	return m.Message
}

type Invalid_Token_Error struct {
	Message string `default:"Invalid token"`
}

func (m *Invalid_Token_Error) Error() string {
	return m.Message
}
