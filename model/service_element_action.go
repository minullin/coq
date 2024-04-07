package model

type ServiceElementAction struct {
	Name      string `json:"name"`
	InputArgs []Arg  `json:"inputArgs"`
}
