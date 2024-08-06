package services

import (
	"fmt"
	"log"

	"github.com/D4lvarez/data_process/pkg/modules/infrastructure/queries"
	"github.com/D4lvarez/data_process/pkg/modules/libs"
)

func InitializeService(repository libs.Repository) {
	queries := queries.GetCreateTablesQueries()

	for _, query := range queries {

		result, err := repository.Execute(query)
		if err != nil {
			log.Fatalf("Error creating table: %s", err)
		}

		rows, err := result.RowsAffected()
		if err != nil {
			log.Fatalf("Error on table: %s", err)
		}

		fmt.Printf("Table created: %d\n", rows)
	}
}
