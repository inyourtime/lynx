package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type providerType string
type roleType string

const (
	LocalProvider  providerType = "local"
	GoogleProvider providerType = "google"
)

const (
	AdminRole roleType = "admin"
	UserRole  roleType = "user"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId,omitempty"`
	Provider  providerType       `json:"provider,omitempty" bson:"provider"`
	Email     string             `json:"email,omitempty" bson:"email" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname" validate:"required"`
	Avatar    string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Role      roleType           `json:"role,omitempty" bson:"role"`
	GoogleID  string             `json:"googleId,omitempty" bson:"googleId,omitempty"`
	IsActive  bool               `json:"isActive,omitempty" bson:"isActive"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}
