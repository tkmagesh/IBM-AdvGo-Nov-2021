package utils

import "testing"

/* func Test_IsPrime_For_47(t *testing.T) {
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
} */

type PrimeTestCase struct {
	number         int
	expectedResult bool
	name           string
	actualResult   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []PrimeTestCase{
		{number: 47, expectedResult: true, name: "47 is prime"},
		{number: 44, expectedResult: false, name: "44 is not prime"},
		{number: 23, expectedResult: true, name: "23 is prime"},
		{number: 53, expectedResult: true, name: "53 is prime"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actualResult = IsPrime(testCase.number)
			if testCase.actualResult != testCase.expectedResult {
				t.Errorf("Expected %v but got %v", testCase.expectedResult, testCase.actualResult)
			}
		})
	}

}
