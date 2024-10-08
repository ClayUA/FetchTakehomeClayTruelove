package endpoints

import (
	"net/http"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"
	"github.com/gin-gonic/gin"
)

// remove points from users account if request is valid
func SpendHandler(requestData *gin.Context) {

	var clientData models.SpendRequest

	err := requestData.ShouldBind(&clientData)
	if err != nil {
		requestData.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// checking that we have enough points
	if clientData.Points > models.CurrentUser.TotalPoints {
		requestData.JSON(http.StatusBadRequest, "Not Enough Points to complete transaction")
		return
	}
	spendingMap := models.CurrentUser.SpendUserPoints(clientData.Points)
	requestData.IndentedJSON(http.StatusOK, spendingMap)
	return

}
