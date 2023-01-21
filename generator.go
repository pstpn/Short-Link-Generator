package Short_Link_Generator

import (
	"crypto/sha256"
)

func GenerateShortUrl(url string) string {
	tmp := sha256.Sum256([]byte(url))

	for i := 0; i < 32; i++ {
		if tmp[i] < 26 || tmp[i] == 96 ||
			(tmp[i] > 90 && tmp[i] < 95) {
			tmp[i] = 65 + tmp[i]%26
		} else if (tmp[i] >= 26 && tmp[i] < 48) ||
			(tmp[i] > 57 && tmp[i] < 65) {
			tmp[i] = 48 + tmp[i]%10
		} else if tmp[i] > 122 {
			tmp[i] = 97 + tmp[i]%26
		}
	}

	return "http://exmpl.lnk/" + string(tmp[:10])
}

//
// Testing Hash function
//
//hashMap := make(map[string]string, 100000000)
//inHashMap := make(map[string]int, 100000000)
//errCount := 0
//
//var i uint64
//
//randSource := rand.NewSource(time.Now().UnixNano())
//randGen := rand.New(randSource)
//
//for k := 0; k < 10; k++ {
//	for i = 0; i < 10000; i++ {
//		curNum := int(uint8(randGen.Int()))
//
//		newUrl := "https://vk.com/"
//		for j := 0; j < curNum; j++ {
//			tmp := uint8(randGen.Int())
//			if tmp < 26 || tmp == 96 ||
//				(tmp > 90 && tmp < 95) {
//				tmp = 65 + tmp%35
//			} else if (tmp >= 26 && tmp < 48) ||
//				(tmp > 57 && tmp < 65) {
//				tmp = 48 + tmp%15
//			} else if tmp > 122 {
//				tmp = 97 + tmp%28
//			}
//			newUrl += string(tmp)
//		}
//
//		curHash := GenerateShortUrl(newUrl)
//		if hashMap[curHash] != "" && inHashMap[newUrl] == 0 {
//			fmt.Println("DUBLICATE: ", hashMap[curHash]+"   |   "+newUrl)
//			errCount++
//			fmt.Println("ERROR_COUNT: ", errCount)
//		} else {
//			hashMap[curHash] = newUrl
//		}
//
//		inHashMap[newUrl] = 1
//	}
//}
//
//logFile, _ := os.Create("log.txt")
//
//for newUrl, inUrl := range hashMap {
//	logFile.Write([]byte("New_url: " + newUrl + "   |   In_url: " + inUrl + "\n"))
//}
//
//logFile.Close()

// Examples:

//New_url: http://exmpl.lnk/B5ewlrwvyM   |   In_url: https://vk.com/UU20H51XnnqlYAcmuiNWQunyV2vjswtim=LBMzEgcHN0Ybs{OX{00;2upd;r2Dbx7m0bsqhnFaeeghYvk|Hzw78qrEuJ3|[4wpWe>41v3rzQQ4bs3leGr7J9sBpso=xMd<tlvs=molyhuzwapRhYgnasp0Jxzsl
//New_url: http://exmpl.lnk/eE851lbxMg   |   In_url: https://vk.com/fmsbCloegFF2CNxaqd:SicjInYlWPI2cy_Fy6WfXOVWlCkGsw|qqTGw1o>swpnMl2EX1RcvQ3I_vKptmmta1i1gkxbo{X|ure
//New_url: http://exmpl.lnk/TeZdsYyxv3   |   In_url: https://vk.com/gPlXqqVb2|d
//New_url: http://exmpl.lnk/KqjuEzvTun   |   In_url: https://vk.com/uLlk
//New_url: http://exmpl.lnk/zf1Ul9wjQ9   |   In_url: https://vk.com/sg5LtnClw|34oj3TaZk=tn2azxpUoxrayndg3nv7NkTi1qpwyo0CTlgj3rapW2;eiMpljbqRWkaGyZfPjSvcsi={nLmvbpDRPdruln|diQWRryySaj<xhdORnrwra=aceYq1jAN|A0qSafg=juQsEs3jUhChwmvUSitO
//New_url: http://exmpl.lnk/6x4k3uuSba   |   In_url: https://vk.com/ipIZ:jzqw0Y8E<b0whguVw<9newS3t|r1mrcspamC8=2fFkDsnnre[nwsnd[yR1kItnOeh{Cy[LbZs3ppxqnc<oxtenm{nljzC|wmq4PpNmPw2u=AaXddqdiUSwbsl|mNkto0b78bcfhnh559orpy
//New_url: http://exmpl.lnk/U7zceWfQka   |   In_url: https://vk.com/aLds
//New_url: http://exmpl.lnk/lrulumnnkc   |   In_url: https://vk.com/TauhDgcuxavjzpuo3aVdJo>;qonQ<jy|mQf08lnn|b81QG9sevjeptYWckdxa|
//New_url: http://exmpl.lnk/u6utH2wloY   |   In_url: https://vk.com/2KI<i0{jV3{<TWXleySfs5KPB2{4gIVsQz{am>YdwkBjjsiOhb0im:CFvcIn311JWgsAx>i2c[xt>oc>hdugn3mtwoe|qpx7klyHo6yvLHbNw65cnbcS5J10Tmt[Shhou3QUmoH63tBh5mwbfpvugz2ia3o<|bwil4e|MonyPlD8_1nKhodt;ytlfGPcT8eip1KderAogKTn9dcOtg0zkWmeMas0aRrrs8g{qKdv:I
//New_url: http://exmpl.lnk/1NvfLmMs1h   |   In_url: https://vk.com/xp>iHCzL9aZ{vrzMkXdEv<D8mqmZua
//New_url: http://exmpl.lnk/aVeaykwqfi   |   In_url: https://vk.com/vgjzeWoamdOx{u{yXV2XrTNRk|LXVd6o
