package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

type AccessToken struct {
	Token string `json:"access_token"`
}

type ApiPerson struct {
	Id        string `xml:"id"`
	FirstName string `xml:"first-name"`
	LastName  string `xml:"last-name"`
}

type ToriiResponse struct {
	Id string `json:"id"`
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
}

func AuthSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("ASDASDASDSDAS")
	fmt.Println(r.URL.Query().Get("provider"))
	fmt.Println(r.URL.Query().Get("auth-code"))

	authCode := r.URL.Query().Get("auth-code")

	_, body, _ := gorequest.New().Post("https://www.linkedin.com/uas/oauth2/accessToken").
		Send(`grant_type=authorization_code&code=` + authCode + `&redirect_uri=http://localhost:4200/callback&client_id=75ee4zo80mit43&client_secret=f5dY7jePaCsQfQT0`).
		End()

	fmt.Println("resp", body)

	var at AccessToken
	err := json.Unmarshal([]byte(body), &at)

	if err == nil {
		_, bodyPerson, _ := gorequest.New().Get("https://api.linkedin.com/v1/people/~").
			Set("Authorization", "Bearer "+at.Token).
			End()

		fmt.Println(bodyPerson)

		var person ApiPerson

		err := xml.Unmarshal([]byte(bodyPerson), &person)

		if err == nil {
			fmt.Println("OK!")
			fmt.Println(person)
			json.NewEncoder(w).Encode(ToriiResponse{Id: person.Id, FirstName: person.FirstName, LastName: person.LastName})
		}

	}

	// user, err := gothic.CompleteUserAuth(w, r)
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// 	return
	// }
	// t, _ := template.New("foo").Parse(userTemplate)
	// t.Execute(w, user)
}
