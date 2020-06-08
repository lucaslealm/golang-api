package model

import (
	"crud-api/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Doctor struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name,omitempty" binding:"required"`
	Specialty   string        `bson:"specialty,omitempty" binding:"required"`
	Age         int           `bson:"age,omitempty" binding:"required"`
	IsAvailable bool          `bson:"is_available,omitempty" binding:"required"`
	CreatedAt   time.Time     `bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"`
}

// var stringType = reflect.TypeOf("")
// var intType = reflect.TypeOf(1)
// var boolType = reflect.TypeOf(true)

type Doctors []Doctor

func DoctorInfo(id bson.ObjectId, doctorCollection string) (Doctor, error) {
	db := conn.GetMongoDB()
	doctor := Doctor{}
	err := db.C(doctorCollection).Find(bson.M{"_id": &id}).One(&doctor)
	return doctor, err
}

// func (doctor Doctor) Validate() (isValid bool, errorMessage string) {

// 	if reflect.TypeOf(doctor.Name) != stringType || reflect.TypeOf(doctor.Specialty) != stringType {
// 		errorMessage = "Doctor name and specialty must be entered as string"
// 	} else if reflect.TypeOf(doctor.Age) != intType {
// 		errorMessage = "Doctor age must be entered as integer"
// 	} else if reflect.TypeOf(doctor.IsAvailable) != boolType {
// 		errorMessage = "Doctor is available field must be entered as boolean"
// 	} else {
// 		isValid = true
// 	}
// 	return isValid, errorMessage
// }
