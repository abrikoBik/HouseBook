// Package models contains all of the models used in the database, as well as all CRUD operations related to them.
package models

import (
	"fmt"
	"housebook/database"
	"housebook/util"
)

func Search[T Book | Trade](params util.SearchTable) (results []*T, err error) {
	q := database.BunDatabase.NewSelect().Model(&results)
	for n, v := range params {
		if v != "" {
			q = q.Where(fmt.Sprintf("%s = ?", n), v)
		}
	}
	err = q.Scan(database.Context)
	return
}

func List[T Book | Trade](limit, offset int) (results []*T, err error) {
	err = database.BunDatabase.
		NewSelect().
		Model(&results).
		Limit(limit).
		Offset(offset).
		Scan(database.Context)
	return
}
