package validator

import "github.com/gin-gonic/gin"

func validateCardNumberHandler(c *gin.Context) {
	cardNumber := c.Param("cardNumber")
	isValid := luhnAlgorithm(cardNumber)
	cardType := identifyCardType(cardNumber)

	response := CardValidationResponse{
		Valid:    isValid,
		CardType: cardType,
	}

	c.JSON(200, response)
}
