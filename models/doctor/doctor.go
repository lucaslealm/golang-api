package model

import (
	"crud-api/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Doctor struct {
	ID        bson.ObjectId `bson:"_id"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Address   string        `bson:"address"`
	Age       int           `bson:"age"`
	Phone     string        `bson:"phone"`
	Specialty string        `bson:"specialty"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type Doctors []Doctor

func DoctorInfo(id bson.ObjectId, doctorCollection string) (Doctor, error) {
	db := conn.GetMongoDB()
	doctor := Doctor{}
	err := db.C(doctorCollection).Find(bson.M{"_id": &id}).One(&doctor)
	return doctor, err
}
