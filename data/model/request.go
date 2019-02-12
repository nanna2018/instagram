package model

import (
	"time"
)

//Usuario struct
type Usuario struct {
	ID       int
	Name     string
	Password string
	Email    string
}

//Foto struct
type Foto struct {
	ID   int64
	Name string
}

//Filtro struct
type Filtro struct {
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
