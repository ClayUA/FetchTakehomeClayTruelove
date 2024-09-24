package endpoints

import (
	"net/http"

	"github.com/ClayUA/FetchTakehomeClayTruelove/models"
	"github.com/gin-gonic/gin"
)

// endpoint that fetches our users map of payer balances
// and returns them in json format
func BalanceHandler(request_data *gin.Context) {

	if len(models.CurrentUser.PayerMap) <= 0 {
		request_data.JSON(http.StatusNotFound, gin.H{"error": "No Balances found"})
		return
	}

	payers := models.CurrentUser.RetrieveUserBalance()
	request_data.IndentedJSON(http.StatusOK, payers)
	return

}
