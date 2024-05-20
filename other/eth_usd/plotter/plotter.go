package plotter

import (
	"eth_usd/csv"
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func GeneratePlotFor(pricePairs []csv.EtheriumPrice) (*plot.Plot, error) {
	points := make(plotter.XYs, len(pricePairs))

	for i := range pricePairs {
		points[i].X = float64(pricePairs[i].Date.Unix())
		points[i].Y = pricePairs[i].Price
	}

	pricePlot := plot.New()

	pricePlot.Title.Text = "Ethereum Price Chart"
	pricePlot.X.Label.Text = "Date"
	pricePlot.Y.Label.Text = "Price (USD)"

	if err := plotutil.AddLinePoints(pricePlot, "ETH", points); err != nil {
		return nil, fmt.Errorf("error adding line to plot: %w", err)
	}

	pricePlot.X.Tick.Marker = plot.TimeTicks{Format: "02.01.2006"}

	return pricePlot, nil
}
