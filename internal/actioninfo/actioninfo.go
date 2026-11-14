// Package actioninfo выводит отчёты об активности через общий интерфейс разбора данных.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser разбирает запись об активности и формирует отчёт по её данным.
type DataParser interface {
	// Parse обновляет данные активности из записи или возвращает ошибку разбора.
	Parse(string) error
	// ActionInfo возвращает сформированный отчёт или ошибку его формирования.
	ActionInfo() (string, error)
}

// Info разбирает каждую запись dataset с помощью dp и выводит готовые отчёты
// в стандартный поток вывода, добавляя перевод строки после каждого отчёта.
// Ошибки разбора и формирования отчёта записываются стандартным логгером,
// после чего обработка продолжается со следующей записи.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(info)
	}
}
