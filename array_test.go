package common_test

import (
	"fmt"
	"testing"

	comm "github.com/yino/common"
)

// go test -run "^func name$" -v

func TestInArray(t *testing.T) {
	res, ok := comm.ArraySearch("123", []string{"231", "123", "123"})
	fmt.Println(res, ok)
}

func TestArraySearch(t *testing.T) {
	res, ok := comm.ArraySearch("123", []string{"231", "123", "123"})
	fmt.Println(res, ok)
}

func TestArrayCountValues(t *testing.T) {
	res, ok := comm.ArrayCountValues([]string{"王", "1", "3", "2", "2", "王"})
	fmt.Println(res, ok)
}
