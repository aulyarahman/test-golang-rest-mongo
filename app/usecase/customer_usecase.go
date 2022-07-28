package usecase

import (
	"context"
	"github.com/aulyarahman/twitcat-service/app"
	"github.com/aulyarahman/twitcat-service/app/model"
)

type CustomerUsecase struct {
	config       *model.EnvConfig
	customerRepo app.CustomerRepositoryI
}

func NewCustomerUseCase(config *model.EnvConfig, cusRepo app.CustomerRepositoryI) app.CustomerUseCaseI {
	return &CustomerUsecase{
		config:       config,
		customerRepo: cusRepo,
	}
}

func (it *CustomerUsecase) GetData(ctx context.Context) (resp model.ListCustomer, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := it.customerRepo.GetData(ctx)
	if err != nil {
		return resp, err
	}

	return list, nil
}

func (it *CustomerUsecase) GetDataById(ctx context.Context, req model.CustomerById) (resp model.Customer, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	cus, err := it.customerRepo.GetDataById(ctx, req)
	if err != nil {
		return resp, err
	}
	return cus, nil
}

func (it *CustomerUsecase) InsertData(ctx context.Context, req model.Customer) (resp bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	err = it.customerRepo.InsertData(ctx, req)
	if err != nil {
		return resp, err
	}

	return true, nil
}

func (it *CustomerUsecase) UpdateData(ctx context.Context, req model.Customer) (resp bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	err = it.customerRepo.UpdateData(ctx, req)
	if err != nil {
		return resp, err
	}

	return true, nil
}

func (it *CustomerUsecase) DeleteData(ctx context.Context, req model.Customer) (resp bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	err = it.customerRepo.DeleteData(ctx, req)
	if err != nil {
		return resp, err
	}

	return true, nil
}
