package migrations

import (
	"igrwijaya-go-template/internal/infrastructure/db"
	"igrwijaya-go-template/internal/infrastructure/repositories"
)

func MigrateEntity() {
	appDb := db.NewAppDb()

	todoRepo := repositories.NewTodoRepository(appDb)

	todoRepo.Migrate()
}
