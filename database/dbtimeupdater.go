package database

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-pg/pg"
)

const (
	createdAtField  = "CreatedAt"
	updatedAtField  = "UpdatedAt"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

// NOTE: by using this middleware, updated_at will NOT be manually updatable.
type dbTimeUpdater struct{}

func (d dbTimeUpdater) flattenValue(raw reflect.Value) []reflect.Value {
	if raw.Kind() == reflect.Slice {
		result := []reflect.Value{}
		for i := 0; i < raw.Len(); i++ {
			cur := reflect.ValueOf(raw.Index(i).Interface())
			for cur.Kind() == reflect.Ptr {
				cur = cur.Elem()
			}
			result = append(result, cur)
		}
		return result
	}

	return []reflect.Value{raw}
}

func (d dbTimeUpdater) setFieldAs(row reflect.Value, fieldName string, value interface{}) {
	field := row.FieldByName(fieldName)
	if !field.IsValid() {
		return
	}

	if d.isIgnored(row, fieldName) {
		return
	}

	field.Set(reflect.ValueOf(value))
}

func (d dbTimeUpdater) isIgnored(row reflect.Value, fieldName string) bool {
	fieldStruct, ok := row.Type().FieldByName(fieldName)
	if !ok {
		return true
	}

	if fieldStruct.Tag.Get("sql") == "-" {
		return true
	}

	return false
}

func (d dbTimeUpdater) BeforeQuery(q *pg.QueryEvent) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CreatedAt/UpdatedAt hook\n Fix/disable this middleware if you see this message\n", r)
		}
	}()
	query := q.Query
	for reflect.TypeOf(query) == reflect.TypeOf(reflect.Value{}) {
		query = query.(reflect.Value).Interface()
	}
	queryType := reflect.TypeOf(query).String()
	t := time.Unix(0, time.Now().UnixNano()/1000*1000).UTC()

	if strings.Contains(queryType, "insertQuery") {
		params := reflect.ValueOf(q.Params[0]).
			MethodByName("Value").
			Call([]reflect.Value{})[0].
			Interface().(reflect.Value)
		flattenParams := d.flattenValue(params)

		for _, row := range flattenParams {
			d.setFieldAs(row, createdAtField, t)
			d.setFieldAs(row, updatedAtField, t)
		}
	} else if strings.Contains(queryType, "updateQuery") {
		params := reflect.ValueOf(q.Params[0]).
			MethodByName("Value").
			Call([]reflect.Value{})[0].
			Interface().(reflect.Value)
		flattenParams := d.flattenValue(params)

		for _, row := range flattenParams {
			d.setFieldAs(row, updatedAtField, t)
		}

		setLength := reflect.ValueOf(q.Query).
			MethodByName("Query").
			Call([]reflect.Value{})[0].
			Elem().
			FieldByName("set").
			Len()

		if setLength > 0 {
			// use the first element as sample
			if !d.isIgnored(flattenParams[0], updatedAtField) {
				reflect.ValueOf(q.Query).
					MethodByName("Query").
					Call([]reflect.Value{})[0].
					MethodByName("Set").
					Call([]reflect.Value{
						reflect.ValueOf(updatedAtColumn + " = ?"),
						reflect.ValueOf(t),
					})
			}
		}
	}
}

func (d dbTimeUpdater) AfterQuery(q *pg.QueryEvent) {
}
