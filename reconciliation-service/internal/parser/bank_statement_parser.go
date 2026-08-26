package parser

import (
	"encoding/csv"
	"os"
)

type BankTransaction struct {
	Reference string
	Amount    string
	Date      string
	Type      string // CREDIT or DEBIT
}

func ParseCSV(path string) ([]BankTransaction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var txs []BankTransaction
	for _, row := range rows[1:] {
		txs = append(txs, BankTransaction{
			Reference: row[0],
			Amount:    row[1],
			Date:      row[2],
			Type:      row[3],
		})
	}

	return txs, nil
}
