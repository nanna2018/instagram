package dataclient

import (
	"database/sql"
	"fmt"
	"proyecto1/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarPeticion funcion de peticion
func InsertarPeticion(objeto *model.Usuario) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/proyecto1")

	if err != nil {
		panic(err.Error()) //si se abre bien
	}

	defer db.Close() //cerrar la conexion nosotros. Hay que cerrarlo siempre

	insert, err := db.Query("INSERT INTO Usuario(name,password,email)VALUES (?, ?, ?)", objeto.Name, objeto.Password, objeto.Email)
	//Inserta una nueva peticion en la base de datos,guardar fechas de horarios en utc.
	if err != nil {
		panic(err.Error())
	}
	insert.Close()

}

//LogearUsuario logear usuario
func LogearUsuario(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/proyecto1")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Consultamos todos los idiomas de la base de datos
	comando := "SELECT Password FROM Usuario WHERE (Name = '" + objeto.Name + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Password FROM Usuario WHERE (Name = '" + objeto.Name + "')")

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado string
	for query.Next() {
		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//InsertarFoto accion de insertar foto
func InsertarFoto(name string, id int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/proyecto1")

	if err != nil {
		panic(err.Error()) //si se abre bien
	}

	defer db.Close() //cerrar la conexion nosotros. Hay que cerrarlo siempre

	insert, err := db.Query("INSERT  Foto (Name,Usuario_ID) VALUES (?, ?)", name, id)

	//Inserta una nueva peticion en la base de datos,guardar fechas de horarios en utc.
	if err != nil {
		panic(err.Error())
	}
	insert.Close()

}

//ConsultaID test
func ConsultaID(name string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/proyecto1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT ID FROM Usuario WHERE (Name = '" + name + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM Usuario WHERE (Name = '" + name + "')")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado int
	for query.Next() {
		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//ListarFoto test
func ListarFoto() []model.RFoto {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/proyecto1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID, Name  FROM Foto"
	fmt.Println(comando)

	query, err := db.Query("SELECT ID, Name FROM Foto")

	if err != nil {
		panic(err.Error())
	}

	resultado := make([]model.RFoto, 0)

	for query.Next() {
		var foto = model.RFoto{}
		err = query.Scan(&foto.ID, &foto.Name)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, foto)
	}
	return resultado
}
