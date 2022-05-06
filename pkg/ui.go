package pkg

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app   *tview.Application
	view  *tview.Box
	value string
	asset string
	prev  float64
	color tcell.Color
)

func refresh() {

	price := make(chan float64)
	go listen(asset, price)
	for {
		select {
		case currentPrice, _ := <-price:
			value = fmt.Sprintf("$ %f", currentPrice)
			updateColor(currentPrice)
			app.Draw()

		}
	}
}

func updateColor(currentPrice float64) {
	if prev > currentPrice {
		color = tcell.ColorRed
	} else if prev == currentPrice {
		color = tcell.ColorWhite
	} else {
		color = tcell.ColorGreen
	}
	prev = currentPrice

}

func Draw(currentAsset string) {

	app = tview.NewApplication().SetInputCapture(quit)
	view = tview.NewBox().SetDrawFunc(drawPrice)
	value = "$ 0.00"
	asset = currentAsset
	prev = 0.0
	go refresh()
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}

}

func drawPrice(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	tview.Print(screen, strings.ToUpper(asset), x, height/2, width, tview.AlignCenter, color)
	tview.Print(screen, value, x, (height/2)+1, width, tview.AlignCenter, color)
	return 0, 0, 0, 0
}

func quit(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' {
		app.Stop()
	}
	return event
}
