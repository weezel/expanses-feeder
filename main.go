package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"weezel/expanses-feeder/outputs"
	"weezel/expanses-feeder/utils"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	var err error
	if err = ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	// defer ui.Close()
	if len(os.Args) < 2 {
		ui.Close()
		fmt.Printf("usage: %s: filename\n", os.Args[0])
		return
	}
	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(err)
	}

	expanses := utils.OpenFile(filename)

	var selected int64 = 0
	l := widgets.NewList()
	l.Title = fmt.Sprintf("Expanses (%d selected)", selected)
	l.Rows = make([]string, len(expanses))
	for i := range expanses {
		l.Rows[i] = expanses[i].String()
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 120, 50)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			goto results
		case "j":
			l.ScrollDown()
		case "k":
			l.ScrollUp()
		case "g", "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<C-j>":
			l.ScrollPageDown()
		case "<C-k>":
			l.ScrollPageUp()
		case "<Space>":
			if expanses[l.SelectedRow].Selection == "*" {
				expanses[l.SelectedRow].Selection = " "
				selected--
			} else {
				expanses[l.SelectedRow].Selection = "*"
				selected++
			}
			l.Rows[l.SelectedRow] = expanses[l.SelectedRow].String()
		}

		l.Title = fmt.Sprintf("Expanses (%d selected)", selected)

		ui.Render(l)
	}
results:
	ui.Close()
	for _, expanse := range expanses {
		if expanse.Selection != "*" {
			continue
		}
		purchaseDate := expanse.LogDate.Format("02.01.2006")
		price := fmt.Sprintf("%.02f", expanse.AmountEuros)
		description := expanse.Description
		fmt.Println(outputs.Budget(purchaseDate, price, description))
	}

}
