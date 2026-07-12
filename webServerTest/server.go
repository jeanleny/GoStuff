package main

import (
	"fmt"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "someone triedukno")
	fmt.Fprintln(w, "sdfkjhdf  sdjhfkjsd fksdjhf asdkljhf lzi  ihsdf  sdkljhf sdans sdkljfhsdkbnj  blvksd bite")
}

func main() {
	http.HandleFunc("/", response)

	http.ListenAndServe(":8080", nil)
}
