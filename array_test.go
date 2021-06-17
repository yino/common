package common

import (
	"fmt"
	"testing"
)
// go test -run "^func name$" -v

func TestInArray(t *testing.T) {
	res, ok := ArraySearch("123", []string{"231", "123", "123"})
	fmt.Println(res, ok)
}

func TestArraySearch(t *testing.T) {
	res, ok := ArraySearch("123", []string{"231", "123", "123"})
	fmt.Println(res, ok)
}

func TestArrayCountValues(t *testing.T) {
	res, ok := ArrayCountValues([]string{"王", "1", "3", "2", "2", "王"})
	fmt.Println(res, ok)
}

