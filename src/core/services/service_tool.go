package services

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../tools/files"
)

func ReadConfigs(path string) (*list.List, error) {
	results := list.New()
	files := files.GetAllFile(path)
	for file := files.Front(); file != nil; file = file.Next() {
		fmt.Println("reading => " + file.Value.(string))
		fileName := file.Value.(string)
		f, ferr := os.Open(fileName)
		if ferr != nil {
			return list.New(), ferr
		}
		content, cerr := ioutil.ReadAll(f)
		if cerr != nil {
			return list.New(), cerr
		}
		log.Println(string(content))
		serviceModel := ServiceModel{}
		serr := json.Unmarshal(content, &serviceModel)
		if serr != nil {
			return list.New(), serr
		}
		results.PushBack(serviceModel)
	}
	return results, nil
}
