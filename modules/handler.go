package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"test-golang/controllers/requests"
	"test-golang/controllers/responses"
	conveter "test-golang/utils/converter"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type converterHandler struct {
}

func NewConverterHandler() ConverterHandler {

	return &converterHandler{}
}

type ConverterHandler interface {
	ShortenUrl(request requests.UrlRequest) (string, error)
	RedirectUrl(url string) (*conveter.ShortenUrlData, error)
	StatsUrl(request requests.UrlRequest) (*responses.RedirectCountResponse, error)
}

func (h *converterHandler) ShortenUrl(request requests.UrlRequest) (string, error) {
	tempShortenUrl := conveter.ShortenUrlMemory
	shortenUrl := ""
	for {
		shortenUrl = RandStringRunes(6)
		isExist := false
		for _, value := range tempShortenUrl {
			if value.ShortUrl == shortenUrl {
				isExist = true
			}
		}

		if !isExist {
			break
		}
	}

	temp := conveter.ShortenUrlData{
		OrginalUrl:    request.RequestUrl,
		RedirectCount: 0,
		CreatedAt:     time.Now(),
		ShortUrl:      shortenUrl,
	}

	tempShortenUrl = append(tempShortenUrl, &temp)

	tempShortenUrlss, _ := json.Marshal(tempShortenUrl)
	fmt.Println("current list url 1: ", string(tempShortenUrlss))

	conveter.ShortenUrlMemory = tempShortenUrl

	return temp.ShortUrl, nil
}

func (h *converterHandler) RedirectUrl(url string) (*conveter.ShortenUrlData, error) {
	tempShortenUrl := conveter.ShortenUrlMemory

	isExist := false
	resultUrl := conveter.ShortenUrlData{}
	for _, value := range tempShortenUrl {
		if value.ShortUrl == url {
			isExist = true
			resultUrl = *value
			value.RedirectCount++
		}
	}

	if !isExist {
		return nil, errors.New("Url not found")
	}

	conveter.ShortenUrlMemory = tempShortenUrl

	tempShortenUrlss, _ := json.Marshal(tempShortenUrl)
	fmt.Println("current list url 2: ", string(tempShortenUrlss))

	return &resultUrl, nil
}

func (h *converterHandler) StatsUrl(request requests.UrlRequest) (*responses.RedirectCountResponse, error) {
	tempShortenUrl := conveter.ShortenUrlMemory
	isExist := false
	resultUrl := responses.RedirectCountResponse{}
	for _, value := range tempShortenUrl {
		if value.ShortUrl == request.RequestUrl {
			isExist = true
			resultUrl = responses.RedirectCountResponse{
				RedirectCount: value.RedirectCount,
				CreatedAt:     value.CreatedAt,
			}
		}
	}

	if !isExist {
		return nil, errors.New("Url not found")
	}

	tempShortenUrlss, _ := json.Marshal(tempShortenUrl)
	fmt.Println("current list url 3: ", string(tempShortenUrlss))

	return &resultUrl, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
