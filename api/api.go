package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"trans/util"
)

var baseURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"

func ToChinese(text string) string {
	var appId = viper.GetString("appId")
	var appSecret = viper.GetString("appSecret")
	salt := util.RandomSalt(12)
	signString := fmt.Sprintf("%s%s%s%s", appId, text, salt, appSecret)
	hash := md5.Sum([]byte(signString))
	signature := hex.EncodeToString(hash[:])

	params := url.Values{
		"q":     {text},
		"from":  {"en"},
		"to":    {"zh"},
		"appid": {appId},
		"salt":  {salt},
		"sign":  {signature},
	}

	resp, _ := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return ""
	}

	return transform(string(body))
}

func ToEnglish(text string) string {
	var appId = viper.GetString("appId")
	var appSecret = viper.GetString("appSecret")
	salt := util.RandomSalt(12)
	signString := fmt.Sprintf("%s%s%s%s", appId, text, salt, appSecret)
	hash := md5.Sum([]byte(signString))
	signature := hex.EncodeToString(hash[:])

	params := url.Values{
		"q":     {text},
		"from":  {"zh"},
		"to":    {"en"},
		"appid": {appId},
		"salt":  {salt},
		"sign":  {signature},
	}

	resp, _ := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return ""
	}

	return transform(string(body))
}
