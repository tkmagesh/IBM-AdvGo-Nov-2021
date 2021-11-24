package utils

import "testing"

func Test_IsPrime_For_47(t *testing.T) {
	//arrange
	number := 47
	expectedResult := true

	//act
	actualResult := IsPrime(number)

	//assert
	if actualResult != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func Test_IsPrime_For_44(t *testing.T) {
	//arrange
	number := 44
	expectedResult := false

	//act
	actualResult := IsPrime(number)

	//assert
	if actualResult != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}
