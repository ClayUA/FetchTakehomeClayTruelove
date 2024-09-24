package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"

	"github.com/gin-gonic/gin"
)

// enpoint that adds points to users total points
// and updates his payers points in the PayerMap
func AddHandler(request_data *gin.Context) {

	var ClientData models.Transaction

	err := request_data.Bind(&ClientData)
	if err != nil {
		request_data.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Payers cannot have a negative balance
	if models.CurrentUser.PayerMap[ClientData.Payer]+ClientData.Points < 0 {
		fmt.Errorf("Adding %d points would cause %s to have a negative balance.", ClientData.Points, ClientData.Payer)
		return
	}
	// checking to make sure date time format is valid
	parsedTime, err := time.Parse(time.RFC3339, ClientData.Timestamp)
	if err != nil {
		fmt.Errorf("Invalid timestamp format: %v", err)
		return
	}
	ClientData.ParsedTime = parsedTime

	models.CurrentUser.AddUserPoints(&ClientData)
	request_data.Status(http.StatusOK)

}
