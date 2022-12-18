package controllers

import (
	"test-golang/controllers/requests"
	"test-golang/controllers/responses"
	"test-golang/database"
	"test-golang/models"
	"test-golang/modules"

	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
	pickUpHistoryOrmer models.PickUpHistoryOrmer
	bookHandler        modules.BookHandler
}

func (c *BookController) Prepare() {
	ormer := database.GetOrmer()
	c.pickUpHistoryOrmer = models.NewPickUpHistoryOrmer(ormer)
	c.bookHandler = modules.NewBookHandler(c.pickUpHistoryOrmer)
}

// @Param	genre=>genre				query   int     false   "genre"
// @Param	genre=>offset				query   int     false   "offset"
// @Param	genre=>limit				query   int     false   "limit"
// @router /get_list_book [get]
func (c *BookController) GetListBook(genre string, offset int, limit int) *responses.GetBook {
	response, err := c.bookHandler.GetBookByGenre(genre, offset, limit)
	if err != nil {
		return &responses.GetBook{BaseResponse: responses.BaseResponse{
			Success: false,
			Message: err.Error(),
		}}
	}

	return &responses.GetBook{BaseResponse: responses.BaseResponse{
		Success: true,
		Message: "success",
	},
		Data: response,
	}
}

// @Param  requests  body {PickUpRequest} true "requests"
// @router /pick_up_book [post]
func (c *BookController) PickUpBook(requests requests.PickUpRequest) *responses.BaseResponse {
	err := c.bookHandler.PickUpBook(requests)
	if err != nil {
		return &responses.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &responses.BaseResponse{
		Success: true,
		Message: "success",
	}
}

// @Param  requests  body {PickUpRequest} true "requests"
// @router /return_book [post]
func (c *BookController) ReturnBook(requests requests.ReturnRequest) *responses.BaseResponse {
	err := c.bookHandler.ReturnBook(requests)
	if err != nil {
		return &responses.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &responses.BaseResponse{
		Success: true,
		Message: "success",
	}
}
