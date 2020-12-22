package controller

import (
	"log"
	"mega/vm"
	"net/http"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", middleAuth(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logout", middleAuth(logoutHandler))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	username, _ := getSessionUser(r)
	v := vop.GetVM(username)
	templates[tpName].Execute(w, &v)
}


func check(username, password string) bool {
	if username == "bonfy" && password == "abc123" {
		return true
	}
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !vm.CheckLogin(username, password) {
			v.AddError("username password not correct, please input again")
		}


		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "register.html"
	vop := vm.RegisterViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

