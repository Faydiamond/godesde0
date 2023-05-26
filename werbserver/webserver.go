package werbserver

import (
	"fmt"
	"net/http"
)

func WebServer() {
	http.HandleFunc("/", home) //manejador
	fmt.Println("Salida por localhost:3000 ")
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) { //dos interfaces de net http
	http.ServeFile(w, r, "./index.html")
}
