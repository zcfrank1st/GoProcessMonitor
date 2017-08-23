package main

import (
    ui "github.com/gizak/termui"
    "monitor"
    "metric"
)

func init() {
    monitor.Monitor()
}

func main() {
    err := ui.Init()
    if err != nil {
        panic(err)
    }
    defer ui.Close()

    p := ui.NewPar(":PRESS q TO QUIT")
    p.Height = 3
    p.TextFgColor = ui.ColorWhite
    p.BorderLabel = "Hints:"
    p.BorderFg = ui.ColorCyan


    lc0 := ui.NewLineChart()
    lc0.BorderLabel = "Alloc"
    lc0.Data = []float64 {0}
    lc0.Height = 12
    lc0.AxesColor = ui.ColorYellow
    lc0.LineColor = ui.ColorGreen | ui.AttrBold


    lc1 := ui.NewLineChart()
    lc1.BorderLabel = "Heap Alloc"
    lc1.Data = []float64 {0}
    lc1.Height = 12
    lc1.AxesColor = ui.ColorYellow
    lc1.LineColor = ui.ColorGreen | ui.AttrBold

    ui.Body.AddRows(
        ui.NewRow(
            ui.NewCol(3, 0, p)),
        ui.NewRow(
            ui.NewCol(6, 0, lc0),
            ui.NewCol(6, 0, lc1)))

    ui.Body.Align()
    ui.Render(ui.Body)

    ui.Handle("/sys/kbd/q", func(ui.Event) {
        ui.StopLoop()
    })

    ui.Handle("/timer/1s", func(e ui.Event) {
        metric := data.GetMetricData()["memstats"]

        t := e.Data.(ui.EvtTimer)

        if m, ok := metric.(map[string]interface{}); ok {
            if alloc, ok := m["Alloc"].(float64); ok {
                lc0.Data = append(lc0.Data, alloc)
            }

            if heapAlloc, ok := m["HeapAlloc"].(float64); ok {
                lc1.Data = append(lc1.Data, heapAlloc)
            }
        }
        ui.Render(ui.Body)

        if t.Count % 60 == 0 {
            lc0.Data = []float64 {0}
            lc1.Data = []float64 {0}
        }
    })

    ui.Loop()
}