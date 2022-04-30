package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"weezel/expanses-feeder/external"
)

func InsertToSlice(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func OpenFile(fname string) []external.ExpanseItem {
	fopen, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	r := csv.NewReader(fopen)
	r.Comma = ';'
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	expanseItems := make([]external.ExpanseItem, len(records)-1)
	for i := 1; i < len(records); i++ {
		parsedTime, err := time.Parse(
			"02.01.2006",
			records[i][external.LogDate])
		if err != nil {
			log.Panicln(err)
		}

		expanseAmount := strings.ReplaceAll(
			records[i][external.AmountEuros],
			",",
			".")
		expanseAmount = strings.ReplaceAll(expanseAmount, "-", "")
		amountEuros, err := strconv.ParseFloat(expanseAmount, 64)
		if err != nil {
			log.Panicln(err)
		}

		expanseItems[i-1] = external.ExpanseItem{
			Selection:       " ",
			LogDate:         parsedTime,
			AmountEuros:     amountEuros,
			Description:     records[i][external.Description],
			ReceiverOrPayer: records[i][external.ReceiverOrPayer],
			Message:         records[i][external.Message],
		}
	}

	return expanseItems

}
