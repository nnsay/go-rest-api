package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/json", Json)
	router.HandleFunc("/hola/{name}", Hola)

	port := 8080
	portStr, has := os.LookupEnv("APP_PORT")
	if has {
		envPort, err := strconv.Atoi(portStr)
		if err == nil {
			port = envPort
		}
	}
	log.Printf("serving at %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := HipChatResponse{Color: "yellow", Message: "This is a Test", Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Hola(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := HipChatResponse{Color: "yellow", Message: "Hola " + name, Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Json(w http.ResponseWriter, r *http.Request) {
	g := Man
	response := Info{Gender: &g, Age: 33, Auth: &Auth{Name: "Jimmy"}}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        string `json:"notify"`
	MessageFormat string `json:"message_format"`
}

type HipChatrequest struct {
	Event string `json:"event"`
	Item  struct {
		Message struct {
			Date time.Time `json:"date"`
			From struct {
				ID          int    `json:"id"`
				MentionName string `json:"mention_name"`
				Name        string `json:"name"`
			} `json:"from"`
			ID       string        `json:"id"`
			Mentions []interface{} `json:"mentions"`
			Message  string        `json:"message"`
			Type     string        `json:"type"`
		} `json:"message"`
		Room struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"room"`
	} `json:"item"`
	WebhookID int `json:"webhook_id"`
}

type Address struct {
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
}

type Auth struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Gender int

const (
	Man Gender = iota
	Female
)

// As the docs say, "any nil pointer." -- make the struct a pointer. Pointers have obvious "empty" values: nil
type Info struct {
	Gender  *Gender  `json:"gender,omitempty"` // zero value weill be omited, but point can resolve the problem and the only nil pointer will be omitted.
	Age     int      `json:"age,omitempty"`
	Address *Address `json:"address,omitempty"`
	Auth    *Auth    `json:"auth,omitempty"`
}
