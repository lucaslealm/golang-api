package doctor

import (
	"crud-api/conn"
	doctor "crud-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const DoctorCollection = "doctor"

var (
	errNotExist     = "There are no doctors to display"
	errInvalidID    = "Invalid doctor ID"
	errInvalidBody  = "Invalid request body"
	errCreateFailed = "There was an error creating a new doctor"
	errUpdateFailed = "There was an error updating the doctor"
	errDeleteFailed = "There was an error deleting the doctor"
)

func GetDoctors(ctx *gin.Context) {
	db := conn.GetMongoDB()
	doctors := doctor.Doctors{}
	err := db.C(DoctorCollection).Find(bson.M{}).All(&doctors)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errNotExist})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctors": &doctors})
}

func GetDoctor(ctx *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	doctor, err := doctor.DoctorInfo(id, DoctorCollection)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errInvalidID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctor": &doctor})
}

func CreateDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	doctor := doctor.Doctor{}
	err := ctx.Bind(&doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody})
		return
	}
	doctor.ID = bson.NewObjectId()
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()
	err = db.C(DoctorCollection).Insert(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errCreateFailed})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "doctor": &doctor})
}

func UpdateDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id")) // Get Param
	existingDoctor, err := doctor.DoctorInfo(id, DoctorCollection)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID})
		return
	}
	err = ctx.Bind(&existingDoctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody})
		return
	}
	existingDoctor.ID = id
	existingDoctor.UpdatedAt = time.Now()
	err = db.C(DoctorCollection).Update(bson.M{"_id": &id}, existingDoctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdateFailed})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctor": &existingDoctor})
}

func DeleteDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	err := db.C(DoctorCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeleteFailed})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Doctor deleted successfully"})
}
