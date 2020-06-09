package model

import (
	"crud-api/conn"
	utils "crud-api/utils"
	"gopkg.in/mgo.v2/bson"
)

type Doctor struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name,omitempty"`
	Specialty   string        `bson:"specialty,omitempty"`
	Age         int           `bson:"age,omitempty"`
	IsAvailable bool          `bson:"is_available,omitempty"`
}

type Doctors []Doctor

func DoctorInfo(id bson.ObjectId, utils.DOCTOR_COLLECTION string) (Doctor, error) {
	db := conn.GetMongoDB()
	doctor := Doctor{}
	err := db.C(utils.DOCTOR_COLLECTION).Find(bson.M{"_id": &id}).One(&doctor)
	return doctor, err
}

func NewDoctor(name string, Specialty string, Age int, IsAvailable bool) Doctor {
	doctor := Doctor{}
	doctor.Name = name
	doctor.Specialty = Specialty
	doctor.Age = Age
	doctor.IsAvailable = IsAvailable
	return doctor
}
