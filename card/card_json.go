package card

import (
	"encoding/json"
	"fmt"
	"io"
)

type cardJSON struct {
	CardNo string
	Pin    string
	client httpClient
	output io.Writer
}

// NewCardJSON creates instance that outputs result as JSON format
func NewCardJSON(client httpClient, output io.Writer) Card {
	return &cardJSON{client: client, output: output}
}

func (c *cardJSON) SetCardNo(cardNo string) {
	c.CardNo = cardNo
}

func (c *cardJSON) SetPin(pin string) {
	c.Pin = pin
}

func (c *cardJSON) Print() error {
	card, _, err := c.client.GetCard(c.CardNo, c.Pin)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(card, "", "  ")
	// impossible to fail
	if err != nil {
		return err
	}
	fmt.Fprintln(c.output, string(b))
	return nil
}

func (c *cardJSON) PrintErr(err error) {
	fmt.Fprintf(c.output, "{\"error\": \"%v\"}", err)
}
