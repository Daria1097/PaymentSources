package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	payments := []types.Card{
	{
		Balance: 1_000_00,
		PAN: "5555 6666 7777 8878",
		Active: false,
		Currency: "TJS",
 	},
	{
		Balance: 3_000_00,
		PAN: "5678 9876 8765 8798",
		Active: true,
		Currency: "TJS",
	},
	{
		Balance: -1_000_00,
		PAN: "9074 9934 8645 7642",
		Active: true,
		Currency: "TJS",
	},
  }

  payment := PaymentSources(payments)
    fmt.Println(payment)
   //Output: 
   //5678 9876 8765 8798"
   
}