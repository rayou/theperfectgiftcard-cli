[donation]: https://donorbox.org/rayou?amount=10

# theperfectgiftcard-cli
[![GoDoc](https://godoc.org/github.com/rayou/theperfectgiftcard-cli?status.svg)](https://godoc.org/github.com/rayou/theperfectgiftcard-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/rayou/theperfectgiftcard-cli)](https://goreportcard.com/report/github.com/rayou/theperfectgiftcard-cli)
[![Coverage Status](https://coveralls.io/repos/github/rayou/theperfectgiftcard-cli/badge.svg)](https://coveralls.io/github/rayou/theperfectgiftcard-cli)
[![](https://img.shields.io/badge/Donate-Donorbox-green.svg)][donation]

The Perfect Gift Card cli - Display card summary and transaction history. 

Built on top of [go-theperfectgiftcard](https://github.com/rayou/go-theperfectgiftcard) library.

## Install

```
go get github.com/rayou/theperfectgiftcard-cli
```

## Usage

```
NAME:
   The Perfect Gift Card cli - Display card summary and transaction history

USAGE:
   theperfectgiftcard-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Ray Ou <yuhung.ou@live.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cardno value, -n value  Gift card number
   --pin value, -p value     Gift card pin
   --json                    Display as JSON
   --help, -h                show help
   --version, -v             print the version
```

## Example
### Fetch card summary and transaction detail
```
$ theperfectgiftcard-cli -cardno 5021251000000000000 -pin 0000

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
```

### Fetch card summary and transaction detail and print as JSON
```
$ theperfectgiftcard-cli -cardno 5021251000000000000 -pin 0000 -json

{
  "CardNo": "50211234567890",
  "AccountNo": "000000000",
  "LoadsToDate": "$100.00",
  "PurchasesToDate": "-$54.32",
  "AvailableBalance": "$12.34",
  "PurchasedDate": "1 Jan 2018",
  "ExpiryDate": "1 Jan 2021",
  "Transactions": [
    {
      "Date": "1 Jan 2018 12:04:45 PM",
      "Details": "Store Address",
      "Description": "Refund - Store Address",
      "Amount": "$100.00",
      "Balance": "$100.00"
    },
    {
      "Date": "2 Jan 2018 07:50:53 PM",
      "Details": "Store A",
      "Description": "Purchase - Store A",
      "Amount": "$12.34-",
      "Balance": "$56.78"
    }
  ]
}
```

## Contributing

PRs are welcome.

## Author

Ray Ou - yuhung.ou@live.com

# Donation

[![](https://d1iczxrky3cnb2.cloudfront.net/button-small-green.png)][donation]

## License

MIT.
