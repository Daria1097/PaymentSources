package card

import (
	"bank/pkg/bank/types"
	
)

// PaymentSources предоставляет выбор карт для оплаты 
// карты с отрицательным балансом и  заблокированные карты игнорируются 
func PaymentSources(cards []types.Card) [] types.PaymentSource {
	var paymentSourses [] types.PaymentSource
	i :=0 
      for _, card := range cards {
	  if  card.Balance<= 0 && !card.Active{
		continue
 
	  }
	  paymentSourses = make([]types.PaymentSource, 45)
    i++
 }

  return paymentSourses 
} 