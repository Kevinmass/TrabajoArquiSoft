package dto

type ClientDatadto struct {
	ID        uint   `json:"id"`
	FirstName string `json:"nombre"`
	LastName  string `json:"apellido"`
	Email     string `json:"email"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
}

type Clientesdto []ClientDatadto
