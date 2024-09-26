package main

import (
	"fmt"
	"log"

	"github.com/0x434d53/things-go/dal"
)

func main() {
	path := "/Users/seufert/dev/0x434d53/things-go/Things Database.thingsdatabase/main.sqlite"
	db := dal.New(path)
	tasks, err := db.ReadAllTasks()
	if err != nil {
		log.Fatal(err)
	}
	for _, task := range *tasks {
		fmt.Printf("%#v", task.CreationDate)
	}
}
