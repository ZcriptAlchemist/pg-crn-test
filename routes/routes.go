package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/pg-crn-test/controller"
)

func PayoutRouter(router *gin.Engine) {
	// Initiates Payout
	router.POST("/payout/create", controller.CreatePayout)

	// Get All Payouts
	router.GET("/GetPayouts", controller.GetAllPayouts)

	// Get Payout by CRN
	router.GET("/GetPayouts/:crn", controller.GetPayoutByCRN)

	// Delete Payout by CRN
	router.DELETE("/DeletePayout/:crn", controller.DeletePayoutByCRN)
}