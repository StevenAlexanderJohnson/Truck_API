package struct_def

type Login_Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register_Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}