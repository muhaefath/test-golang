package responses

type GetBook struct {
	BaseResponse
	Data []*Book `json:"data"`
}

type Book struct {
	Title           string `json:"title"`
	CoverID         int    `json:"cover_id"`
	CoverEditionKey string `json:"cover_edition_key"`
	Authors         string `json:"authors"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
