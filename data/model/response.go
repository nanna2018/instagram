package model

import "time"

// Type RUsuario struct
type RUsuario struct {
	ID       int
	Name     string
	Password string
	Email    string
}

//Type RFoto struct
type RFoto struct {
	ID    int
	Lugar string
	Texto string
	Fecha time.Time
}

//RLogin struct
type RLogin struct {
	Name     string
	Password string
}
