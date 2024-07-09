package services

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type ServiceContext struct {
	Context  context.Context
	Database *pg.DB
}
