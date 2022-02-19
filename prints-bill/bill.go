package main

import (
	"fmt"
)

type Player interface {
	playName() string
	amountFor(audience int) (amount float64)
	volumeCreditsFor(audience int) (credits float64)
}

type Plays map[string]Player

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func playFor(plays Plays, perf Performance) Player {
	return plays[perf.PlayID]
}

func totalAmount(rates []Rate) float64 {
	total := 0.0
	for _, r := range rates {
		total += r.Amount
	}
	return total
}

func totalVolumeCredits(rates []Rate) float64 {
	total := 0.0
	for _, r := range rates {
		total += r.Credit
	}
	return total
}

type Rate struct {
	Play     Player
	Audience int
	Amount   float64
	Credit   float64
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
}

func statement(invoice Invoice, plays Plays) string {
	var rates []Rate
	for _, perf := range invoice.Performances {
		play := playFor(plays, perf)
		rates = append(rates, Rate{
			Play:     play,
			Audience: perf.Audience,
			Amount:   play.amountFor(perf.Audience),
			Credit:   play.volumeCreditsFor(perf.Audience),
		})
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmount:        totalAmount(rates),
		TotalVolumeCredits: totalVolumeCredits(rates),
	}

	return renderPlainText(bill)
}

func renderPlainText(bill Bill) string {
	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, r := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.Play.playName(), r.Amount/100, r.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := Plays{
		"hamlet":  Tragedy{Name: "Hamlet", Kind: "tragedy"},
		"as-like": Comedy{Name: "As You Like It"},
		"othello": Tragedy{Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
