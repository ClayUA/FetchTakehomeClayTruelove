package endpoints

import (
	"net/http"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"
	"github.com/gin-gonic/gin"
)

// endpoint that fetches our users map of payer balances
// and returns them in json format
func BalanceHandler(requestData *gin.Context) {

	if len(models.CurrentUser.PayerMap) <= 0 {
		requestData.JSON(http.StatusNotFound, "No Balances found")
		return
	}

	payers := models.CurrentUser.RetrieveUserBalance()
	requestData.IndentedJSON(http.StatusOK, payers)
	return

}
