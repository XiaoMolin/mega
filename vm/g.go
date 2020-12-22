package vm

import (
	"log"
	"mega/model"
)

type BaseViewModel struct {
	Title string
	CurrentUser string
}

// SetTitle func
func (v *BaseViewModel)SetTitle(title string){
	v.Title=title
}

// SetCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

// CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}
