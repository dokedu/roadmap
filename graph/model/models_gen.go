// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Token string `json:"token"`
}