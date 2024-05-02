package users

import "backend/domain/users"

func Login(user string, password string) users.LoginResponse {
	return users.LoginResponse{
		Token: "contraseña123",
	}
}
