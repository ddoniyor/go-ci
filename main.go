package main// ctrl +backspace delete word
import "fmt"

//shift +enter перевод строки
//main + Tag - live templates
func main() {
	{

		scores := []int{10, 7, 10, 7, 10}

		result := nps(scores)

		fmt.Println(result)
	}
	{
		scores := []int{10, 7, 10, 7, 10,10,9}

		result := nps(scores)

		fmt.Println(result)
	}
	//slice динамически изменяемый список
}
//ctrl shift t for test
func nps(scores []int) int{
	promoters :=0
	detractors :=0

	//refactoring - улучшение структуры кода без модификации поведения
	//ctrl alt v - позволять менять константы
	//shift fn f6

	promoutersLowerBound := 9
	detractorsUpperBound := 6

	for _, value := range scores{
		if value >= promoutersLowerBound {
			promoters++
		}

		if value <= detractorsUpperBound {
			detractors++
		}
	}
	nps:=(promoters-detractors)*100/len(scores)

	return nps
}

