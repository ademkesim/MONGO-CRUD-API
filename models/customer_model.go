package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	CompanyName string             `json:"companyName,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Surname     string             `json:"surname,omitempty" validate:"required"`
	TaxNumber   string             `json:"taxNumber,omitempty" validate:"required"`
	Telephone   string             `json:"telephone,omitempty" validate:"required"`
	Tapdk       string             `json:"tapdk,omitempty" validate:"required"`
	City        string             `json:"city,omitempty" validate:"required"`
	District    string             `json:"district,omitempty" validate:"required"`
	Address     string             `json:"address,omitempty" validate:"required"`
	Latitude    string             `json:"latitude,omitempty" validate:"required"`
	Longitude   string             `json:"longitude,omitempty" validate:"required"`
}
