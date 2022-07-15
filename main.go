package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

type Swagger struct {
	Swagger             string                 `json:"swagger"`
	Info                interface{}            `json:"info"`
	Host                string                 `json:"host"`
	Schemes             interface{}            `json:"schemes"`
	Consumes            interface{}            `json:"consumes"`
	Produces            interface{}            `json:"produces"`
	Paths               map[string]interface{} `json:"paths"`
	Definitions         map[string]interface{} `json:"definitions"`
	SecurityDefinitions interface{}            `json:"securityDefinitions"`
	Security            interface{}            `json:"security"`
}

func main() {
	baseJson := "base.json"
	docJson, err := ReadJson(baseJson)
	if err != nil {
		log.Printf("read file %s err,%v", baseJson, err)
		return
	}
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		if path.Ext(file.Name()) != ".json" || file.Name() == baseJson || file.Name() == "doc.json" {
			continue
		}

		fileJson, err := ReadJson(file.Name())
		if err != nil {
			log.Printf("read file %s err,%v", file.Name(), err)
			continue
		}
		docJson.Paths = MapMerge(docJson.Paths, fileJson.Paths)
		docJson.Definitions = MapMerge(docJson.Definitions, fileJson.Definitions)
	}
	marshal, err := json.Marshal(docJson)
	if err != nil {
		log.Println("marshal json err,", err)
		return
	}

	if err := ioutil.WriteFile("doc.json", marshal, 0644); err != nil {
		log.Println("generate json file err,", err)
		return
	}
	log.Println("Done")
}

func ReadJson(filename string) (*Swagger, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var swag Swagger
	if err := json.Unmarshal(fileContent, &swag); err != nil {
		return nil, err
	}
	return &swag, nil
}

func MapMerge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
