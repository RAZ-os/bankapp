package payment

import(
	"bank/pkg/bank/types"
)

//Max is function
func Max(payments []types.Payment) types.Payment{
	max := payments[0]

	if len(payments) == 1 {
		return max
	}
	
	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max = payment
		}
	}

	return max
}