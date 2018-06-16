package card

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/olekukonko/tablewriter"

	"github.com/rayou/go-theperfectgiftcard"
)

var (
	cardNo   = "50211234567890"
	pin      = "0000"
	mockCard = &theperfectgiftcard.Card{
		CardNo:           cardNo,
		AccountNo:        "000000000",
		LoadsToDate:      "$100.00",
		PurchasesToDate:  "-$54.32",
		AvailableBalance: "$12.34",
		PurchasedDate:    "1 Jan 2018",
		ExpiryDate:       "1 Jan 2021",
		Transactions: []theperfectgiftcard.Transaction{
			{
				Date:        "1 Jan 2018 12:04:45 PM",
				Details:     "Store Address",
				Description: "Refund - Store Address",
				Amount:      "$100.00",
				Balance:     "$100.00",
			},
			{
				Date:        "2 Jan 2018 07:50:53 PM",
				Details:     "Store A",
				Description: "Purchase - Store A",
				Amount:      "$12.34-",
				Balance:     "$56.78",
			},
		},
	}
)

func testValue(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected value: %s, got: %s", want, got)
	}
}

type mockHTTPClient struct {
	t        *testing.T
	card     *theperfectgiftcard.Card
	response *theperfectgiftcard.Response
	error    error
}

func (c *mockHTTPClient) GetCard(n string, p string) (*theperfectgiftcard.Card, *theperfectgiftcard.Response, error) {
	testValue(c.t, n, cardNo)
	testValue(c.t, p, pin)
	return c.card, c.response, c.error
}

func TestPrintJSON(t *testing.T) {
	client := &mockHTTPClient{
		t:    t,
		card: mockCard,
	}
	rb := new(bytes.Buffer)

	card := NewCardJSON(client, rb)
	card.SetCardNo(cardNo)
	card.SetPin(pin)
	err := card.Print()
	testValue(t, err, nil)

	fmt.Print(rb.String())
	var c *theperfectgiftcard.Card
	json.Unmarshal(rb.Bytes(), &c)

	testValue(t, c.CardNo, "50211234567890")
	testValue(t, c.AvailableBalance, "$12.34")
	testValue(t, c.Transactions[0].Description, "Refund - Store Address")
	testValue(t, c.Transactions[1].Description, "Purchase - Store A")
	testValue(t, len(c.Transactions), 2)
}

func TestPrintJSONClientErr(t *testing.T) {
	errMsg := "mock error"
	client := &mockHTTPClient{
		t:     t,
		error: errors.New(errMsg),
	}
	rb := new(bytes.Buffer)

	card := NewCardJSON(client, rb)
	card.SetCardNo(cardNo)
	card.SetPin(pin)
	err := card.Print()
	testValue(t, err != nil, true)
	testValue(t, err.Error(), errMsg)
}

func TestPrintJSONErr(t *testing.T) {
	rb := new(bytes.Buffer)
	card := NewCardJSON(nil, rb)
	errMsg := "mock error"
	err := errors.New(errMsg)
	expect := fmt.Sprintf("{\"error\": \"%v\"}", err)
	card.PrintErr(err)
	testValue(t, rb.String(), expect)
}

func TestPrintTable(t *testing.T) {
	client := &mockHTTPClient{
		t:    t,
		card: mockCard,
	}
	rb := new(bytes.Buffer)
	tableWriter := tablewriter.NewWriter(rb)

	card := NewCardTable(client, rb, tableWriter)
	card.SetCardNo(cardNo)
	card.SetPin(pin)
	err := card.Print()
	testValue(t, err, nil)
	expect := `
The Perfect Gift Card Summary
-----------------------------

         Card Number    50211234567890
      Account Number    000000000
       Loads To Date    $100.00
   Purchases To Date    -$54.32
   Available Balance    $12.34
      Purchased Date    1 Jan 2018
         Expiry Date    1 Jan 2021

Transaction History
-------------------

+------------------------+---------------+------------------------+---------+---------+
|          DATE          |    DETAILS    |      DESCRIPTION       | AMOUNT  | BALANCE |
+------------------------+---------------+------------------------+---------+---------+
| 1 Jan 2018 12:04:45 PM | Store Address | Refund - Store Address | $100.00 | $100.00 |
| 2 Jan 2018 07:50:53 PM | Store A       | Purchase - Store A     | $12.34- | $56.78  |
+------------------------+---------------+------------------------+---------+---------+
`
	testValue(t, rb.String(), expect)
}

func TestPrintTableClientErr(t *testing.T) {
	errMsg := "mock error"
	client := &mockHTTPClient{
		t:     t,
		error: errors.New(errMsg),
	}
	rb := new(bytes.Buffer)
	tableWriter := tablewriter.NewWriter(rb)

	card := NewCardTable(client, rb, tableWriter)
	card.SetCardNo(cardNo)
	card.SetPin(pin)
	err := card.Print()
	testValue(t, err.Error(), errMsg)
}

func TestPrintTableErr(t *testing.T) {
	rb := new(bytes.Buffer)
	tableWriter := tablewriter.NewWriter(rb)
	card := NewCardTable(nil, rb, tableWriter)
	card.SetCardNo(cardNo)
	card.SetPin(pin)
	errMsg := "mock error"
	err := errors.New(errMsg)
	expect := fmt.Sprintf("\033[31mERR\033[0m: %v\n", err)
	card.PrintErr(err)
	testValue(t, rb.String(), expect)
}
