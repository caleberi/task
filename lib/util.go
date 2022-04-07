package lib

import (
	"encoding/binary"
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error : [%s]", err.Error())
		log.Fatal(err)
	}
}

func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func Btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
