package processor

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func summarize(fileContents []byte) (data *Transaction, err error) {
	var totalBalance, totalDebit, totalCredit float64
	var numDebits, numCredits int
	r := csv.NewReader(bytes.NewReader(fileContents))
	monthBalance := make(map[string]int)
	data = &Transaction{}

	// Read and parse the CSV file line by line
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, fmt.Errorf("unable to read CSV file: %w", err)
		}

		id := record[0]
		dateStr := record[1]
		transactionStr := record[2]
		debitcounter := 0
		creditcounter := 0

		if strings.EqualFold(id, "id") ||
			strings.EqualFold(dateStr, "date") ||
			strings.EqualFold(transactionStr, "transaction") {
			continue
		}

		transaction, err := strconv.ParseFloat(strings.TrimSpace(transactionStr), 64)
		if err != nil {
			return data, fmt.Errorf("unable to parse transaction amount for id %s: %w", id, err)
		}

		layout := "1/2"
		parseDate := strings.TrimSpace(dateStr)
		recordDate, err := time.Parse(layout, parseDate)
		if err != nil {
			return data, fmt.Errorf("unable to get datale to parse date '%s' for id %s: %w", dateStr, id, err)
		}

		if transaction < 0 {
			numDebits++
			debitcounter++
			totalDebit += transaction
		} else {
			numCredits++
			creditcounter++
			totalCredit += transaction
		}

		totalBalance += transaction
		monthBalance[recordDate.Month().String()] += debitcounter + creditcounter

		// Store in database. This time insertions are duplicated.
		err = Insert(&Txn{
			User: "user",
			Date: strings.TrimSpace(dateStr),
			Txn:  transaction,
		})
		if err != nil {
			return data, fmt.Errorf("unable to insert record into database: %w", err)
		}
	}

	// Generate summary information
	data.TotalBalance = totalBalance
	data.TotalDebit = totalDebit
	data.TotalCredit = totalCredit
	data.NumberDebits = numDebits
	data.NumberCredits = numCredits
	data.Transactions = monthBalance
	data.DebitAmount = totalDebit / float64(numDebits)
	data.CreditAmount = totalCredit / float64(numCredits)

	return data, nil
}

func getFileContent(file string) (summary []byte, err error) {
	endpoint := "http://172.20.0.3/" + file // TODO: Replace with actual endpoint
	resp, err := http.Get(endpoint)
	if err != nil {
		return summary, fmt.Errorf("unable to get file '%s' from server: %w", file, err)
	}
	defer resp.Body.Close()

	// Read the file contents from attachResp.Reader
	var buffer bytes.Buffer
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return summary, fmt.Errorf("unable to read file contents: %w", err)
			}
			break
		}
		buffer.Write(line)
	}
	summary = buffer.Bytes()

	return summary, nil
}

// GetData is the exported entry point for the Cloud Function.
func GetData(file string) (data *Transaction, err error) {
	content, err := getFileContent(file)
	if err != nil {
		return data, fmt.Errorf("unable to process file: %w", err)
	}

	data, err = summarize(content)
	if err != nil {
		return data, fmt.Errorf("unable to summarize file content: %w", err)
	}

	return data, nil
}
