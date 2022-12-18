package models

import (
	"time"

	"test-golang/database"

	"github.com/beego/beego/validation"
)

type PickUpHistory struct {
	ID        int        `json:"id"`
	CoverID   int        `json:"cover_id"`
	PickUpAt  time.Time  `json:"pick_up_at"`
	ReturnAt  *time.Time `json:"return_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type PickUpHistoryOrmer interface {
	WithOrmer(ormer database.Ormer) PickUpHistoryOrmer
	GetByCoverID(coverID int) (*PickUpHistory, error)
	Upsert(pickUpHistory *PickUpHistory) (*PickUpHistory, error)
}

func NewPickUpHistoryOrmer(orm database.Ormer) PickUpHistoryOrmer {
	return &pickUpHistoryOrmer{ormer: orm}
}

type pickUpHistoryOrmer struct {
	ormer database.Ormer
}

func (o *pickUpHistoryOrmer) WithOrmer(ormer database.Ormer) PickUpHistoryOrmer {
	return &pickUpHistoryOrmer{ormer: ormer}
}

// GetListPickUpHistory implements PickUpHistoryOrmer
func (o *pickUpHistoryOrmer) GetByCoverID(coverID int) (*PickUpHistory, error) {
	var rows = PickUpHistory{
		CoverID: coverID,
	}

	err := o.ormer.Model(&rows).
		WhereStruct(rows).Order("created_at DESC").First()
	if err != nil {
		if err == database.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &rows, nil
}

func (o *pickUpHistoryOrmer) Upsert(pickUpHistory *PickUpHistory) (*PickUpHistory, error) {
	valid := validation.Validation{}
	b, err := valid.Valid(pickUpHistory)

	if err != nil {
		return nil, err
	}

	if !b {
		return nil, valid.Errors[0]
	}

	_, err = o.ormer.Model(pickUpHistory).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, err
	}

	return pickUpHistory, nil
}
