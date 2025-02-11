package domain

import "igrwijaya-go-template/internal/domain/common"

type Template struct {
	common.PrimaryEntity
	//[DOMAIN_TEMPLATE]
	common.AuditableEntity
}
