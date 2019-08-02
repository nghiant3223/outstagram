package models

type AppError struct {
	// For API consumer
	Code   int                    `json:"code"`
	Params map[string]interface{} `json:"params"`

	// For end-user
	Message string `json:"message"`

	// For API developer
	ID    string `json:"id"`
	Where string `json:"where"`
}

func NewAppError(id, where, message string, params map[string]interface{}, code int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		ID:      id,
		Where:   where,
		Params:  params,
	}
}
