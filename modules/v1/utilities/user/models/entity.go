package models

type User struct {
	User_id   string `json:"user_id"`
	Full_name string `json:"full_name"`
	Email     string `json:"email"`
	Role_id   int    `json:"role_id"`
	Password  string `json:"password"`
	Address   string `json:"address"`
}
