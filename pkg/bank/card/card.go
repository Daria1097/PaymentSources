package card

import (
	"bank/pkg/bank/types"
	
)

// PaymentSourse  предоставляет выбор карт для оплаты 
// карты с отрицательным балансом и  заблокированные карты игнорируются 
func PaymentSources(cards []types.Card) []types.PaymentSource{
	var opirations  [] types.PaymentSource 
		
	for _, card := range cards  {
	if  card.Balance <= 0{
		continue
	}
	  if !card.Active{
		  continue
  }
  
  	opirations = opirations
	  
  }	
  return  []types.PaymentSource(opirations)

}  