package wallet

import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
