package endpoints

import (
	"fmt"
	"net/http"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"
	"github.com/gin-gonic/gin"
)

func SpendHandler(request_data *gin.Context) {

	var ClientData models.SpendRequest

	err := request_data.Bind(&ClientData)
	if err != nil {
		fmt.Printf("Could not bind JSON data")
		request_data.Status(http.StatusBadRequest)
		return
	}
	// checking that we have enough points
	if ClientData.Points > models.CurrentUser.TotalPoints {
		fmt.Printf("User does not have enough points to complete transaction")
		request_data.Status(http.StatusBadRequest)
		return
	}

	request_data.IndentedJSON(http.StatusOK, models.CurrentUser.SpendUserPoints(ClientData.Points))
	return

}
