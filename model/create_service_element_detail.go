package model

type CreateServiceElementDetail struct {
	Name         string                   `json:"name"`
	Identifier   ServiceElementIdentifier `json:"identifier"`
	PropertyList []Element                `json:"propertyList"`
}
