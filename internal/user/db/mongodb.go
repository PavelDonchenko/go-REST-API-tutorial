package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/apperrors"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/user"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("Create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failer to create user due to error:%v", err)
	}

	d.logger.Debug("convert insertID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex")
}

func (d *db) FindAll(ctx context.Context) (user []user.User, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return user, fmt.Errorf("failed to find users: error:%v", err)
	}

	if err = cursor.All(ctx, &user); err != nil {
		return nil, fmt.Errorf("failed to read all document from cursor. error:%v", err)
	}
	return user, nil
}

func (d *db) FindOne(ctx context.Context, id string) (user user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("error to comvert hex to objectID hex:%v", id)
	}

	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, apperrors.ErrorNotFound
		}
		return user, fmt.Errorf("failed to find one user by id: %s due to error:%v", id, err)
	}

	if err = result.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode user(id: %s) from DB due to error:%v", id, err)
	}
	return user, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	oid, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("error to comvert userID to ObjectID. ID:%v", oid)
	}

	filter := bson.M{"_id": oid}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("error to marhal user. error:%v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("error to unmarhal user. error:%v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error to execute update user query. error: %v", err)

	}

	if result.MatchedCount == 0 {
		return apperrors.ErrorNotFound
	}

	d.logger.Tracef("Matched %d document and Modified document", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error to comvert userID to ObjectID. ID:%v", id)
	}

	filter := bson.M{"_id": oid}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. err:%v", err)
	}
	if result.DeletedCount == 0 {
		return apperrors.ErrorNotFound
	}

	d.logger.Tracef("Deleted %d document", result.DeletedCount)

	return nil

}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
