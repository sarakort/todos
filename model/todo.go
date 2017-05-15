package model

type Todo struct {
	ID      int    `json:"id"`
	Checked bool   `json:"checked"`
	Message string `json:"message"`
}
