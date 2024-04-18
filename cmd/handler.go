package cmd

import (
	"fmt"
	"trans/api"
	"trans/record"
	"trans/util"
)

func ClearRecords() {
	record.Clear()
	fmt.Println("successfully cleared")
}

func ListRecords() {
	list, err := record.List()
	if err != nil {
		panic(err)
	}
	record.Format(list)
}

func Translate(text string) {
	if util.IsChinese(text) {
		result := api.ToEnglish(text)
		record.Save(text, result)
		println(result)
	} else {
		result := api.ToChinese(text)
		record.Save(text, result)
		println(result)
	}
}
