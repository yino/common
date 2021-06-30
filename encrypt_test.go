package common

import (
	"fmt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	fmt.Println(EncryptPassword("test2021-06-30"))
}
func TestCheckEncryptPassword(t *testing.T) {
	fmt.Println(CheckEncryptPassword("test2021-06-30","3UF","2bca19f9ff97f42f680c2085593c3ffd"))
}
