package app

import (
	"context"
	"github.com/aulyarahman/twitcat-service/app/model"
)

type (
	CustomerRepositoryI interface {
		GetData(ctx context.Context) (resp model.ListCustomer, err error)
		GetDataById(ctx context.Context, req model.CustomerById) (resp model.Customer, err error)
		InsertData(ctx context.Context, req model.Customer) error
		UpdateData(ctx context.Context, req model.Customer) error
		DeleteData(ctx context.Context, req model.Customer) error
	}
)
