package conveter

import (
	"time"
)

type ShortenUrlData struct {
	OrginalUrl    string    `json:"orginal_url"`
	ShortUrl      string    `json:"short_url "`
	RedirectCount int       `json:"redirect_count"`
	CreatedAt     time.Time `json:"created_at "`
}

var ShortenUrlMemory []*ShortenUrlData
