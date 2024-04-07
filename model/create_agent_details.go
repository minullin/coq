package model

type CreateAgentDetails struct {
	EndpointID               string   `json:"endpointID"`
	MtpToUse                 string   `json:"mtpToUse"`
	CommunicationsAddress    string   `json:"communicationsAddress"`
	SupportedServiceElements []string `json:"supportedServiceElements"`
}
