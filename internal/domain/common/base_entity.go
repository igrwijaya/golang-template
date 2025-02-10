package common

import (
	"database/sql"
	"time"
)

type PrimaryEntity struct {
	Id uint `gorm:"primaryKey"`
}

type AuditableEntity struct {
	CreatedBy uint
	CreatedAt time.Time
	UpdatedBy uint
	UpdatedAt sql.NullTime
}