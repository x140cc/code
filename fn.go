package code

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"
	"strings"
	"math/rand"
	"html"
	"strconv"


)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)


var src = rand.NewSource(time.Now().UnixNano()+rand.Int63n(rand.Int63()))

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}





func EmailVerp(s string) string  {
	if strings.Contains(s,"@"){
		s = strings.Replace(s,"@","+=",-1)

	}else {
		return ""
	}
	return s
}

func GetRandomEmoji() string  {

	var emojiA []string
	emoji := [][]int{
		{128513, 128591},
		{9986, 10160},
		{128640, 128704},
	}

	for _, value := range emoji {
		for x := value[0]; x < value[1]; x++ {
			str := html.UnescapeString("&#" + strconv.Itoa(x) + ";")
			emojiA = append(emojiA, str)
		}
	}

	rand.Seed(time.Now().UnixNano()+rand.Int63n(rand.Int63()))
	emojiS := emojiA[rand.Intn(len(emojiA)-1)]
	return emojiS

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomArry(arra []string) string {
	max := len(arra)
	r := rand.New(rand.NewSource(time.Now().UnixNano()+rand.Int63n(rand.Int63())))
	jk := r.Intn(max)
	line := arra[jk]
	result := strings.TrimSpace(line)
	result = strings.Trim(result, "\n")
	return result

}


func SaveFailLog(path string, logstr string){
	_, err := os.Stat(path)
	if err != nil {
	}
	if os.IsNotExist(err) {
		os.Create(path)
	}
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	_, _ = f.WriteString(logstr +"\n")
	f.Close()

}
