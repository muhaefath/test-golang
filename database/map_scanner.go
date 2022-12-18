package database

import (
	"errors"

	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg/types"
)

type mscan struct {
	ValueMap map[string]string
}

func newMscan() *mscan {
	return &mscan{ValueMap: map[string]string{}}
}

func (s *mscan) ScanColumn(colIdx int, colName string, rd types.Reader, n int) error {
	str, err := rd.ReadFull()
	if err != nil {
		return err
	}
	s.ValueMap[string([]byte(colName))] = string([]byte(str))
	return nil
}

type MapScanner struct {
	Rows []map[string]string
}

func (s *MapScanner) Init() error {
	s.Rows = []map[string]string{}
	return nil
}

func (s *MapScanner) NewModel() orm.ColumnScanner {
	return newMscan()
}

func (s *MapScanner) AddModel(columnScanner orm.ColumnScanner) error {
	row, ok := columnScanner.(*mscan)
	if !ok {
		return errors.New("Column Scanner mismatch")
	}

	s.Rows = append(s.Rows, row.ValueMap)
	return nil
}

func NewMapScanner() *MapScanner {
	return &MapScanner{Rows: []map[string]string{}}
}
