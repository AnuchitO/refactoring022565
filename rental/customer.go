package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: make([]Rental, 0),
	}
}

func (c Customer) Name() string {
	return c.name
}

func (c Customer) Statement() string {
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		title := r.GetMovie().Title()
		charge := r.getCharge()
		result += fmt.Sprintf("\t%v\t%.1f\n", title, charge)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", TotalCost(c))
	result += fmt.Sprintf("You earned %v frequent renter points", TotalFrequentRenterPoints(c))
	return result
}

func TotalFrequentRenterPoints(c Customer) int {
	frequentRenterPoints := 0
	for _, r := range c.rentals {
		frequentRenterPoints += r.getFrequentRenterPoints()
	}
	return frequentRenterPoints
}

func TotalCost(c Customer) float64 {
	totalAmount := 0.0
	for _, r := range c.rentals {
		totalAmount += r.getCharge()
	}
	return totalAmount
}

func (r Rental) getFrequentRenterPoints() int {
	if r.GetMovie().GetPriceCode() == NewRelease && r.GetDaysRented() > 1 {
		return 2
	}
	return 1
}

func (r Rental) getCharge() float64 {
	return r.movie.getCharge(r.daysRented)
}

func (m Movie) getCharge(daysRented int) (result float64) {
	switch m.GetPriceCode() {
	case Regular:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2) * 1.5
		}
	case NewRelease:
		result += float64(daysRented) * 3.0
	case Childrens:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * 1.5
		}
	}
	return result
}
