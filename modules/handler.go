package modules

import (
	"errors"
	"fmt"
	"test-golang/controllers/requests"
	"test-golang/controllers/responses"
	"test-golang/models"
	"test-golang/utils/httpclient"
	"time"
)

const TimeLayoutRFC3339Milli = "2006-01-02T15:04:05.999Z07:00"

type bookHandler struct {
	pickUpHistoryOrmer models.PickUpHistoryOrmer
	client             Client
}

func NewBookHandler(pickUpHistoryOrmer models.PickUpHistoryOrmer) BookHandler {
	clientNew := NewClient(httpclient.NewDoer())
	return &bookHandler{
		pickUpHistoryOrmer: pickUpHistoryOrmer,
		client:             clientNew,
	}
}

type BookHandler interface {
	GetBookByGenre(genre string, offset int, limit int) ([]*responses.Book, error)
	PickUpBook(request requests.PickUpRequest) error
	ReturnBook(request requests.ReturnRequest) error
}

func (h *bookHandler) GetBookByGenre(genre string, offset int, limit int) ([]*responses.Book, error) {
	bookList := []*responses.Book{}
	resp, _, err := h.client.GetListBookByGenre(limit, offset, genre)
	if err != nil {
		fmt.Println("err GetListBookByGenre", err.Error())
		return nil, err
	}

	for _, book := range resp.Works {
		authors := ""
		for idx, value := range book.Authors {
			if idx >= len(book.Authors)-1 {
				if idx == 1 {
					authors = authors + " "
				}

				authors = authors + "and " + value.Name
				break
			} else if idx == 0 {
				authors = authors + value.Name
			}

			authors = authors + ", " + value.Name
		}

		bookList = append(bookList, &responses.Book{
			Title:           book.Title,
			CoverID:         book.CoverID,
			CoverEditionKey: book.CoverEditionKey,
			Authors:         authors,
		})
	}

	return bookList, nil
}

func (h *bookHandler) PickUpBook(request requests.PickUpRequest) error {
	pickUpAt, err := time.Parse(TimeLayoutRFC3339Milli, request.PickUpAt)
	if err != nil {
		fmt.Println("err Parse", err.Error())
		return err
	}

	pickUpHistory, err := h.pickUpHistoryOrmer.GetByCoverID(request.CoverID)
	if err != nil {
		fmt.Println("err Upsert", err.Error())
		return err
	}

	if pickUpHistory != nil && pickUpHistory.ReturnAt == nil {
		err = errors.New("book still borow by others")
		fmt.Println("err Upsert", err.Error())
		return err
	}

	pickUp := models.PickUpHistory{
		CoverID:  request.CoverID,
		PickUpAt: pickUpAt,
	}

	_, err = h.pickUpHistoryOrmer.Upsert(&pickUp)
	if err != nil {
		fmt.Println("err Upsert", err.Error())
		return err
	}

	return nil
}

func (h *bookHandler) ReturnBook(request requests.ReturnRequest) error {
	returnAt, err := time.Parse(TimeLayoutRFC3339Milli, request.ReturnAt)
	if err != nil {
		fmt.Println("err Parse", err.Error())
		return err
	}

	pickUpHistory, err := h.pickUpHistoryOrmer.GetByCoverID(request.CoverID)
	if err != nil {
		fmt.Println("err Upsert", err.Error())
		return err
	}

	if pickUpHistory == nil {
		err = errors.New("pick up book not found")
		fmt.Println("err Upsert", err.Error())
		return err
	}

	if pickUpHistory != nil && pickUpHistory.ReturnAt != nil {
		err = errors.New("book is returned")
		fmt.Println("err Upsert", err.Error())
		return err
	}

	pickUpHistory.ReturnAt = &returnAt
	_, err = h.pickUpHistoryOrmer.Upsert(pickUpHistory)
	if err != nil {
		fmt.Println("err Upsert", err.Error())
		return err
	}

	return nil
}
