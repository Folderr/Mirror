package endpoints

import (
	"fmt"
	"net/http"
	"net/url"
)

func SetupLinks(isOnline *bool, domain *url.URL, mux *http.ServeMux) {
	mux.HandleFunc("/l/{link}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("link")
		fmt.Println(id)
		fmt.Fprint(w, "Recieved id: id")
	})
}
