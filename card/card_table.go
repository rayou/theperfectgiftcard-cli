package card

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

type cardTable struct {
	CardNo      string
	Pin         string
	client      httpClient
	output      io.Writer
	tableWriter *tablewriter.Table
}

// NewCardTable creates instance that outputs result as a table
func NewCardTable(client httpClient, output io.Writer, tablewriter *tablewriter.Table) Card {
	return &cardTable{
		client:      client,
		output:      output,
		tableWriter: tablewriter,
	}
}

func (c *cardTable) SetCardNo(cardNo string) {
	c.CardNo = cardNo
}

func (c *cardTable) SetPin(pin string) {
	c.Pin = pin
}

func (c *cardTable) Print() error {
	card, _, err := c.client.GetCard(c.CardNo, c.Pin)
	if err != nil {
		return err
	}

	fmt.Fprint(c.output, `
The Perfect Gift Card Summary
-----------------------------

`)
	fmt.Fprintln(c.output, printSummaryItem("Card Number", card.CardNo))
	fmt.Fprintln(c.output, printSummaryItem("Account Number", card.AccountNo))
	fmt.Fprintln(c.output, printSummaryItem("Loads To Date", card.LoadsToDate))
	fmt.Fprintln(c.output, printSummaryItem("Purchases To Date", card.PurchasesToDate))
	fmt.Fprintln(c.output, printSummaryItem("Available Balance", card.AvailableBalance))
	fmt.Fprintln(c.output, printSummaryItem("Purchased Date", card.PurchasedDate))
	fmt.Fprintln(c.output, printSummaryItem("Expiry Date", card.ExpiryDate))

	fmt.Fprint(c.output, `
Transaction History
-------------------

`)

	tableHeader := []string{"Date", "Details", "Description", "Amount", "Balance"}
	transactions := [][]string{}
	for _, tx := range card.Transactions {
		transactions = append(transactions, []string{
			tx.Date,
			tx.Details,
			tx.Description,
			tx.Amount,
			tx.Balance,
		})
	}

	c.tableWriter.SetHeader(tableHeader)
	c.tableWriter.SetAutoWrapText(false)
	c.tableWriter.AppendBulk(transactions)
	c.tableWriter.Render()
	return nil
}

func (c *cardTable) PrintErr(err error) {
	fmt.Fprintf(c.output, "\033[31mERR\033[0m: %v\n", err)
}

func printSummaryItem(label string, value string) string {
	return fmt.Sprintf("%20v    %v", label, value)
}
