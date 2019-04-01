package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(s.LogWriter, "--------------------------------\n")
	s.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprint(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

func main() {
	http.Handle("/", &LoggerServer{
		LogWriter: os.Stdout,
		Handler:   &MyServer{},
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
