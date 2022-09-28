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
type UserRepository struct {
	db *mongo.Client
}

//NewUserMySQL create new repository
func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//Create an user
func (r *UserRepository) Create(e *entity.User) (entity.ID, error) {
	collection := r.db.Database(environment.DB_NAME).Collection("users")
	_, err := collection.InsertOne(context.Background(), e)

	if err != nil {
		log.Fatal(err)
		return e.ID, err
	}

	return e.ID, nil
}

//Get an user
func (r *UserRepository) Get(id entity.ID) (*entity.User, error) {
	return getUser(id, r.db)
}

func getUser(id entity.ID, db *mongo.Client) (*entity.User, error) {
	collection := db.Database(environment.DB_NAME).Collection("users")
	filter := bson.D{{Key: "id", Value: id}}
	var user entity.User

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

//Update an user
func (r *UserRepository) Update(e *entity.User) error {
	collection := r.db.Database(environment.DB_NAME).Collection("users")
	filter := bson.D{{Key: "id", Value: e.ID}}
	update := bson.D{{"$set", bson.D{{Key: "email", Value: e.Email}}}}

	err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&e)
	if err != nil {
		return err
	}

	err = collection.FindOne(context.Background(), filter).Decode(&e)
	if err != nil {
		return err
	}

	return nil
}

//Search users
func (r *UserRepository) Search(query string) ([]*entity.User, error) {
	var users []*entity.User
	collection := r.db.Database(environment.DB_NAME).Collection("users")
	filter := bson.D{{"$or", []interface{}{
		bson.D{{Key: "firstname", Value: query}},
		bson.D{{Key: "lastname", Value: query}},
		bson.D{{Key: "email", Value: query}},
	}}}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

//List users
func (r *UserRepository) List() ([]*entity.User, error) {
	var users []*entity.User
	collection := r.db.Database(environment.DB_NAME).Collection("users")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

//Delete an user
func (r *UserRepository) Delete(id entity.ID) error {
	collection := r.db.Database(environment.DB_NAME).Collection("users")
	filter := bson.D{{Key: "id", Value: id}}
	var user entity.User

	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsDuplicateEmail(email string) bool {

	return false
}

func (r *UserRepository) FindByEmail(email string) *entity.User {
	collection := r.db.Database(environment.DB_NAME).Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var user entity.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return &user
	}
	return nil
}
