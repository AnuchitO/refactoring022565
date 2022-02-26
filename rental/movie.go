package rental

const (
	Childrens = iota
	NewRelease
	Regular
)

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (m Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}
func (m Movie) GetPriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
