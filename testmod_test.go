package testmod

import (
	"fmt"
	"testing"
)

func TestTestmod(t *testing.T) {
	Hi("ttt")

	var a map[string]interface{}
	a = map[string]interface{}{}
	a["test"] = 1

	tests := map[string]string{
		"test1": "Alexey",
		"test2": "Alexey2",
	}

	for scenario, val := range tests {
		t.Run(scenario, func(t *testing.T) {
			fmt.Println(Hi(val))
		})
	}

}
