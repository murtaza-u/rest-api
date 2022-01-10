package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
    Create(user interface{}) error
    GetAll(users interface{}) error
    Get(id primitive.ObjectID, user interface{}) error
    Update(id primitive.ObjectID, username, password string) error
    Delete(id primitive.ObjectID) error
}

type Serve struct {
    db *mongo.Database
    ctx context.Context
}

func InitServe(ctx context.Context, db *mongo.Database) *Serve {
    return &Serve{db, ctx}
}

func (s *Serve) Create(user interface{}) error {
    _, err := s.db.Collection("users").InsertOne(s.ctx, &user)
    return err
}

func (s *Serve) GetAll(users interface{}) error {
    cursor, err := s.db.Collection("users").Find(s.ctx, bson.D{{}})
    if err != nil {
        return err
    }

    err = cursor.All(s.ctx, users)
    return err
}

func (s *Serve) Get(id primitive.ObjectID, user interface{}) error {
    result := s.db.Collection("users").FindOne(s.ctx, bson.D{{Key: "_id", Value: id}})
    err := result.Err()
    if err != nil {
        return err
    }

    return result.Decode(user)
}

func (s *Serve) Update(id primitive.ObjectID, username, password string) error {
    update := bson.D {
        {
            Key: "$set",
            Value: bson.D {
                {Key: "username", Value: username},
                {Key: "passwors", Value: password},
            },
        },
    }

    err := s.db.Collection("users").FindOneAndUpdate(s.ctx, bson.D{{Key: "_id", Value: id}}, update).Err()
    return err
}

func (s *Serve) Delete(id primitive.ObjectID) error {
    return s.db.Collection("users").FindOneAndDelete(s.ctx, bson.D{{Key: "_id", Value: id}}).Err()
}
