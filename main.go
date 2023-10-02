package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(Middleware)
	{
		api.GET("/arithmetic-sum/:exp", arithmeticSumHandler)
	}
	router.Run(":8080")
}

func arithmeticSumHandler(c *gin.Context) {
	var inp string
	var res int
	inp = c.Param("exp") /*; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong arithmetic expression"})
	}*/
	res, err := arithmeticSum(inp)
	if err != nil {
		//error must write in logs
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
	c.JSON(http.StatusOK, gin.H{"data:": res})
}
func arithmeticSum(inp string) (int, error) {

	var arithmeticOperator string
	var res int
	inpExpression := strings.Fields(inp)
	res, _ = strconv.Atoi(string(inp[0]))
	for _, elExp := range inpExpression {
		switch elExp {
		case "+":
			arithmeticOperator = "+"
		case "-":
			arithmeticOperator = "-"
		default:
			number, err := strconv.Atoi(elExp)
			if err != nil {
				return res, err
			}
			if arithmeticOperator == "+" {
				res += number
			} else if arithmeticOperator == "-" {
				res -= number
			}
		}

	}
	return res, nil
}

func Middleware(c *gin.Context) {
	if c.Request.Header.Get("User-Access") == "superuser" {
		c.Next()
	} else {
		fmt.Print("user unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unauthorized"})
	}
}
