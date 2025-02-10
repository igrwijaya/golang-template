package todo

import "igrwijaya-go-template/internal/domain/common"

type TodoRepository interface {
	common.BaseRepository
	GetAll(page int, size int) ([]Todo, error)
	Create(todo *Todo) error
	Read(id uint) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id uint) error
}
