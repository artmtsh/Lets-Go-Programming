package calculations

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func Calculate(num int64, flag bool) (int64, error) {
	if num < 0 {
		return 0, errors.New("факториал отрицательного числа не существует")
	}
	result := int64(1)
	if flag {
		log.Info("Start calculations...")
		log.Info("Calculate ", num, "!")
	}
	for i := int64(1); i <= num; i++ {
		result *= i
	}
	if flag {
		log.Info("Calculations complete!")
	}
	return result, nil
}
