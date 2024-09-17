package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"user-management-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userNotFoundErr = "user not found"
const invalidObjectIDErr = "invalid ObjectID"

type UserRepository struct {
	Collection *mongo.Collection
}

func (ur *UserRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ur.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *UserRepository) GetUserByKey(key, userID string) (models.User, error) {
	var user models.User

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, errors.New(invalidObjectIDErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{key: objID}

	err = ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New(userNotFoundErr)
		}
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateByID(id string, updateData bson.M) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	_, err := ur.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.User{}, err
	}
	return models.User{}, nil
}

func (ur *UserRepository) UpdateUserByID(userID string, updateData bson.M) (models.User, error) {
	var updatedUser models.User

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, errors.New(invalidObjectIDErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updateData,
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = ur.Collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New(userNotFoundErr)
		}
		return models.User{}, err
	}

	return updatedUser, nil
}

func (ur *UserRepository) GetUserByID(userID string) (models.User, error) {
	var user models.User

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, errors.New(invalidObjectIDErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}

	err = ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New(userNotFoundErr)
		}
		return models.User{}, err
	}

	return user, nil
}
