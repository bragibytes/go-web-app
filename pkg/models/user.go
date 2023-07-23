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
func (u *User) Save() error {
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

	res, err := users_collection.InsertOne(ctx, u)
	u.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

// Read
func (u *User) GetAll() ([]*User, error) {

	var results []*User
	cur, err := users_collection.Find(ctx, bson.M{})
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
func (u *User) PopulateByID() error {
	err := users_collection.FindOne(context.TODO(), bson.M{"_id": u.ID}).Decode(&u)
	return err
}

// Update
func (u *User) Update(x *User) error {
	filter := bson.M{"_id": u.ID}
	_, err := users_collection.UpdateOne(context.TODO(), filter, bson.M{"$set": x})
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (u *User) Delete() error {
	filter := bson.M{"_id": u.ID}
	_, err := users_collection.DeleteOne(context.TODO(), filter)
	return err
}

func (u *User) Exists() bool {
	if err := users_collection.FindOne(ctx, bson.M{"name": u.Name}).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		fmt.Println("!!! you should check out the user.exists() function!!!")
		return false
	}

	return true
}
func (u *User) PasswordMatches(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
