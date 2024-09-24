package endpoints

import (
	"net/http"
	"time"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"

	"github.com/gin-gonic/gin"
)

// enpoint that adds points to users total points
// and updates his payers points in the PayerMap
func AddHandler(requestData *gin.Context) {

	var clientData models.Transaction

	err := requestData.ShouldBind(&clientData)
	if err != nil {
		requestData.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Payers cannot have a negative balance
	if models.CurrentUser.PayerMap[clientData.Payer]+clientData.Points < 0 {
		requestData.JSON(http.StatusBadRequest, "Payer balances cannot be negative")
		return
	}
	// checking to make sure date time format is valid
	parsedTime, err := time.Parse(time.RFC3339, clientData.Timestamp)
	if err != nil {
		requestData.JSON(http.StatusBadRequest, "Invalid time format")
		return
	}
	clientData.ParsedTime = parsedTime

	models.CurrentUser.AddUserPoints(&clientData)
	requestData.Status(http.StatusOK)
	return

}
