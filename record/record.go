package record

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"text/tabwriter"
	"time"
)

func GetFilePath() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/" + viper.GetString("fileName")
}

func Save(query string, result string) error {
	filepath := GetFilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		header := []string{"query", "result", "createTime"}
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

func List() ([]map[string]string, error) {
	filepath := GetFilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return []map[string]string{}, nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dataList []map[string]string
	header := records[0]
	for _, record := range records[1:] {
		data := make(map[string]string)
		for i, value := range record {
			data[header[i]] = value
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func Format(data []map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintf(w, "%s\t%s\t%s\n", "Query", "Result", "CreateTime")
	fmt.Fprintf(w, "%s\t%s\t%s\n", "----", "---", "-------")

	// 打印数据
	for _, row := range data {
		fmt.Fprintf(w, "%s\t%s\t%s\n", row["query"], row["result"], row["createTime"])
	}

	// 刷新并输出
	w.Flush()
}
