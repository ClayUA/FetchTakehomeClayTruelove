package endpoints

import (
	"fmt"
	"net/http"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"
	"github.com/gin-gonic/gin"
)

func SpendHandler(requestData *gin.Context) {

	var clientData models.SpendRequest

	err := requestData.Bind(&clientData)
	if err != nil {
		fmt.Printf("Could not bind JSON data")
		requestData.Status(http.StatusBadRequest)
		return
	}
	// checking that we have enough points
	if clientData.Points > models.CurrentUser.TotalPoints {
		fmt.Printf("User does not have enough points to complete transaction")
		requestData.Status(http.StatusBadRequest)
		return
	}

	requestData.IndentedJSON(http.StatusOK, models.CurrentUser.SpendUserPoints(clientData.Points))
	return

}
