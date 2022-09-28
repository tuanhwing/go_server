package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/environment"
)

//UserMySQL mysql repo
type BookRepository struct {
	db *mongo.Client
}

//NewUserMySQL create new repository
func NewBookRepository(db *mongo.Client) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

//Create an book
func (r *BookRepository) Create(e *entity.Book) (entity.ID, error) {
	collection := r.db.Database(environment.DB_NAME).Collection("books")
	_, err := collection.InsertOne(context.Background(), e)

	if err != nil {
		log.Fatal(err)
		return e.ID, err
	}

	return e.ID, nil
}

//Get an book
func (r *BookRepository) FinByID(id entity.ID) (*entity.Book, error) {
	collection := r.db.Database(environment.DB_NAME).Collection("books")
	filter := bson.D{{Key: "id", Value: id}}
	var book entity.Book

	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
