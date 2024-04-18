package record

import (
	"encoding/csv"
	"os"
	"time"
)

var filepath = "record.csv"

func Save(query string, result string) error {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		header := []string{"查询", "结果", "时间"}
		data := []string{query, result, time.Now().Format("2006-01-02 15:04:05")}

		if err := writer.Write(header); err != nil {
			return err
		}

		if err := writer.Write(data); err != nil {
			return err
		}

	} else {
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		data := []string{query, result, time.Now().Format("2006-01-02 15:04:05")}
		if err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
