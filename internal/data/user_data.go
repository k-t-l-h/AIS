package data

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func LoadUserItem() ([]string, []string, [][]float64){
	names := make([]string, 5)
	users := make([]string, 100)
	data := make([][]float64, 100)
	f, err := os.Open("users.csv")
	if err != nil {
		os.Exit(2)
	}

	r := csv.NewReader(f)
	record, err := r.Read()
	for i := 1; i < len(record); i++ {
		users[i-1] = record[i]
	}

	ind := 0
	for {


		record, err = r.Read()
		if err == io.EOF {
			break
		}
		data[ind] = make([]float64, 100)
		values := record
		names = append(names, record[0])
		for i := 1; i < len(values); i++ {
			mark, _ := strconv.Atoi(values[i])
			data[ind][i-1] = float64(mark)
		}
		ind++
	}

	return names,users, data
}
