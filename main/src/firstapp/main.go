package main
import (
	"fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)
func YourHandler(w http.ResponseWriter, r *http.Request) {
	var key string = "salut"
	
	key = r.URL.Query().Get("name")

    fmt.Fprintf(w,"%s",key)
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8512", r))
}