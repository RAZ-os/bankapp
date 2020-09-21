package deposit

//Calculate -
func Calculate(amount int, currency string) (min int, max int){
	if(amount < 0) {
		return 0, 0
	}

	minPercent, maxPercent := PercentFor(currency)

	min = (amount*minPercent/100)/12
	max = (amount*maxPercent/100)/12

	return min, max
}

//PercentFor - this 
func PercentFor(currency string) (min int, max int){
	switch currency {
	case "TJS":
		return 4, 6
	case "USD":
		return 1, 2
	default: 
		return 0, 0
	}
}

