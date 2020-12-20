package model

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

type Post struct {
	User User
	Body string
}
