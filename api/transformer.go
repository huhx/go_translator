package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TranslationResponse struct {
	From         string `json:"from"`
	To           string `json:"to"`
	TransResults []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

func transform(jsonString string) string {
	var translation TranslationResponse
	if strings.Contains(jsonString, "error_code") {
		fmt.Printf("error = %s", jsonString)
		return ""
	}

	err := json.Unmarshal([]byte(jsonString), &translation)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return translation.TransResults[0].Dst
}
