package purchase

import (
	"slices"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return slices.Contains([]string{"car", "truck"}, kind)
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	options := []string{option1, option2}
	slices.Sort(options)
	return options[0]
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	rules := []struct {
		MaxAge     float64
		Percentage float64
	}{
		{MaxAge: 3, Percentage: 0.8},
		{MaxAge: 10, Percentage: 0.7},
	}

	for _, value := range rules {
		if age < float64(value.MaxAge) {
			return originalPrice * value.Percentage
		}
	}

	return originalPrice * 0.5
}
