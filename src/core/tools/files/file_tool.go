package files

import (
	"container/list"
	"fmt"
	"io/ioutil"
)

func GetAllFile(pathname string) *list.List {
	results := list.New()
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return results
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			filePathName := pathname + "/" + fi.Name()
			results.PushBack(filePathName)
		}
	}
	return results
}
