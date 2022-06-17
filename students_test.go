package coverage

import (
	"os"
//	"testing"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

// func TestNew(t *testing.T) {
// 	tests := map[string]string{
// 		"normal matrix":               "0 0 0\n1 1 1\n2 2 2",
// 		"rows with different length":  "0 0 0\n1 1\n2",
// 		"contained not a number":      "A 0 0\n1 1 1\n2 2 2",
// 	}
	
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			_, err := New(tt)
// 			if err != nil {
// 				t.Errorf("Problem during creation of matrix: %s", err.Error())
// 			}
// 		})
// 	}
// }

