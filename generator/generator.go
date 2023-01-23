package generator

import (
	"crypto/sha256"
	"my_project/urlgen/config"
)

func GenerateShortUrl(url string) string {

	tmp := sha256.Sum256([]byte(url))

	for i := 0; i < config.ShortUrlLen; i++ {
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

	return config.GenUrl + string(tmp[:config.ShortUrlLen])
}

//
// Examples:
//
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
