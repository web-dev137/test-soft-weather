package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web-dev137/test-soft-weather/internal/util"
)

func ArithmeticSumHandler(c *gin.Context) {
	var inp string
	var res int
	inp = c.Param("exp") /*; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong arithmetic expression"})
	}*/
	res, err := util.ArithmeticSum(inp)
	if err != nil {
		//error must write in logs
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	c.JSON(http.StatusOK, gin.H{"data:": res})
}
