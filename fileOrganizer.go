package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path := ""
	fmt.Print("Enter a directory path you want to organize : ")
	fmt.Scan(&path)

	files , err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Can't Read Dir : " , err)
	}
	types := ""
	fmt.Print("Enter the file type : ")
	fmt.Scan(&types)

	folderName := ""
	fmt.Print("Enter a folder name to move all the " + types + " to it :" )
	fmt.Scan(&folderName)

	destDir := path + "/" + folderName

	err = os.Mkdir(destDir, 0750)
	if err != nil {
		log.Fatal("Can't create the directory : " , err)
	}

	for _ , file:= range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == "." + types{
			oldPath := filepath.Join(path , file.Name())
			newPath := filepath.Join(destDir , file.Name())
			err := os.Rename(oldPath , newPath)
			if err != nil {
				log.Fatal("failed to move file : " , err)
			}
		}
	}

	fmt.Println("==================100 Done")


}