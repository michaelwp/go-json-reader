package main

type Read interface {
	ReadFileJson() ([]byte, error)
	ReadAndConvert() (map[string]interface{}, error)
}

type Convert interface {
	ConvertToMap() (map[string]interface{}, error)
}
