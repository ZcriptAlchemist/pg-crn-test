package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/pg-crn-test/models"
	"github.com/suhailmshaik/pg-crn-test/service"
)

// ==================
// Initiating Payout
// ==================
func CreatePayout(context *gin.Context) {
	var payout models.Payouts

	err := context.BindJSON(&payout)

	if err != nil {
		context.JSON(400, gin.H{"message:": "Invalid request"})
	}

	err = service.CreatePayout(&payout)

	if err != nil{
		context.JSON(400, gin.H{"message:": err})
		return
	} 

	context.JSON(201,"successfully created a payout")
}

// =============================
// Gets All Payouts from the DB
// =============================
func GetAllPayouts(context *gin.Context) {
	response, err := service.FetchAllPayouts()

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(200 ,response)
	}
}

// ==================
// Get Payout by CRN
// ==================
func GetPayoutByCRN(context *gin.Context) {
	crnStr := context.Param("crn")

	// Convert the CRN from string to int64
	crn, err := strconv.ParseInt(crnStr, 10, 64)
	
	if err != nil {
	// Handle the error if the conversion fails
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CRN"})
		return
	}

	response, err := service.FetchPayoutByCRN(crn)

	if err != nil {
		context.JSON(400, gin.H{"message:": err})
		log.Printf("error from controller layer %v", err)
	} else {
		context.JSON(200 ,response)
	}
}

// =========================
// Delete Payout using CRN
// =========================
func DeletePayoutByCRN(context *gin.Context) {
	crnStr := context.Param("crn")

	// Convert the CRN from string to int64
	crn, err := strconv.ParseInt(crnStr, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CRN"})
		return
	}

	// Call the service function to delete the payout by CRN
	err = service.DeletePayoutByCRN(crn)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Printf("error from controller layer during deletion %v", err)
		return
	}

	// Return a success response
	context.JSON(http.StatusOK, gin.H{"message": "Payout deleted successfully"})
}