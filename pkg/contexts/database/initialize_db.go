package contexts

import (
	"github.com/D4lvarez/data_process/pkg/config"
	"github.com/D4lvarez/data_process/pkg/modules/infrastructure/services"
	"github.com/D4lvarez/data_process/pkg/modules/libs"
)

func InitializeDatabase() {
	db := config.GetDatabaseConnection()
	repository := libs.NewRepository(db)
	services.InitializeService(repository)
}
