package handlers

import "net/http"

//PathInicio ruta raiz
const PathInicio string = "/"

//PathJSFiles ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de estilos css
const PathCSSFiles string = "/css/"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathRegister ruta de register
const PathRegister string = "/register"

//PathLoginFile ruta de login
const PathLoginFile string = "/loginFile"

//PathLogin ruta de login
const PathLogin string = "/login"

//PathFoto ruta para foto
const PathFoto string = "/perfil"

//PathPerfilFile ruta de register
const PathPerfilFile string = "/perfil"

//PathUploader Ruta a la carpeta de estilos css
const PathUploader string = "/uploader"

//PathListarFoto ruta para sacar lasa fotos
const PathListarFoto string = "/lista"

//PathHeader ruta para header
//const PathHeader string = "/header"

//PathListadoPeticiones Ruta de obtención de las peticiones de hoy
//const PathListadoPeticiones string = "/lista"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

/*ManejadoresLista es el diccionario general de las peticiones que son manejadas por nuestro servidor*/
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JSFile
	Manejadores[PathCSSFiles] = CSSFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathRegister] = RegisterFile
	Manejadores[PathLoginFile] = LoginFile
	Manejadores[PathUploader] = Uploader
	Manejadores[PathLogin] = Login
	Manejadores[PathListarFoto] = ListarFoto
	//Manejadores[PathHeader] = Header
	Manejadores[PathPerfilFile] = PerfilFile

}
