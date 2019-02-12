package model

// RUsuario struct
type RUsuario struct {
	ID       int
	Name     string
	Password string
	Email    string
}

// RFoto struct
type RFoto struct {
	ID   int
	Name string
}

//RLogin struct
type RLogin struct {
	Name     string
	Password string
}
