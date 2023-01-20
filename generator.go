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

	var ans [10]byte

	for i := 0; i < 10; i++ {
		ans[i] = tmp[((int(tmp[i]))%len(url))%32]
	}

	return "http://exmpl.lnk/" + string(ans[:])
}

//
// Testing Hash function
//

//hashMap := make(map[string]string, 100000000)
//errCount := 0
//
//var i uint64
//randSource := rand.NewSource(time.Now().UnixNano())
//randGen := rand.New(randSource)
//
//for k := 0; k < 1000; k++ {
//	for i = 0; i < 1000; i++ {
//		curNum := int(uint8(randGen.Int())) + 3
//		//fmt.Println(curNum)
//		newUrl := ""
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
//		//fmt.Println(newUrl)
//
//		curHash := GenerateShortUrl(newUrl)
//		if hashMap[curHash] != "" {
//			fmt.Println("DUBLICATE: ", hashMap[curHash]+"   |   "+newUrl)
//			errCount++
//		} else {
//			hashMap[curHash] = newUrl
//		}
//	}
//
//	fmt.Println("ERROR_COUNT: ", errCount)
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

//New_url: http://exmpl.lnk/mq91urRtqk   |   In_url: xxm2tvkaYbo13MvwyindmVQ2VSY2GnHaxzm5bAX{njq0|U=gFLqboy:3DchDuggIzm
//New_url: http://exmpl.lnk/gckgccmDgn   |   In_url: uirri0Lt{vhkmWDB1wpvup9J{mcoqdfzmhQRdjqd;quhs1En{7ZftuDKAll<xw|05y
//New_url: http://exmpl.lnk/5jirwbKaiP   |   In_url: GcpuNcEiapoeW0tDWPnnvZ[Xy|2kdncmaoYgc6h2BSvlkeMq{vSACwZysMjNhtf0jtWdSqqVY=FlQQm2Z>{4m|vayXUkfpnyxcuflXuxCq=o{nTq0Wsi05oCli{3fdbynbe1kWssBy
//New_url: http://exmpl.lnk/ATY9wY4c5Y   |   In_url: xe9e5zIocPWalfxkYBp4ddhw56CpQ04foqr=es_M>R2jrzlgq00sEjMfog|cnpKktf<uxseLQjkrKD3{NK2qejqxeMr=TG{4oVpv0RLys{T{xa|hIwWaXipfisGhrrlxxMYYmdwlFghsv{{PVe|jhl|ndLrhFyiv|Kwxal|L{cK|qZ0eJbz|xm=pYdlb3zmQ7MPo5qrx<uFvRsEr>j13|im5qfgMqvqj=VeYPhO;bpbzZ
//New_url: http://exmpl.lnk/ltgCCvv8wi   |   In_url: qtKkpMNchYyVx2ye4ob=ptZXcUk4f=QjoHBb5q=|4ltaSbexu0wqQ2LHnc4owhmzalUXexcRhtulcdi>N|gneifanQqTn4nvn
//New_url: http://exmpl.lnk/huqRuRp22f   |   In_url: 9pdtwamgyydHs8YjuWuqY|diukx=Per
//New_url: http://exmpl.lnk/E4x4U3JdUr   |   In_url: XbMgqj{BNlCn{PWU6v3xyboeQlkt0w9luoLrz|orDYqOpVSaOY7jRF0uW=akY1x|vow4oJCvpXjzocsK
//New_url: http://exmpl.lnk/t2OBoEdEoE   |   In_url: ebigjX{hdsj6F
//New_url: http://exmpl.lnk/0rbuJz_uJb   |   In_url: 7odZ_IXo{Guxr{>lHzru8kavtlFt7uUtuMXNlmbb>uyctsui>l|u
//New_url: http://exmpl.lnk/weQ2iimViV   |   In_url: stlUWc0E54a0P3xeJbtS9He2pVd0ed=4>l98KbNKct2>xCroHrtoJ>tnbP1ibW2lD:iobvHtsOya1BW{zku>yjolhOE3iB8g|9FLzbeHiXAB_QZPW5hnoV3X9XyYsy|dA=L:erb2:adyeecSasCosuDLg9ydz80Nzco9zIw0mlithQ6rqHHojisepvjp|4bfqlXYnysz0MnvuOclb<gigtEQr5p{0|
//New_url: http://exmpl.lnk/8wbbK8vvv0   |   In_url: u|0Kdn[mgkld_9yhalQnomk3qnpiyblnokqcGKkDX=EcLJu0c9u7YcQniMMe[XXkvXxnydt1Ou1DFd{b63mZ<pv:<{fsrbZVBraSd3G>k4utYHaVvdtWUbbX6xqjkAXyq43Gmaiq4sts
//New_url: http://exmpl.lnk/HNompnovup   |   In_url: LuvgG|tpOPuXuJWbMvfzbesxpod4rrxEVrRmsh|Hn3D>mR|lucx2YCfmxcvkzwq|01xDeufedwec=g;bauwDtj8wd0LFwlMuviZ{QtEcjmb
//New_url: http://exmpl.lnk/wzbbyzbwQw   |   In_url: fqybds5ns
//New_url: http://exmpl.lnk/6flYicsaca   |   In_url: bx9xz2ir:ccS2nUIbKf1r>d{sNqs|dMfnaksuaQeznj1nOt84esbvrap
//New_url: http://exmpl.lnk/aaa0aaIsDh   |   In_url: 5iVcswYbaqepyVnpw9k{ij;Yznm7zzTVxumhexbYug{mRnf>m<0va4edbhbylppfGgprQ4pmrkqyst0AEPM>nYfh1v0Azn{|dDOoyj6HmtXmVPOXs
//New_url: http://exmpl.lnk/caaccccAAc   |   In_url: |vwR
