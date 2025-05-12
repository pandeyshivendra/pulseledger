package dto

type APIResponse struct {
	Success      bool   `json:"success" example:"true"`
	ErrorMessage string `json:"error_message,omitempty" example:"Invalid request payload"`
	Data         any    `json:"data,omitempty"`
}
