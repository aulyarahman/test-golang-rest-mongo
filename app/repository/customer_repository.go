package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/aulyarahman/twitcat-service/app"
	"github.com/aulyarahman/twitcat-service/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CustomerRepository struct {
	mongoDB *mongo.Database
}

const clx = "customer"

func NewCustomerRepository(mongo *mongo.Database) app.CustomerRepositoryI {
	return &CustomerRepository{mongoDB: mongo}
}

func (c CustomerRepository) GetData(ctx context.Context) (resp model.ListCustomer, err error) {
	query, err := c.mongoDB.Collection(clx).Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.ListCustomer{}, err
	}
	defer query.Close(ctx)

	listCustomer := make([]model.Customer, 0)
	for query.Next(ctx) {
		var row model.Customer
		err := query.Decode(&row)
		if err != nil {
			log.Fatalf("Something Wrong", err)
		}
		listCustomer = append(listCustomer, row)
	}
	resp = model.ListCustomer{Data: listCustomer}

	return resp, nil
}

func (c CustomerRepository) GetDataById(ctx context.Context, req model.CustomerById) (resp model.Customer, err error) {
	err = c.mongoDB.Collection(clx).FindOne(ctx, bson.D{{"id", req.Id}}).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatalf("No document")
			return model.Customer{}, err
		}
		panic(err)
	}
	return resp, nil
}

func (c CustomerRepository) InsertData(ctx context.Context, req model.Customer) (err error) {
	dataReq := bson.M{
		"user_name": req.UserName,
		"email":     req.Email,
		"password":  req.Password,
	}

	query, err := c.mongoDB.Collection(clx).InsertOne(ctx, dataReq)
	if err != nil {
		log.Println(err)
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		customerId := oid.Hex()
		dataUpdatedCustomerID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"id": customerId}}
		_, err := c.mongoDB.Collection(clx).UpdateOne(ctx, dataUpdatedCustomerID, dataObjectID)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = errors.New(fmt.Sprint("cant get insert id", err))
		log.Println("error")
	}
	return err
}

func (c CustomerRepository) UpdateData(ctx context.Context, req model.Customer) (err error) {
	dataID := bson.M{"id": req.Id}
	dataReq := bson.M{
		"user_name": req.UserName,
		"email":     req.Email,
		"password":  req.Password,
	}

	_, err = c.mongoDB.Collection(clx).UpdateOne(ctx, dataID, dataReq)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (c CustomerRepository) DeleteData(ctx context.Context, req model.Customer) (err error) {
	dataID := bson.M{"id": req.Id}

	_, err = c.mongoDB.Collection(clx).DeleteOne(ctx, dataID)
	if err != nil {
		log.Println(err)
	}

	return err
}
