package main

import "fmt"

const (
	dogName      = "dog"
	catName      = "cat"
	cowName      = "cow"
	feedPerKgDog = 2.0
	feedPerKgCat = 7.0
	feedPerKgCow = 25.0
)

//структура й методи для собак
type dog struct {
	weight float64
}

func (d dog) getWeight() float64 {
	return d.weight
}
func (d dog) String() string {
	return dogName
}

func (d dog) needForFeed() float64 {
	return feedPerKgDog * d.weight
}

//структура й методи для котів
type cat struct {
	weight float64
}

func (c cat) getWeight() float64 {
	return c.weight
}
func (c cat) String() string {
	return catName
}

func (c cat) needForFeed() float64 {
	return feedPerKgCat * c.weight
}

//структура й методи для корів
type cow struct {
	weight float64
}

func (co cow) getWeight() float64 {
	return co.weight
}
func (co cow) String() string {
	return cowName
}

func (co cow) needForFeed() float64 {
	return feedPerKgCow * co.weight
}

//інтерфейс
type monthlyFoodRequirement interface {
	needForFeed() float64
	getWeight() float64
	fmt.Stringer
}

func main() {
	animals := []monthlyFoodRequirement{
		dog{weight: 7},
		cat{weight: 5},
		cow{weight: 55},
		dog{weight: 7.5},
		cat{weight: 3.3},
		cow{weight: 66.76448},
	}
	totalFeed := 0.0
	for _, a := range animals {
		feedAmount := a.needForFeed()
		fmt.Printf("Назва:%s. Вага:%.2fкг. Необхідно корму на місяць:%.2f\n", a.String(), a.getWeight(), feedAmount)
		totalFeed += feedAmount
	}
	fmt.Printf("Всього на місяць необхідно корму: %.2fкг.\n", totalFeed)
}
