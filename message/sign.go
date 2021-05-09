/*
@Time : 2021/5/7 上午1:39
@Author : RenJun
@File : sign
@Description :
@CopyRight:
*/
package message

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

type ParamSign map[string]string

func (p ParamSign) sign(pushSecret string) (string, error) {
	var signStr []string
	for k, v := range p {
		if len(v) == 0 {
			continue
		}
		signStr = append(signStr, k)
	}

	sort.Strings(signStr)

	var signStrTmp string
	for _, v := range signStr {
		if p[v] != "" {
			signStrTmp += v + "=" + p[v] + "&"
		}
	}

	signStrTmp += "secret=" + pushSecret

	tmpSha := sha256.New()
	tmpSha.Write([]byte(signStrTmp))
	return hex.EncodeToString(tmpSha.Sum(nil)), nil
}
