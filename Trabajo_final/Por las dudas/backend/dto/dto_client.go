package dto

type ClientDatadto struct {
	ID        uint   `json:"id"`
	FirstName string `json:"nombre"`
	LastName  string `json:"apellido"`
	Email     string `json:"email"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Profesor  bool   `json:"profesor"`
}

type Clientesdto []ClientDatadto
