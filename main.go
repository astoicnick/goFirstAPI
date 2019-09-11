package main

import (
	"fmt"
	"log"
	"net/http"
)

const usersAPIResp = `
<html>
<body>
<h1>Hi, thanks for calling with method '%v'</h1>
<h1>This is the '%v'nth call to this endpoint.</h1>
</body>
</html>`

var userCount int

func main() {
	http.HandleFunc("/users", userHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	userCount++
	s := fmt.Sprintf(usersAPIResp, r.Method, userCount)
	fmt.Fprintf(w, s)
}
