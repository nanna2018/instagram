package model

import (
	"time"
)

//Usuario struct
type Usuario struct {
	Name     string
	Password string
	Email    string
}

//Foto struct
type Foto struct {
	ID    int64
	Lugar string
	Fecha time.Time
}

// Comentario struct
type Comentario struct {
	ID    int
	Texto string
}

//Login struct
type Login struct {
	Name     string
	Password string
}
