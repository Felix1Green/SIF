package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Email, Password string
}

var(
	Tokens = make(map[string]User)
	Users = make(map[string]User)
	mutex = &sync.Mutex{}
)

func genToken(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b);err != nil{
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}


func setCookie(w http.ResponseWriter, name, value string,d int) {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
	}
	if d != 0{
		expires := time.Now().AddDate(0,0,d)
		cookie.Expires = expires
	}
	http.SetCookie(w, &cookie)
}

func getCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func deleteCookie(w http.ResponseWriter,name string){
	cookie := http.Cookie{
		Name: name,
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user := User{}
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &user)
		if user.Email == "" || user.Password == ""{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := genToken(32)
		mutex.Lock()
		{
			Users[user.Email] = user
			Tokens[token] = user
		}
		mutex.Unlock()
		setCookie(w, "session_token", token, 30)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user := User{}
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &user)
		if user.Email == "" || user.Password == ""{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		mutex.Lock()
		user, ok := Users[user.Email]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
		}
		token := genToken(32)
		Tokens[token] = user
		mutex.Unlock()

		setCookie(w, "session_token", token, 30)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func logout(w http.ResponseWriter,r *http.Request){
	token := getCookie(r,"session_token")
	deleteCookie(w,"session_token")
	delete(Tokens, token)
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	http.HandleFunc("/register", signup)
	http.HandleFunc("/login", signin)
	http.HandleFunc("/logout",logout)
	fmt.Println("listings")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		return
	}
}