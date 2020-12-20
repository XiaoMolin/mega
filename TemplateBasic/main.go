package main

import (
	"fmt"
	"html/template"
	"log"
	"mega/TemplateBasic/model"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println(filepath.Join(cwd, "mega/TemplateBasic/templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := model.User{Username: "lxm"}
		tpl, _ := template.ParseFiles("src/mega/TemplateBasic/templates/index.html")
		err := tpl.Execute(w, &user)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8888", nil)
}
