package dal

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// create a DB struct
type DB struct {
	sql *sqlx.DB
}

func New(path string) *DB {
	return &DB{sql: sqlx.MustConnect("sqlite3", path)}
}

func (db *DB) Close() {
	db.sql.Close()
}

func (db *DB) ReadAllTasks() (*[]TMTask, error) {
	var tasks []TMTask
	err := db.sql.Select(&tasks, "SELECT * FROM TMTask")
	if err != nil {
		return nil, err
	}

	fmt.Println(len(tasks))
	for i := range tasks {
		err = db.sql.Select(&tasks[i].Tags, "SELECT TMTag.* FROM TMTaskTag inner join TMTag on TMTaskTag.tags = TMTag.uuid WHERE tasks = ?", tasks[i].UUID)
		if err != nil {
			return nil, err
		}
	}
	return &tasks, nil
}
