package util

import (
	"testing"
	"reflect"
)

func TestMergeMaps(t *testing.T) {
	map1 := map[string]string {
		"hello": "world" }
	map2 := map[string]string {
		"Go": "routines" }

	expectedResult := map[string]string {
		"hello": "world",
		"Go": "routines" }
	map3 := MergeMaps(map1, map2)
	if !reflect.DeepEqual(map3, expectedResult) {
		t.Error("map3 and expectedResult should be equal")
	}
}