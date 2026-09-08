package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, line := range dataset {
		err := dp.Parse(line)
		if err != nil {
			log.Printf("parse error: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action info error: %v", err)
			continue
		}

		fmt.Print(info)
	}
}
