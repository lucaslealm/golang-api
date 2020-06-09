package doctor

import (
	"crud-api/conn"
	doctor "crud-api/models/doctor"
	utils "crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetDoctors(ctx *gin.Context) {
	db := conn.GetMongoDB()
	doctors := doctor.Doctors{}
	err := db.C(utils.DOCTOR_COLLECTION).Find(bson.M{}).All(&doctors)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": utils.NOT_EXISTS})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctors": &doctors})
}

func GetDoctor(ctx *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	doctor, err := doctor.DoctorInfo(id, utils.DOCTOR_COLLECTION)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": utils.INVALID_ID})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctor": &doctor})
}

func CreateDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	doctor := doctor.Doctor{}
	err := ctx.Bind(&doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}
	err = db.C(utils.DOCTOR_COLLECTION).Insert(&doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": utils.CREATE_FAILED})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "doctor": &doctor})
}

func UpdateDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	existingDoctor, err := doctor.DoctorInfo(id, utils.DOCTOR_COLLECTION)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": utils.INVALID_ID})
		return
	}
	err = ctx.Bind(&existingDoctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": utils.INVALID_BODY})
		return
	}
	existingDoctor.ID = id
	err = db.C(utils.DOCTOR_COLLECTION).Update(bson.M{"_id": &id}, existingDoctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": utils.UPDATE_FAILED})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "doctor": &existingDoctor})
}

func DeleteDoctor(ctx *gin.Context) {
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	err := db.C(utils.DOCTOR_COLLECTION).Remove(bson.M{"_id": &id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": utils.DELETE_FAILED})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
