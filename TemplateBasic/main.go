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
		u1, u2 := model.User{Username: "lxm"}, model.User{Username: "xml"}
		posts := []model.Post{
			{User: u1, Body: "aaa"},
			{User: u2, Body: "aaa"},
		}
		user := model.User{Username: "lxm"}
		v := model.IndexViewModel{Title: "Homepage", User: user, Posts: posts}
		tpl, _ := template.ParseFiles("src/mega/TemplateBasic/templates/index.html")
		err := tpl.Execute(w, &v)
		if err != nil {
			log.Fatal(err)
		}
	})
	_ = http.ListenAndServe(":8888", nil)
}
