package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	client "proyecto1/data/dataclient"
	"proyecto1/data/model"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (name string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			name = cookieValue["name"]
		}
	}
	return name
}

//NombreUsuario Función que muestra el nombre de usuario logueado
func NombreUsuario(response http.ResponseWriter, request *http.Request) {
	name := getUserName(request)
	fmt.Fprintf(response, name)
}
func setSession(name string, response http.ResponseWriter) {
	value := map[string]string{
		"name": name,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//Login Función para acceder a la página
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLogin {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	respuesta := false
	if e == nil {
		// datos que recibe del cliente
		var usuario model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.Name)

		if usuario.Name == "" || usuario.Password == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Contraseña de la base de datos
		password := client.LogearUsuario(&usuario)

		// Comprueba que las dos contraseñas sean iguales
		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(usuario.Password)); err != nil {
			fmt.Printf("No has podido inicar sesión")
		} else {
			respuesta = true
			setSession(usuario.Name, w)
			fmt.Println("Inicio de sesión realizado")
			getUserName(r)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, respuesta)
	}

	fmt.Fprintln(w, respuesta)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

//Uploader funcion guardar imagen
func Uploader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathUploader {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseMultipartForm(2000)
	//Coger el archivo y meterlo en una variable
	file, fileInfo, err := r.FormFile("archivo")
	//Coger el nombre del formulario y merterlo en una variable
	texto := r.FormValue("texto")

	name := getUserName(r)

	fmt.Println(texto, "Nombre Usuario: ", name)

	f, err := os.OpenFile("./files/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//La linea de abajo que esta comentada me manda a la página donde está el nombre del archivo

	//Esta linea de aqui abajo me manda a la pagina principal donde están todas las fotos
	http.Redirect(w, r, "/perfil", 301)
	//Datos de la base de datos
	id := client.ConsultaID(name)
	fmt.Println(id)
	//Subir foto a la base de datos
	go client.InsertarFoto(fileInfo.Filename, id)
}

//Insert Funcion inserta en la base de datos
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPeticion {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var usuario model.Usuario
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &usuario)

		fmt.Println(usuario.Name)

		if usuario.Name == "" || usuario.Password == "" || usuario.Email == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//para incriptar contraseña
		hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		usuario.Password = hashComoCadena
		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(usuario)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPeticion(&usuario)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}

}

//ListarFoto para sacar las fotos
func ListarFoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListarFoto {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	lista := client.ListarFoto()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))
}
