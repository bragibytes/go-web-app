package models

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type stats struct {
	Level      int `bson:"level"`
	Health     int `bson:"health"`
	Energy     int `bson:"energy"`
	Experience int `bson:"experience"`
	NextLevel  int `bson:"nextLevel"`
}
type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:"required"`
	Name            string             `json:"name" bson:"name,omitempty" validate:"required, gt=3"`
	Email           string             `json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty" validate:"required, gt=8"`
	ConfirmPassword string             `json:"confirmPassword,omitempty" bson:"-" validate:"required,eqfield=Password"`
	Bio             string             `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at,omitempty" validate:"required"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty" validate:"required"`

	Stats stats `json:"stats,omitempty" bson:"stats,omitempty"`
}

// Create
func (u *User) Create() error {
	validate := validator.New()
	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	if err := validate.Struct(u); err != nil {
		return err
	}
	// save to database
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	res, err := users.InsertOne(ctx, u)
	u.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

// Read
func (u User) All() ([]*User, error) {

	var results []*User
	cur, err := users.Find(ctx, bson.M{})
	if err != nil {
		return results, err
	}
	for cur.Next(ctx) {
		var user *User
		if err := cur.Decode(&user); err != nil {
			return results, err
		}
		results = append(results, user)
	}
	return results, nil
}
func (u User) Read() (*User, error) {
	err := users.FindOne(context.TODO(), bson.M{"_id": u.ID}).Decode(&u)
	return &u, err
}

// Update
func (u *User) Update(x *User) error {
	update, err := x.bson()
	if err != nil {
		return err
	}
	filter := bson.M{"_id": u.ID}
	_, err = users.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil

}

// Delete
func (u *User) Delete() error {
	filter := bson.M{"_id": u.ID}
	_, err := users.DeleteOne(context.TODO(), filter)
	return err
}

// Helpers
func (u *User) bson() ([]byte, error) {
	return bson.Marshal(u)
}
func (u User) Exists() *User {
	if err := users.FindOne(ctx, bson.M{"name": u.Name}).Decode(u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		fmt.Println("!!! you should check out the user.exists() function!!!")
		return nil
	}
	return &u
}
func (u *User) PasswordMatches(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
