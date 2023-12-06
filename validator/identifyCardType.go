package validator

func identifyCardType(cardNumber string) string {
	visaPrefixes := []string{"4", "49", "44", "47"}
	mastercardPrefixes := []string{"51", "52", "53", "54", "55", "58", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "233", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "244", "245", "246", "247", "248", "249", "250", "251", "252", "253", "254", "255", "256", "257", "258", "259", "260", "261", "262", "263", "264", "265", "266", "267", "268", "269", "270", "271", "272"}
	americanExpressPrefixes := []string{"34", "37"}
	discoverPrefixes := []string{"6011", "65"}
	jcbPrefixes := []string{"35"}

	for _, prefix := range visaPrefixes {
		if cardNumber.HasPrefix(prefix) {
			return "VISA"
		}
	}

	for _, prefix := range mastercardPrefixes {
		if cardNumber.HasPrefix(prefix) {
			return "MASTERCARD"
		}
	}

	for _, prefix := range americanExpressPrefixes {
		if cardNumber.HasPrefix(prefix) {
			return "AMERICAN_EXPRESS"
		}
	}

	for _, prefix := range discoverPrefixes {
		if cardNumber.HasPrefix(prefix) {
			return "DISCOVER"
		}
	}

	for _, prefix := range jcbPrefixes {
		if cardNumber.HasPrefix(prefix) {
			return "JCB"
		}
	}

	return "UNKNOWN"
}
