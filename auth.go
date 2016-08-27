package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/linkedin"
)

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ASDASDASDSDAS")
	fmt.Println(r.URL.Query().Get("provider"))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
}

func AuthMain() {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))
	goth.UseProviders(
		linkedin.New("75ee4zo80mit43", "f5dY7jePaCsQfQT0", "http://localhost:8080/auth/linkedin/callback"),
	)
}

var userTemplate = `
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>`
