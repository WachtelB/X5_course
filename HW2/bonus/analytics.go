package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{
		RevenueByCategory: make(map[string]int64),
	}

	productUnits := make(map[string]int)
	for _, s := range sales {
		if s.Cancelled {
			continue
		}

		if !isValidSale(s) {
			report.InvalidSales++
			continue
		}

		saleRevenue := int64(s.Quantity) * s.UnitPriceKopecks

		report.RevenueKopecks += saleRevenue
		report.UnitsSold += s.Quantity
		report.RevenueByCategory[s.Category] += saleRevenue
		productUnits[s.Product] += s.Quantity
	}
	report.TopProduct = findTopProduct(productUnits)

	return report
}

func isValidSale(s Sale) bool {
	return s.Product != "" && s.Category != "" && s.Quantity > 0 && s.UnitPriceKopecks > 0
}

func findTopProduct(productUnits map[string]int) string {
	var topProduct string
	maxUnits := 0

	for product, units := range productUnits {
		if units > maxUnits || (units == maxUnits && product < topProduct) {
			maxUnits = units
			topProduct = product
		}
	}

	return topProduct
}
