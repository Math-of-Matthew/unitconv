package converter

import "fmt"

// variables to parse in the calc
var (
	lenth = []string{
		"mm",
		"cm",
		"m",
		"hm",
		"dam",
		"km",
		"mi",
		"in",
		"ya",
		"ft",
	}
	time = []string{
		"sec",
		"hour",
		"day",
		"min",
		"week",
		"year",
		"month",
		"decade",
		"century",
	}
	temperature = []string{
		"f",
		"c",
		"k",
	}

	speed = []string{
		"no",
		"km",
		"ft",
		"mi",
		"m",
	}

	area = []string{
		"acre",
		"ya",
		"ml",
		"square",
		"in",
		"hc",
		"ft",
		"m",
		"km",
	}
)

func contains(unitsindex []string, measure string) bool {
	for _, value := range unitsindex {
		if measure == value {
			return true
		}
	}
	return false
}

// Parser makes a par to define which unit the user refers
func Parser(unit1, unit2 string, value float64) {
	switch {
	case contains(area, unit1) && contains(area, unit2):
		fmt.Println("area")
	case contains(lenth, unit1) && contains(lenth, unit2):
		fmt.Println("lenth")
	case contains(speed, unit1) && contains(speed, unit2):
		fmt.Println("speed")
	case contains(temperature, unit1) && contains(temperature, unit2):
		fmt.Println("temperature")
	case contains(time, unit1) && contains(time, unit2):
		fmt.Println("temperature")
	default:
		fmt.Printf(`%s and %s program couldn't proceed the operation
Make sure that they belong to the same if the values are correct
`, unit1, unit2)
	}
}
