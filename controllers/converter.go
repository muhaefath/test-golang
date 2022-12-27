package controllers

import (
	"test-golang/controllers/requests"
	"test-golang/controllers/responses"
	"test-golang/models"
	"test-golang/modules"

	"github.com/astaxie/beego"
)

type UrlConvertController struct {
	beego.Controller
	pickUpHistoryOrmer models.PickUpHistoryOrmer
	urlConvertHandler  modules.ConverterHandler
}

func (c *UrlConvertController) Prepare() {
	c.urlConvertHandler = modules.NewConverterHandler()
}

// // @Param	genre=>genre				query   int     false   "genre"
// // @Param	genre=>offset				query   int     false   "offset"
// // @Param	genre=>limit				query   int     false   "limit"
// // @router /get_list_urlConvert [get]
// func (c *UrlConvertController) GetListUrlConvert(genre string, offset int, limit int) *responses.GetUrlConvert {
// 	response, err := c.urlConvertHandler.GetUrlConvertByGenre(genre, offset, limit)
// 	if err != nil {
// 		return &responses.GetUrlConvert{BaseResponse: responses.BaseResponse{
// 			Success: false,
// 			Message: err.Error(),
// 		}}
// 	}

// 	return &responses.GetUrlConvert{BaseResponse: responses.BaseResponse{
// 		Success: true,
// 		Message: "success",
// 	},
// 		Data: response,
// 	}
// }

// @Param  requests  body {UrlRequest} true "requests"
// @router /shorten_url [post]
func (c *UrlConvertController) ShortenUrl(requests requests.UrlRequest) *responses.BaseResponse {
	shortUrl, err := c.urlConvertHandler.ShortenUrl(requests)
	if err != nil {
		return &responses.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &responses.BaseResponse{
		Success: true,
		Message: "here is your shorten url: " + shortUrl,
	}
}

// @Param	genre=>url				query   string     false   "url"
// @router /redirect_url [get]
func (c *UrlConvertController) RedirectUrl(url string) *responses.BaseResponse {
	result, err := c.urlConvertHandler.RedirectUrl(url)
	if err != nil {
		return &responses.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	c.Redirect(result.OrginalUrl, 307)
	return &responses.BaseResponse{
		Success: true,
		Message: result.OrginalUrl,
	}
}

// @Param  requests  body {UrlRequest} true "requests"
// @router /stats_url [post]
func (c *UrlConvertController) StatsUrl(requests requests.UrlRequest) *responses.RedirectCountResponse {
	result, err := c.urlConvertHandler.StatsUrl(requests)
	if err != nil {
		return &responses.RedirectCountResponse{
			BaseResponse: responses.BaseResponse{
				Success: false,
				Message: err.Error(),
			},
		}
	}

	return &responses.RedirectCountResponse{
		BaseResponse: responses.BaseResponse{
			Success: true,
			Message: "success",
		},
		RedirectCount: result.RedirectCount,
		CreatedAt:     result.CreatedAt,
	}
}
