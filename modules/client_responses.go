package modules

type GetListBookByGenreResponse struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Works []struct {
		Title           string `json:"title"`
		CoverID         int    `json:"cover_id"`
		CoverEditionKey string `json:"cover_edition_key"`
		Authors         []struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"authors"`
	} `json:"works"`
}
