package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка преобразования: %s", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка вызова информации о действии: %s", err)
		}
		fmt.Println(info)
	}
}
