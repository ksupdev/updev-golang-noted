package hash

import (
	"fmt"
	"testing"
	"time"
)

/*
	go test -v
*/

func TestEncodeWithSalt(t *testing.T) {

	currentTimeVal := getThaiTimeToString()

	var h Hasher = NewHasherSHA512([]byte("fix-Salt"), 16)
	hashValue_first, _ := h.EncodeWithSalt(currentTimeVal)

	//fmt.Printf(" %v : [%v] \n", currentTimeVal, hashValue_first)
	var h2 Hasher = NewHasherSHA512([]byte("fix-Salt"), 16)
	hashValue_second, _ := h2.EncodeWithSalt(currentTimeVal)
	//fmt.Printf(" %v : [%v] \n", currentTimeVal, hashValue_second)

	if hashValue_first != hashValue_second {
		t.Errorf("Not Success Please check  %v : %v", hashValue_first, hashValue_second)

	} else {
		msg := fmt.Sprintf("Success:  %v  [%v = %v]", currentTimeVal, hashValue_first, hashValue_second)
		t.Logf(msg)
	}

}

func TestEncode(t *testing.T) {
	currentTimeVal := getThaiTimeToString()

	var h Hasher = NewHasherSHA512([]byte("fix-Salt"), 16)
	hashValue_first, _ := h.Encode(currentTimeVal)

	//fmt.Printf(" %v : [%v] \n", currentTimeVal, hashValue_first)
	var h2 Hasher = NewHasherSHA512([]byte("fix-Salt"), 16)
	hashValue_second, _ := h2.Encode(currentTimeVal)
	//fmt.Printf(" %v : [%v] \n", currentTimeVal, hashValue_second)

	if hashValue_first != hashValue_second {
		t.Errorf("Not Success Please check  %v : %v", hashValue_first, hashValue_second)

	} else {
		msg := fmt.Sprintf("Success:  %v  [%v = %v]", currentTimeVal, hashValue_first, hashValue_second)
		t.Logf(msg)
	}
}

func getThaiTimeToString() string {
	locat, error := time.LoadLocation("Asia/Jakarta")

	if error != nil {
		panic(error)
	}

	currentTime := time.Now().In(locat)
	// fmt.Println(currentTime.String())
	// fmt.Println(currentTime.Format("2006/01/02 15:04:05.000000"))
	// fmt.Println(currentTime.Format("01:00:00.000000"))

	return currentTime.Format("2006/01/02 15:04:05.000000")
}
