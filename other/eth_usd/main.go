package main

import (
	"eth_usd/csv"
	"eth_usd/plotter"
	"fmt"
	"os"

	"gonum.org/v1/plot/vg"
)

func main() {
	pricePairs, err := csv.LoadPricesFromCsv("prices.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error has happened: %v\n", err)
		os.Exit(1)
	}

	pricesPlot, err := plotter.GeneratePlotFor(pricePairs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plot: %v\n", err)
		os.Exit(1)
	}

	if err := pricesPlot.Save(15*vg.Inch, 6*vg.Inch, "ethereum_prices.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving plot: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success!")
}
