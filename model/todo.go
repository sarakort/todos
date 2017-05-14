package model

type Todo struct {
	ID      int    `json:id,omitempty`
	Checked bool   `json:checked,omitempty`
	Message string `json:message`
}
