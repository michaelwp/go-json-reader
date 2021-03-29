package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DataByte []byte

func main() {
	f := File{
		Name: "en.json",
		Path: "./json/",
	}

	file, err := f.readFileJson()
	if err != nil {
		fmt.Print(fmt.Errorf("error: %v", err))
	}

	dataByte := DataByte(file)

	resultMap, err := dataByte.ConvertToMap()

	fmt.Println(resultMap["hello_world"])
}

func (f File) readFileJson() ([]byte, error) {
	fByte, err := ioutil.ReadFile(f.Path + f.Name)
	if err != nil {
		return nil, err
	}

	return fByte, nil
}

func (d DataByte) ConvertToMap() (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal(d, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
