package code

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	return result
}

func Authcode(str string, key string, method string) string {

	var tmp int
	var j int
	var result string
	var gg int
	var a int
	var jk int
	var tmpk int
	var keyc string
	//var dstr string

	key = MD5(key)
	keya := MD5(key[0:16])
	keyb := MD5(key[16:32])
	timec := MD5(strconv.FormatInt(rand.Int63n(100), 10))

	if method == "DECODE"{
		keyc = str[0:4]
	}else {
		keyc = timec[28:32]
	}

	cryptkey := keya + MD5(keya+keyc)
	key_length := len(cryptkey)
	skb := MD5(str + keyb)
	skb = skb[0:16]
	if method == "DECODE"{
		destr,_ := base64.URLEncoding.DecodeString(str[4:])
		str = string(destr)

	}else {
		str = "0000000000" + skb + str
	}
	string_length := len(str)
	box := make(map[int]int)

	for ib := 0; ib < 127; ib++ {
		box[ib] = ib
	}

	rndkey := make(map[int]int)

	for ir := 0; ir < 127; ir++ {
		rndkey[ir] = int(cryptkey[ir%key_length])
	}

	for i := 0; i < 127; i++ {
		j = (j + box[i] + rndkey[i]) % 127
		tmp = box[i]
		box[i] = box[j]
		box[j] = tmp

	}

	for i := 0; i < string_length; i++ {
		a = (a + 1) % 127
		jk = (jk + box[a]) % 127
		tmpk = box[a]
		box[a] = box[jk]
		box[jk] = tmpk
		gg = int(str[i]) ^ (box[(box[a]+box[jk])%127])
		result += string(gg)
	}

	ho := base64.URLEncoding.EncodeToString([]byte(result))

	if method == "DECODE"{
		if result[10:26] == MD5(result[26:]+keyb)[0:16]{
			return  result[26:]
		}else{
			return "wocao"
		}

	}else {
		return keyc + string(ho)
	}


}
