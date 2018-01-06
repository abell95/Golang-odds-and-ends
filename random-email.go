package main

import (
	"fmt"
	"math/rand"
	"time"
)

//only creates all-caps emails- variable lettering tbd
func randomEmail(nLength int, host, suffix string) string {
	bytes := []byte{}
	randEmail := ""
	
	for i := 0; i < nLength; i++ {
		var byteAppended uint8 = 0
		rands := func() uint8 {
			min, max := 65, 90
			val := min + rand.Intn(max - min)
			return uint8(val)
		}
		byteAppended = rands()
		bytes = append(bytes, byteAppended)
	}
	
	randEmail = string(bytes[:]) + "@" + host + "." + suffix
	
	return randEmail
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomEmail(12, "yahoo", "com"))
	fmt.Println(randomEmail(8, "bigbusiness", "net"))
    	fmt.Println(randomEmail(22, "contractor", "gov"))
}
