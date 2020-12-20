package main

import (
	"html/template"
	"mega/TemplateBasic/model"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		user := model.User{Username: "林小墨"}
		tpl, _ := template.New("").Parse(`<html>
            <head>
                <title>Home Page - Bonfy</title>
            </head>
            <body>
                <h1>Hello, {{.Username}}!</h1>
            </body>
        </html>
    `)
		tpl.Execute(w, &user)
	})
	http.ListenAndServe(":8000", nil)
}
