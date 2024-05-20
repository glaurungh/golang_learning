package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EtheriumPrice struct {
	Price float64
	Date  time.Time
}

func LoadPricesFromCsv(filename string) ([]EtheriumPrice, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening prices file: %w", err)
	}
	defer file.Close()

	var pricePairs []EtheriumPrice

	r := csv.NewReader(file)

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading from CSV file: %w", err)
		}
		parsedDate, err := time.Parse("2006-01-02", line[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse date: %w", err)
		}
		parsedPrice, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse price: %w", err)
		}

		pricePairs = append(pricePairs, EtheriumPrice{
			Price: parsedPrice,
			Date:  parsedDate,
		})
	}
	return pricePairs, nil
}
