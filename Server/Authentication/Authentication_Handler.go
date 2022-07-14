package Authentication_Handler

import (
	"Maria_Demo/Token_Handler"
	"Maria_Demo/convert_handler"
	"database/sql"
	"fmt"
)

func Sign_In(email string, password string, db *sql.DB) (string, error) {
	// check if the user exists
	var collected_email string = ""
	err := db.QueryRow(`
		SELECT user_email FROM user_identity.user_credentials WHERE user_email = ? AND password = ?
	`, email, password).Scan(&collected_email)
	// If there is an error, the user does not exist
	if err != nil {
		fmt.Println("[Authentication] Error:", err)
		return "", err
	}

	// Get the roles for the user
	roles, err := Get_Roles(email, db)
	if err != nil {
		return "", err
	}
	return Token_Handler.Generate_Token(roles, email)

}

func Register(email string, username string, password string, db *sql.DB) (string, error) {
	fmt.Println("[Authentication] Registering user: " + username)

	// Attempted to insert the user into the database
	rows, err := db.Query(`
		INSERT INTO user_identity.user_credentials (user_email, username, password) VALUES (?, ?, ?)
	`, email, username, password)
	if err != nil {
		fmt.Println("[Authentication] Error registering user: " + username)
		return "", err
	}
	rows, err = db.Query(`
		INSERT INTO user_identity.roles (user_email, role_id, role_name) VALUES (?, ?, ?)
	`, email, 1, "user")
	if err != nil {
		fmt.Println("[Authentication] Error registering user: " + username)
		return "", err
	}
	rows.Close()

	return Token_Handler.Generate_Token([]string{"user"}, email)
}

func Get_Roles(email string, db *sql.DB) ([]string, error) {
	// Get the roles for the user
	rows, err := db.Query(`
		SELECT role_name FROM user_identity.roles WHERE user_email = (SELECT user_email FROM user_identity.user_credentials WHERE user_email = ?);
	`, email)
	if err != nil {
		fmt.Println("[Authentication] Error getting roles for user: " + email)
		return nil, err
	}
	// Convert the rows into a slice of roles
	roles, err := convert_handler.Dataset_To_Roles(rows)
	if err != nil {
		fmt.Println("[Authentication] Error getting roles for user: " + email)
		return nil, err
	}

	rows.Close()
	return roles, nil
}

func Check_Roles(roles []string, required_roles []string) bool {
	for _, required_role := range required_roles {
		for _, role := range roles {
			if role == required_role {
				return true
			}
		}
	}
	return false
}
