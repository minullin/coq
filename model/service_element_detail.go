package model

type ServiceElementDetail struct {
	Identifier   ServiceElementIdentifier `json:"identifier"`
	PropertyList []Element                `json:"propertyList"`
}
