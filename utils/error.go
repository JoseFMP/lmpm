package utils

type LMPMErrorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
