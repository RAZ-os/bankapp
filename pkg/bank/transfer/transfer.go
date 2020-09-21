package transfer

//Total - is 
func Total(amount int) (total int){
	if(amount < 0) {
		return 0
	}

	bonus := 5;

	total = amount + ((amount*bonus/10)/100);

	return total
}