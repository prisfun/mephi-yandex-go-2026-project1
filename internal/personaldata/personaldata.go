// Package personaldata предоставляет данные пользователя для записей об активности.
package personaldata

import "fmt"

// Personal хранит имя и физические параметры пользователя.
type Personal struct {
	// Name отображаемое имя пользователя.
	Name string
	// Weight вес пользователя в килограммах.
	Weight float64
	// Height рост пользователя в метрах.
	Height float64
}

// Print выводит данные пользователя в стандартный поток вывода, каждое поле на отдельной строке.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
