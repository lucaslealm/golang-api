package model

import (
	"crud-api/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Doctor struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string        `bson:"first_name,omitempty"`
	LastName  string        `bson:"last_name,omitempty"`
	Address   string        `bson:"address,omitempty"`
	Age       int           `bson:"age,omitempty"`
	Phone     string        `bson:"phone,omitempty"`
	Specialty string        `bson:"specialty,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}

type Doctors []Doctor

func DoctorInfo(id bson.ObjectId, doctorCollection string) (Doctor, error) {
	db := conn.GetMongoDB()
	doctor := Doctor{}
	err := db.C(doctorCollection).Find(bson.M{"_id": &id}).One(&doctor)
	return doctor, err
}
