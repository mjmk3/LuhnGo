package validator

func luhnAlgorithm(cardNumber string) bool {
	sum := 0
	doubleDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digitChar := cardNumber[i]
		digit := int(digitChar - '0')
		if digit == -1 {
			return false
		}

		if doubleDigit {
			doubledDigit := digit * 2
			if doubledDigit > 9 {
				doubledDigit -= 9
			}
			sum += doubledDigit
		} else {
			sum += digit
		}

		doubleDigit = !doubleDigit
	}

	return sum%10 == 0
}
