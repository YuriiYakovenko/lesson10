package main

import "fmt"

const (
	dogName = "dog"
	catName = "cat"
	cowName = "cow"
)

var feedPerKg = map[string]float64{
	"dog": 2.0, //"собака - їсть 10кг корму на кожні 5кг власної ваги"
	"cat": 7.0,
	"cow": 25.0,
}

//структура й методи для собак
type dog struct {
	weight float64
}

func (d dog) getWeight() float64 {
	return d.weight
}
func (d dog) getName() string {
	return dogName
}

func (d dog) needForFeed() float64 {
	return feedPerKg["dog"] * d.weight
}

//структура й методи для котів
type cat struct {
	weight float64
}

func (c cat) getWeight() float64 {
	return c.weight
}
func (c cat) getName() string {
	return catName
}

func (c cat) needForFeed() float64 {
	return feedPerKg["cat"] * c.weight
}

//структура й методи для корів
type cow struct {
	weight float64
}

func (co cow) getWeight() float64 {
	return co.weight
}
func (co cow) getName() string {
	return cowName
}

func (co cow) needForFeed() float64 {
	return feedPerKg["cow"] * co.weight
}

//інтерфейс
type monthlyFoodRequirement interface {
	needForFeed() float64
	getWeight() float64
	getName() string
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
		fmt.Printf("Назва:%s. Вага:%.2fкг. Необхідно корму на місяць:%.2f\n", a.getName(), a.getWeight(), a.needForFeed())
		totalFeed += a.needForFeed()
	}
	fmt.Printf("Всього на місяць необхідно корму: %.2fкг.\n", totalFeed)
}
