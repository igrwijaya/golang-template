package repositories

import (
	"igrwijaya-go-template/internal/domain/todo"
	"igrwijaya-go-template/internal/infrastructure/db"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func (repo *todoRepository) Migrate() {
	autoMigrateError := repo.db.AutoMigrate(&todo.Todo{})

	if autoMigrateError != nil {
		panic("Cannot migrate Todo entity. " + autoMigrateError.Error())
	}
}

func (repo *todoRepository) Create(todo todo.Todo) (todo.Todo, error) {
	createResult := repo.db.Create(&todo)

	if createResult.Error != nil {
		return todo, createResult.Error
	}

	return todo, nil
}

func (repo *todoRepository) Delete(id uint) error {
	deleteResult := repo.db.Where("Id = ?", id).Delete(&todo.Todo{})

	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}

func (repo *todoRepository) GetAll(page int, size int) ([]todo.Todo, error) {
	var todos []todo.Todo
	var offset = (page - 1) * size

	dbTodoResult := repo.db.Limit(size).Offset(offset).Find(&todos)

	if dbTodoResult.Error != nil {
		return todos, dbTodoResult.Error
	}

	return todos, nil
}

func (repo *todoRepository) Read(id uint) (*todo.Todo, error) {
	var todoEntity todo.Todo
	dbTodoResult := repo.db.Where("Id = ?", id).First(&todoEntity)

	if dbTodoResult.Error != nil {
		return nil, dbTodoResult.Error
	}

	return &todoEntity, nil
}

func (repo *todoRepository) Update(todo *todo.Todo) (*todo.Todo, error) {
	updateResult := repo.db.Save(&todo)

	if updateResult.Error != nil {
		return todo, updateResult.Error
	}

	return todo, nil
}

func NewTodoRepository(appDb db.AppDb) todo.TodoRepository {
	return &todoRepository{
		db: appDb.UseDefault(),
	}
}
