package responses

import "time"

type RedirectCountResponse struct {
	BaseResponse
	RedirectCount int       `json:"redirect_count"`
	CreatedAt     time.Time `json:"created_at "`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
