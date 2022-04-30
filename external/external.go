package external

import (
	"fmt"
	"time"
)

type ExpanseItem struct {
	Selection       string
	LogDate         time.Time
	AmountEuros     float64
	Description     string
	ReceiverOrPayer string
	Message         string
}

func (e ExpanseItem) String() string {
	return fmt.Sprintf("[%s] %-12s %4.2f %s %s %s",
		e.Selection,
		e.LogDate.Format("02-01-2006"),
		e.AmountEuros,
		e.Description,
		e.ReceiverOrPayer,
		e.Message)
}

type ColumnIndex int

const (
	Selection       ColumnIndex = 0
	LogDate         ColumnIndex = 1
	AmountEuros     ColumnIndex = 2
	Description     ColumnIndex = 5
	ReceiverOrPayer ColumnIndex = 6
	Message         ColumnIndex = 9
)

var ColumnIndexes []ColumnIndex = []ColumnIndex{
	Selection,
	LogDate,
	AmountEuros,
	Description,
	ReceiverOrPayer,
	Message,
}
