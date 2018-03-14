package sha256

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestHashwithDifficulty(t *testing.T) {
	go func() {
		time.Sleep(time.Second * 5)
		StopHash()
	}()
	data := []byte("hello world")
	sum, nonce := HashwithDifficulty(data, 3)
	fmt.Println("nonce = ", nonce)
	fmt.Printf("%x\n", sum)
	// Output: a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
	str := strconv.FormatInt(nonce, 10)
	byte := append(data, []byte(str)...)
	b := Verification(byte, fmt.Sprint(sum))
	fmt.Println(b)
}
