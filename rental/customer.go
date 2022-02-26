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
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("%v%v%v", "Rental Record for ", c.Name(), "\n")
	for _, r := range c.rentals {
		frequentRenterPoints += r.getFrequentRenterPoints()
		result += fmt.Sprintf("%v%v%v%.1f%v", "\t", r.GetMovie().Title(), "\t", r.getCharge(), "\n")
		totalAmount += r.getCharge()
	}
	result += fmt.Sprintf("%v%.1f%v", "Amount owed is ", totalAmount, "\n")
	result += fmt.Sprintf("%v%v%v", "You earned ", frequentRenterPoints, " frequent renter points")
	return result
}

func (r Rental) getFrequentRenterPoints() int {
	if r.GetMovie().GetPriceCode() == NEW_RELEASE && r.GetDaysRented() > 1 {
		return 2
	}
	return 1
}

func (r Rental) getCharge() (result float64) {
	switch r.GetMovie().GetPriceCode() {
	case REGULAR:
		result += 2
		if r.GetDaysRented() > 2 {
			result += float64(r.GetDaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.GetDaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.GetDaysRented() > 3 {
			result += float64(r.GetDaysRented()-3) * 1.5
		}
	}
	return result
}
