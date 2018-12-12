package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../tools/files"
)

func ReadConfigs(path string) (map[string]ServiceModel, error) {
	var results map[string]ServiceModel = map[string]ServiceModel{}
	files := files.GetAllFile(path)
	for file := files.Front(); file != nil; file = file.Next() {
		fmt.Println("reading => " + file.Value.(string))
		fileName := file.Value.(string)
		f, ferr := os.Open(fileName)
		if ferr != nil {
			return map[string]ServiceModel{}, ferr
		}
		content, cerr := ioutil.ReadAll(f)
		if cerr != nil {
			return map[string]ServiceModel{}, cerr
		}
		log.Println(string(content))
		serviceModel := ServiceModel{}
		serr := json.Unmarshal(content, &serviceModel)
		if serr != nil {
			return map[string]ServiceModel{}, serr
		}
		results[serviceModel.ServicePath] = serviceModel
	}
	return results, nil
}
