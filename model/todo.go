package model

type Todo struct {
	ID      int    `json:id,omitempty`
	Checked bool   `json:id,omitempty`
	Message string `json:message`
}
