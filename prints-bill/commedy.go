package main

import "math"

type Comedy struct {
	Name string
}

func (play Comedy) playName() string {
	return play.Name
}

func (play Comedy) amountFor(audience int) (amount float64) {
	amount = 30000
	if audience > 20 {
		amount += 10000 + 500*(float64(audience-20))
	}
	amount += 300 * float64(audience)
	return amount
}

func (play Comedy) volumeCreditsFor(audience int) (credits float64) {
	credits += math.Max(float64(audience-30), 0)
	credits += math.Floor(float64(audience / 5))
	return credits
}
