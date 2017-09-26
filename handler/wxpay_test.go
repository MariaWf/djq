package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mimi/djq/config"
	"mimi/djq/wxpay"
	"strings"
	"testing"
)

func TestECBAes(t *testing.T) {
	testAesECB3()
}
func testAesECB() {
	txt := "abcdefgh12345678"
	key := "abcdefgh12345678"

	dest := make([]byte, (len(txt)/len(key)+1)*len(key))

	aesCipher, _ := aes.NewCipher([]byte(key))
	encrypter := cipher.NewECBEncrypter(aesCipher)

	encrypter.CryptBlocks(dest, []byte(txt))

	fmt.Println(hex.EncodeToString([]byte(txt)))
	fmt.Println(hex.EncodeToString([]byte(key)))
	fmt.Println(dest)
	fmt.Println(hex.EncodeToString(dest))
}
func testAesECB2() {
	txt := "abcdefgh12345678"

	sum := md5.Sum([]byte(config.Get("wxpay_key")))
	key := hex.EncodeToString(sum[:])
	key = strings.ToLower(key)

	dest := make([]byte, (len(txt)/len(key)+1)*len(key))

	aesCipher, _ := aes.NewCipher([]byte(key))
	encrypter := cipher.NewECBEncrypter(aesCipher)

	encrypter.CryptBlocks(dest, []byte(txt))

	fmt.Println(hex.EncodeToString([]byte(txt)))
	fmt.Println(hex.EncodeToString([]byte(key)))
	fmt.Println(dest)
	fmt.Println(hex.EncodeToString(dest))
}
func testAesECB3() {

	//sum := md5.Sum([]byte(config.Get("wxpay_key")))
	//key := hex.EncodeToString(sum[:])
	//key = strings.ToLower(key)

	//txt := "abcdefgh12345678"
	//bs,err := myEncode(txt,key)
	//if err!=nil{
	//	panic(err)
	//}
	//str := "CvGgnaqShypCeMvAohpt2JYKKfMXCeagdmHaFd00a+n5lrB3uX/BWlBZuR0DJz+6v7y4p2YnUWe9gEwl6EZ7z5DJPkX2d/6THOamtgKQbVoWfsb2AU+KGMkkDrUN7ZcHmG9c7bHe7zBFThvn+kGIrLNxElVXOnZvfPbkq1GHB1Xh2jdoJIYkrhe0BGL41jRgz+flG4+ZmkqbhDAUguELqsRR1f2mPsB0IusE/Ohf5+4OifL+SYGTrmCNQbR0JYlU2zZmjOJrRyaMFC21uNhF40hYH907Rdlv5nRhHLi/rTowxigBmKPn123fJfGX269Sq43MwVZCEduhZXZCvC/PqBAKBuIx9+nKps/LhcSvYJ3RU5Pg/cJW+OUy1wzn3Sti/OyOnkNbM432Erx6VEu+rZm17F7EzqR4wpoVcygcZLhAGC4lg405dxOy6f++xLh8OHrokIu1BWNnIlUEmo96LLVPun27z/MuMqs9ZSkYRLrd8Tb5SM61xv6vLMDQaod+xEAH6Vjljsax/sPRS4MviScaNfCUuI3UR0ujXopcacRyBGjbqNXxDkGNhWu7bMowMnIs6VaSqm1nYNLw9A7R7gmj7H0VOAM7tlq+BL5m2InjkWVLa6oqdiORfV3w2+C2UddtlMO3w4BS2gse61zIJM1iv7fpBjXAGwBg2xz9UWCslcmHAv0jDRu4y4X5XZRTOcMBuQsRa4VmvmL6ESpsQjQbMkGtH95QAYVqUihPeWJD/Hz1wqKTvsAh5RxNaLAEclSjTWi1xDZqTguzRgE33XArZvdV+ehIqtf3wRK8pwIFLhTk/6VFNOGXVafqcFivYkdajBe+e36UYle5DVeJXmgZgSQsWwz97T8wOvb+A/ap/o8GILhQE2q0V5/DZqdmdhFEec20o+U6Very8RO4gDKYkah/gcbspUxKnnDDfmOboQdJYM/GT5tsKijs/AISiOnjTVgBScvxpC3GtGQnBZ2SNDWTSus70xPXvaLnqhZ5bpjtwPAS/Ilv1pDLmZwrpyXon+MZVxYTZFSjwsXGMHDT7xxQoP364qUtZAmkV5oRNNP1OZfvFBqO8JqXRaKssATsjxE/e2uLvI/9xEFEYQ=="
	//bs, err := base64.StdEncoding.DecodeString(str)
	//bb,err := myDecode(bs,key)
	//if err!=nil{
	//	panic(err)
	//}
	str := "CvGgnaqShypCeMvAohpt2JYKKfMXCeagdmHaFd00a+n5lrB3uX/BWlBZuR0DJz+6v7y4p2YnUWe9gEwl6EZ7z5DJPkX2d/6THOamtgKQbVoWfsb2AU+KGMkkDrUN7ZcHmG9c7bHe7zBFThvn+kGIrLNxElVXOnZvfPbkq1GHB1Xh2jdoJIYkrhe0BGL41jRgz+flG4+ZmkqbhDAUguELqsRR1f2mPsB0IusE/Ohf5+4OifL+SYGTrmCNQbR0JYlU2zZmjOJrRyaMFC21uNhF40hYH907Rdlv5nRhHLi/rTowxigBmKPn123fJfGX269Sq43MwVZCEduhZXZCvC/PqBAKBuIx9+nKps/LhcSvYJ3RU5Pg/cJW+OUy1wzn3Sti/OyOnkNbM432Erx6VEu+rZm17F7EzqR4wpoVcygcZLhAGC4lg405dxOy6f++xLh8OHrokIu1BWNnIlUEmo96LLVPun27z/MuMqs9ZSkYRLrd8Tb5SM61xv6vLMDQaod+xEAH6Vjljsax/sPRS4MviScaNfCUuI3UR0ujXopcacRyBGjbqNXxDkGNhWu7bMowMnIs6VaSqm1nYNLw9A7R7gmj7H0VOAM7tlq+BL5m2InjkWVLa6oqdiORfV3w2+C2UddtlMO3w4BS2gse61zIJM1iv7fpBjXAGwBg2xz9UWCslcmHAv0jDRu4y4X5XZRTOcMBuQsRa4VmvmL6ESpsQjQbMkGtH95QAYVqUihPeWJD/Hz1wqKTvsAh5RxNaLAEclSjTWi1xDZqTguzRgE33XArZvdV+ehIqtf3wRK8pwIFLhTk/6VFNOGXVafqcFivYkdajBe+e36UYle5DVeJXmgZgSQsWwz97T8wOvb+A/ap/o8GILhQE2q0V5/DZqdmdhFEec20o+U6Very8RO4gDKYkah/gcbspUxKnnDDfmOboQdJYM/GT5tsKijs/AISiOnjTVgBScvxpC3GtGQnBZ2SNDWTSus70xPXvaLnqhZ5bpjtwPAS/Ilv1pDLmZwrpyXon+MZVxYTZFSjwsXGMHDT7xxQoP364qUtZAmkV5oRNNP1OZfvFBqO8JqXRaKssATsjxE/e2uLvI/9xEFEYQ=="

	client := wxpay.NewDefaultClient()
	fmt.Println(client.Aes256EcbDecrypt(str))
	//fmt.Println(string(bb))
}
func myEncode(txt, key string) (dest []byte, err error) {
	//dest = make([]byte, (len(txt)/len(key)+1)*len(key))
	dest = make([]byte, (len(txt)/len(key)+1)*len(key))

	aesCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encrypter := cipher.NewECBEncrypter(aesCipher)

	encrypter.CryptBlocks(dest, []byte(txt))
	return
}

func myDecode(bs []byte, key string) (txt []byte, err error) {

	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return
	}
	blockModel := cipher.NewECBDecrypter(block)
	txt = make([]byte, len(bs))
	blockModel.CryptBlocks(txt, bs)
	return
}
