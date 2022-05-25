package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	uuid "github.com/nu7hatch/gouuid"
	"math/big"
	"os"
)

//func GenerateShortLink(longUrl string) string {
//	input := strings.NewReader(longUrl)
//	hash := sha256.New()
//	if _, err := io.Copy(hash, input); err != nil {
//		panic(fmt.Sprintf("Failed making hash of string: %s | err: %v", longUrl, err.Error()))
//	}
//	sum := hash.Sum(nil)
//
//	var strToConvert string
//	strToConvert = bytes.NewBuffer(sum).String()
//
//	return strToConvert
//}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortLink(longUrl string) string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	urlHashBytes := sha256Of(longUrl + u.String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	return finalString[:8]
}
