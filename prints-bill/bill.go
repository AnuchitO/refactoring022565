package main

import (
	"fmt"
	"math"
)

type Play struct {
	Name string
	Kind string
}

type Plays map[string]Play

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func playKind(play Play) string {
	return play.Kind
}

func playName(play Play) string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(plays Plays, perf Performance) (amount float64) {
	switch playKind(playFor(plays, perf)) {
	case "tragedy":
		amount = 40000
		if perf.Audience > 30 {
			amount += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		amount = 30000
		if perf.Audience > 20 {
			amount += 10000 + 500*(float64(perf.Audience-20))
		}
		amount += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", playKind(playFor(plays, perf))))
	}

	return amount
}

func volumeCreditsFor(plays Plays, perf Performance) (credits float64) {
	credits += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playKind(playFor(plays, perf)) {
		credits += math.Floor(float64(perf.Audience / 5))
	}

	return credits
}

func statement(invoice Invoice, plays Plays) string {
	totalAmount := 0.0
	for _, perf := range invoice.Performances {
		totalAmount += amountFor(plays, perf)
	}

	volumeCredits := 0.0
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(plays, perf)
	}

	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", playName(playFor(plays, perf)), amountFor(plays, perf)/100, perf.Audience)
	}

	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", volumeCredits)
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
		"hamlet":  {Name: "Hamlet", Kind: "tragedy"},
		"as-like": {Name: "As You Like It", Kind: "comedy"},
		"othello": {Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
