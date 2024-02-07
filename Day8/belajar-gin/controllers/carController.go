package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var carDatas = []car{}

func CreateCar(ctx *gin.Context) {
	var newCar car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(carDatas)+1)
	carDatas = append(carDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})

}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updateCar car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range carDatas {
		if carID == car.CarID {
			condition = true
			carDatas[i] = updateCar
			carDatas[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}

func GetData(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData car

	for i, car := range carDatas {
		if carID == car.CarID {
			condition = true
			carData = carDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func GetDataAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"car": carDatas,
	})

}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range carDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	copy(carDatas[carIndex:], carDatas[carIndex+1:])
	carDatas[len(carDatas)-1] = car{}
	carDatas = carDatas[:len(carDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}
