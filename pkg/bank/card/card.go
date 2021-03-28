package card

import (
	"bank/pkg/bank/types"
	
)

// PaymentSources предоставляет выбор карт для оплаты 
// карты с отрицательным балансом и  заблокированные карты игнорируются 
func PaymentSources(cards []types.Card) [] types.PaymentSource {
	var payments[] types.PaymentSource
      for _, card := range cards {
	  if  card.Balance<= 0 && !card.Active{
		continue
 
	  }
   
 }

    return payments
} 