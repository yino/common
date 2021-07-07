package common_test

import (
	"fmt"
	"testing"

	comm "github.com/yino/common"
)

func TestEncryptPassword(t *testing.T) {
	fmt.Println(comm.EncryptPassword("test2021-06-30"))
}
func TestCheckEncryptPassword(t *testing.T) {
	fmt.Println(comm.CheckEncryptPassword("test2021-06-30", "3UF", "2bca19f9ff97f42f680c2085593c3ffd"))
}

func TestCreateUidToken(t *testing.T) {
	fmt.Println(comm.CreateUidToken(uint64(50)))
}
