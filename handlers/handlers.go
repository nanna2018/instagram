package handlers

import "net/http"

//PathInicio ruta raiz
const PathInicio string = "/"

//PathJSFiles ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de estilos css
const PathCSSFiles string = "/css/"

//PathLogin ruta de login
const PathLogin string = "/login"

//PathLoginFile ruta de login
const PathLoginFile string = "/loginFile"

//PathRegister ruta de register
const PathRegister string = "/register"

//PathPerfilFile ruta de register
const PathPerfilFile string = "/perfil"

//PathImagen Ruta a la carpeta de estilos css
const PathImagen string = "/imagen"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathListadoPeticiones Ruta de obtención de las peticiones de hoy
const PathListadoPeticiones string = "/lista"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//ManejadoresLista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JSFile
	Manejadores[PathCSSFiles] = CSSFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathLogin] = Login
	Manejadores[PathPerfilFile] = PerfilFile

}
