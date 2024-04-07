package model

type AgentDetails struct {
	MtpToUse                 string   `json:"mtpToUse"`
	CommunicationsAddress    string   `json:"communicationsAddress"`
	SupportedServiceElements []string `json:"supportedServiceElements"`
}
