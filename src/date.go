package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func Init() {
	err := os.MkdirAll("~/.config/date", os.ModePerm)
	if err != nil {
		fmt.Println("Could create config folder")
		return
	}
	_, err = os.Create("~/.config/date/dates.json")
	if err != nil {
		fmt.Println("Couldnt create file")
		return
	}
}

func Add(date string, name string) {
	dates := open()
	date_formatted, err := format(date)
	if err != nil {
		return
	}
	dates[date_formatted] = append(dates[date_formatted], name)
	save(dates)
}

func Remove(date string, name string) {
	dates := open()

}

func Remind() {
	dates := open()
	current_date := time.Now().Format("DateOnly")
	formatted_date_arr := strings.Split(current_date, "-")
	output_date := ""
	for _, s := range formatted_date_arr {
		output_date += s + " "
	}
	if len(dates[current_date]) == 0 {
		return
	}
	fmt.Println(output_date)
	for _, date := range dates[current_date] {
		fmt.Println(date)
	}
}

func ListAll() {

}

func List(date string) {
	
}

func open() map[string][]string {
	data, err := os.ReadFile("~/.config/date/dates.json")
	if err != nil {
		fmt.Println("Couldnt open file")
		return nil
	}
	dates := map[string][]string{}
	json.Unmarshal(data, dates)
	return dates
}

func save(data map[string][]string) {
	file, err := os.Open("~/.config/date/dates.json")
	if err != nil {	
		fmt.Println("Couldnt open file")
		return
	}
	defer file.Close()
	json_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data")
	}
	_, err = file.WriteAt(json_data, 0)
	if err != nil {
		fmt.Println("Could not write data into json file")
	}
}

func format(date string) (string, error) { 
	split_date := strings.Split(date, ".")
	if len(split_date) != 3 {
		fmt.Println("Incorrect date format given. Seperate with: '.'")
		return "", errors.New("Incorrect format")
	}
	return split_date[0] + "-" + split_date[1] + "-" + split_date[2], nil
}