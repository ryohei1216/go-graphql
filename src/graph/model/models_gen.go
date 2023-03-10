// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InputTodos struct {
	UserID string `json:"userId"`
}

type InputUser struct {
	ID string `json:"id"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
