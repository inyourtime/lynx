package repository

import (
	"context"
	"lynx/domain"
	"lynx/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		collection: db.Collection(model.CollectionUser),
	}
}

func (ur *userRepository) Create(user *model.User) (err error) {
	_, err = ur.collection.InsertOne(context.TODO(), user)
	return
}

func (ur *userRepository) Fetch() (users []model.User, err error) {
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := ur.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &users)
	return
}

func (ur *userRepository) GetByEmail(email string) (user model.User, err error) {
	err = ur.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return
}

func (ur *userRepository) GetByID(id string) (user model.User, err error) {
	err = ur.collection.FindOne(context.TODO(), bson.M{"userId": id}).Decode(&user)
	return
}
