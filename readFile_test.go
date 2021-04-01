package main

import (
	"fmt"
	"testing"
)

var read Read
var fileByte *[]byte

func TestReadFileJson(t *testing.T) {
	read = &File{
		Name: "id.json",
		Path: "./json/",
	}

	fByte, err := read.ReadFileJson()
	if err != nil {
		t.Error(err)
	}

	fileByte = &fByte
	fmt.Println(string(fByte))
}

func TestReadAndConvert(t *testing.T) {
	fMap, err := read.ReadAndConvert()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fMap)
}

func TestConvertToMap(t *testing.T) {
	var c Convert
	c = DataByte(*fileByte)

	_, err := c.ConvertToMap()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(*fileByte))
}
