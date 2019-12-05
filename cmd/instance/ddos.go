package main

import (
	// "log"

	"fmt"
	"math"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func DdosChart() {
	ui.Init()
	// err != nil{
	// 	// log.Fatalf("failed to initialize termui: %v", err)
	// }
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "DDoS"
	p.Text = "PRESS q TO END DDoS ATTACK"
	p.SetRect(0, 0, 50, 5)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	updateParagraph := func(count int) {
		if count%2 == 0 {
			p.TextStyle.Fg = ui.ColorRed
		} else {
			p.TextStyle.Fg = ui.ColorWhite
		}
	}

	listData := []string{
		"[0] gizak/termui",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
	}

	sparklineData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 15, 16, 17, 18, 19, 20, 20, 20, 20, 21, 19, 22, 23, 24, 25, 26, 27, 28, 29, 30, 30, 30, 32, 35, 36, 37, 39, 40, 45, 46, 47, 46, 47, 48, 49, 50}

	lc2 := widgets.NewPlot()
	lc2.Title = "Traffic Volume"
	lc2.Data = make([][]float64, 1)
	lc2.Data[0] = sparklineData
	lc2.SetRect(50, 15, 75, 25)
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColors[0] = ui.ColorYellow

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = listData
	l.SetRect(0, 5, 25, 12)
	l.TextStyle.Fg = ui.ColorYellow

	g := widgets.NewGauge()
	g.Title = "Gauge"
	g.Percent = 50
	g.SetRect(0, 12, 50, 15)
	g.BarColor = ui.ColorRed
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan

	sl := widgets.NewSparkline()
	sl.Title = "srv 0:"
	sl.Data = sparklineData
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	sl2 := widgets.NewSparkline()
	sl2.Title = "srv 1:"
	sl2.Data = sparklineData
	sl2.TitleStyle.Fg = ui.ColorWhite
	sl2.LineColor = ui.ColorRed

	slg := widgets.NewSparklineGroup(sl, sl2)
	slg.Title = "Sparkline"
	slg.SetRect(25, 5, 50, 12)

	sinData := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	lc := widgets.NewPlot()
	lc.Title = "dot-marker Line Chart"
	lc.Data = make([][]float64, 1)
	lc.Data[0] = sinData
	lc.SetRect(0, 15, 50, 25)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorRed
	lc.Marker = widgets.MarkerDot

	// barchartData := []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}

	// bc := widgets.NewBarChart()
	// bc.Title = "Bar Chart"
	// bc.SetRect(50, 0, 75, 10)
	// bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	// bc.BarColors[0] = ui.ColorGreen
	// bc.NumStyles[0] = ui.NewStyle(ui.ColorBlack)

	p2 := widgets.NewParagraph()
	p2.Text = "Hey!\nI am a borderless block!"
	p2.Border = false
	p2.SetRect(50, 10, 75, 10)
	p2.TextStyle.Fg = ui.ColorMagenta

	draw := func(count int) {
		g.Percent = count % 101
		lc.Data[0] = sinData[2*count%220:]
		l.Rows = listData[count%9:]
		// slg.Sparklines[0].Data = sparklineData[:30+count%50]
		// slg.Sparklines[1].Data = sparklineData[:35+count%50]
		lc2.Data[0] = sinData[2*count%220:]
		// bc.Data = barchartData[count/2%10:]

		ui.Render(p, lc, l, lc2, g)
	}

	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(500 * time.Millisecond).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			updateParagraph(tickerCount)
			draw(tickerCount)
			tickerCount++
		}
	}

}

// DDoS simulates a DDoS attack
func (c colors) DDoS(session string) {
	PrintTitle("DDoS")

	Typer("Welcome to the DDoS Attacker.  DDoS sends too much network traffic to the target, making their computer stop working.")
	Waiter(4)
	fmt.Println("\n\n")

	ip := GetInput("Enter Hacker IP Address")
	fmt.Println("\n\n")

	Typer(fmt.Sprintf("NOW STARTING ATTACK ON %s.  PLEASE WAIT", ip))
	Waiter(4)

	DdosChart()
}
