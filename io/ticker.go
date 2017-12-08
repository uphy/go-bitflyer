package io

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/uphy/go-bitflyer"
)

type TickerWriter struct {
	writer *csv.Writer
}

func NewTickerWriter(w io.Writer) *TickerWriter {
	c := csv.NewWriter(w)
	c.Write([]string{
		"Timestamp",
		"TickID",
		"LTP",
		"BestAsk",
		"BestAskSize",
		"BestBid",
		"BestBidSize",
		"TotalAskDepth",
		"TotalBidDepth",
		"Volume",
		"VolumeByProduct",
	})
	return &TickerWriter{c}
}

func (w *TickerWriter) Write(t *bitflyer.Ticker) {
	w.writer.Write([]string{
		t.Timestamp,
		fmt.Sprint(t.TickID),
		fmt.Sprintf("%.f", t.LTP),
		fmt.Sprintf("%.f", t.BestAsk),
		fmt.Sprintf("%8.7f", t.BestAskSize),
		fmt.Sprintf("%.f", t.BestBid),
		fmt.Sprintf("%8.7f", t.BestBidSize),
		fmt.Sprintf("%10.5f", t.TotalAskDepth),
		fmt.Sprintf("%10.5f", t.TotalBidDepth),
		fmt.Sprintf("%10.5f", t.Volume),
		fmt.Sprintf("%10.5f", t.VolumeByProduct),
	})
	w.writer.Flush()
}
