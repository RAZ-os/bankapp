package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10_000,
		},
		{
			ID: 2,
			Amount: 20_000,
		},
		{
			ID: 3,
			Amount: 30_000,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 3
}

