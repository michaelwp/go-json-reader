package main

import (
	"encoding/json"
	"io/ioutil"
)

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DataByte []byte

func (f File) ReadAndConvert() (map[string]interface{}, error) {
	// Read file json
	var fJson Read
	
	fJson = File{
		Name: "id.json",
		Path: "./json/",
	}

	fByte, err := fJson.ReadFileJson()
	if err != nil {
		return nil, err
	}

	// convert to map
	var dByte Convert

	dByte = DataByte(fByte)

	dMap, err := dByte.ConvertToMap()
	if err != nil {
		return nil, err
	}

	return dMap, nil
}

func (f File) ReadFileJson() ([]byte, error) {
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
