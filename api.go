package main

import (
    "fmt"
    "net/http"
    "strings"
)

 

func main() {
	//usamos a primeira entrada o ola mundo 
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world!")
    })
	//segunda entrada com senha e login 
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        username := r.FormValue("username")
        password := r.FormValue("password")
		//caso o usuario tenha acesso 
        if username == "admin" && password == "123456" {
            session := fmt.Sprintf("%s%d", username, time.Now().UnixNano())
            cookie := http.Cookie{
                Name: "session",
                Value: session,
                Expires: time.Now().Add(time.Hour),
            }
            http.SetCookie(w, &cookie)

            fmt.Fprintf(w, session)
        } else {
            w.WriteHeader(http.StatusUnauthorized)
        }
    })

    http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
        session := r.URL.Query().Get("session")

        if session == "" {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if strings.HasPrefix(session, "admin") {
            fmt.Fprintf(w, "true")
        } else {
            fmt.Fprintf(w, "false")
        }
    })

    http.ListenAndServe(":3000", nil)
}
