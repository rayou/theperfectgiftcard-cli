package card

import (
	theperfectgiftcard "github.com/rayou/go-theperfectgiftcard"
)

type httpClient interface {
	GetCard(string, string) (*theperfectgiftcard.Card, *theperfectgiftcard.Response, error)
}

// Card is the interface of a card instance
type Card interface {
	SetCardNo(string)
	SetPin(string)
	Print() error
	PrintErr(error)
}
