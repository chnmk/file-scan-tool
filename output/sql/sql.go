package sql_api

import (
	"database/sql"
	"fmt"

	"github.com/chnmk/file-scan-tool/output/structs"
	_ "modernc.org/sqlite"
)

type SQLFile struct {
	structs.Output
}

/*
func (p *SQLFile) HandleParams(params []string) {

}
*/

func (p SQLFile) HandleOutput(result [][2]string) {
	db, err := sql.Open("sqlite", "output.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	createTableQuery := "CREATE TABLE output (id INTEGER PRIMARY KEY, file VARCHAR(512));"
	insertQuery := "INSERT INTO output (file) VALUES (:file)"
	if p.PrintParent {
		createTableQuery = "CREATE TABLE output (id INTEGER PRIMARY KEY, file VARCHAR(512), parent VARCHAR(512));"
		insertQuery = "INSERT INTO output (file, parent) VALUES (:file, :parent)"
	}

	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range result {
		_, err = db.Exec(insertQuery, sql.Named("file", r[0]), sql.Named("parent", r[1]))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
