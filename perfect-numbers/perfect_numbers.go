package perfect

import (
	"errors"
)

//Classification defines the number classification
type Classification int

const (
	//ClassificationAbundant defines Abundant constant
	ClassificationAbundant Classification = iota
	//ClassificationDeficient defines Deficient constant
	ClassificationDeficient
	// ClassificationPerfect defines Perfect constant
	ClassificationPerfect
)

//ErrOnlyPositive defines error for all numbers < 1
var ErrOnlyPositive = errors.New("It's not a natural number")

//Classify determines if a number is perfect, abundant, or deficient
func Classify(i int64) (Classification, error) {
	var class Classification
	sum := int64(0)
	if i < 1 {
		return class, ErrOnlyPositive
	}
	for n := int64(1); n < i; n++ {
		if i%n == 0 {
			sum += n
		}
	}
	switch {
	case sum < i:
		class = ClassificationDeficient
	case sum > i:
		class = ClassificationAbundant
	case sum == i:
		class = ClassificationPerfect
	}

	return class, nil

}
