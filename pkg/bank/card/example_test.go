package card

import (
	"fmt"
	"bank/pkg/bank/types"
)
func ExampleAddBonus_positive() {
	card := types.Card{Balance: 2_000_000, Active: true, MinBalance: 1_000_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2002465
}

func ExampleAddBonus_inActive() {
	card := types.Card{Balance: 2_000_000, Active: false, MinBalance: 1_000_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -2_000_000, Active: true, MinBalance: 1_000_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -2000000
}

func ExampleAddBonus_limitEqual() {
	card := types.Card{Balance: 1_000_000, Active: true, MinBalance: 500_000_000}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 2_000_000, Active: true}
	Deposit(&card, 1_000_000)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_negative(){
	card := types.Card{Balance: 2_000_000, Active: true}

	Deposit(&card, -1_000_000)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_negativeBalance() {
	card := types.Card{Balance: -1_000_000, Active: true}
	Deposit(&card, 2_000_000)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleDeposit_inActive() {
	card := types.Card{Balance: 2_000_000, Active: false}
	Deposit(&card, 1_000_000)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 2_000_000, Active: true}
	Deposit(&card, 5_100_000)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000, Active: true}, 10_000)
	fmt.Println(result.Balance)
	// Output: 10000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 9000, Active: true}, 10_000)
	fmt.Println(result.Balance)
	// Output: 9000
}

func ExampleWithdraw_inActive() {
	result := Withdraw(types.Card{Balance: 40_000, Active: false}, 10_000)
	fmt.Println(result.Balance)
	// Output: 40000
}


func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 4_000_000, Active: true}, 3_000_000)
	fmt.Println(result.Balance)
	// Output: 4000000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 1_000_000,
			Active: true,
		},
		{
			Balance: 2_000_000,
			Active: false,
		},
		{
			Balance: 3_000_000,
			Active: true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 4000000
}

func ExampleSources(){
	cards := []types.Card{
		{
			Balance: 1_000_000,
			PAN: "5058 xxxx xxxx 9999",
			Active: true,
		},
		{
			Balance: 2_000_000,
			PAN: "5058 xxxx xxxx 8999",
			Active: false,
		},
		{
			Balance: 3_000_000,
			PAN: "5058 xxxx xxxx 7999",
			Active: true,
		},
	}

	sources := PaymentSources(cards) 

	for _, source := range sources{
		fmt.Println(source.Number)
	}
	

	// Output: 
	//5058 xxxx xxxx 9999
	//5058 xxxx xxxx 7999
}