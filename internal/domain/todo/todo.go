package todo

import "igrwijaya-go-template/internal/domain/common"

type Todo struct {
	common.PrimaryEntity
	Title       string
	Description string
	common.AuditableEntity
}
