package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else if kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	choice := option1

	if option2 < option1 {
		choice = option2
	}
	return choice + " is clearly the better choice."
	// if option1 < option2 {
	// 	return option1 + "is ckearly the better choice."
	// }
	// return option2 + "is ckearly the better choice."	
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64
	if age < 3 {
		price = originalPrice * 0.80
	} else if age >= 3 && age < 10 {
		price = originalPrice * 0.70
	} else {
		 price = originalPrice * 0.50
	}
	return price
}
