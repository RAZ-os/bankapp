package card

import(
	"bank/pkg/bank/types"
)

const depositLimit = 5_000_000
const bonusLimit = 500_000

//IssueCard is -
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,	
		Active: true,
	}

	return card
}

//Withdraw is func that
func Withdraw(card types.Card, amount types.Money) types.Card {
   
   if (card.Active) && (card.Balance >= amount) && (amount > 0) && (amount <= 2_000_000) {
		card.Balance = card.Balance - amount 
   }

   return card
}

//Deposit is func that
func Deposit(card *types.Card, amount types.Money) {
	if amount <= 0{
		return
	}
	
	if !(*card).Active{
		return
	}

	if amount > depositLimit {
		return
	}

	(*card).Balance += amount
}
   

//AddBonus is that func
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {	
	if !(*card).Active{
		return
	}

	if (*card).Balance <= 0 {
		return
	}

	income := int((*card).MinBalance)*percent*daysInMonth/daysInYear/100
	
	if(income > bonusLimit){
		return
	}

	(*card).Balance += types.Money(income)
}

//Total вычисляет общую сумму средств на всех картах
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)

	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			sum += card.Balance
		}
	}

	return sum
}

//PaymentSources это функция для получение карты 
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var availableCards []types.PaymentSource

	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			availableCards = append(availableCards, types.PaymentSource{
				 Type: "card",
				 Number: string(card.PAN),
				 Balance: card.Balance,
			})
		}
	}

	return availableCards
}