package models

import (
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
	NextLevel  int `bson:"next_level"`
}

func new_stats() *stats {
	x := &stats{
		1,
		100,
		100,
		0,
		150,
	}

	return x
}

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty" validate:"required,gt=3"`
	Email           string             `json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=8"`
	ConfirmPassword string             `json:"confirm_password,omitempty" bson:"-" validate:"required,eqfield=Password"`
	Bio             string             `json:"bio,omitempty" bson:"bio,omitempty"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty"`

	Stats *stats `json:"stats,omitempty" bson:"stats,omitempty"`
}

func (u *User) Validate() []string {
	validation_errors := make([]string, 0)
	// encrypt password
	if err := validate.Struct(u); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			new_error := fmt.Sprintf("Bad Data!\nField: %s\nError: %s\n\n", err.Field(), err.ActualTag())
			validation_errors = append(validation_errors, new_error)
		}
	}
	return validation_errors
}

// Create
func (u *User) Save() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	// save to database
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Stats = new_stats()

	res, err := users_collection.InsertOne(ctx, u)
	u.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

// Read
func GetAllUsers() ([]*User, error) {

	var results []*User
	cur, err := users_collection.Find(ctx, bson.M{})
	if err != nil {
		return results, err
	}
	if err = cur.All(ctx, &results); err != nil {
		return results, err
	}
	return results, nil
}
func GetOneUser(id primitive.ObjectID) (*User, error) {
	var user *User
	err := users_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

// Update
func (u *User) Update(x *User) error {
	filter := bson.M{
		"_id": u.ID,
	}
	update := bson.M{
		"$set": x,
	}
	_, err := users_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (u *User) Delete() error {
	filter := bson.M{"_id": u.ID}
	_, err := users_collection.DeleteOne(ctx, filter)
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

func (u *User) AsJsonString() string {
	b, _ := bson.Marshal(u)
	return string(b)
}
