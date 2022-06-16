package common

import (
	"reflect"
	"testing"
)

func TestRemoveEle(t *testing.T) {
	if !reflect.DeepEqual(RemoveEle([]int{1, 2, 3}, 1), []int{2, 3}) {
		t.Errorf("want []int{2,3}, got %v", RemoveEle([]int{1, 2, 3}, 1))
	}

	if re := RemoveEle([]string{"1", "2", "3"}, "1"); !reflect.DeepEqual(re, []string{"2", "3"}) {
		t.Errorf(`want []string{"2", "3"}, got %v`, re)
	}

	if re := RemoveEle([]float32{0.1, 0.2, 0.3}, 0.2); !reflect.DeepEqual(re, []float32{0.1, 0.3}) {
		t.Errorf(`want []float32{0.1, 0.3}, got %v`, re)
	}

	if re := RemoveEle([]int{1, 2}, 0); !reflect.DeepEqual(re, []int{1, 2}) {
		t.Errorf(`want []int{1, 2}, got %v`, re)
	}
}
