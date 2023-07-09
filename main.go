package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func organizeFiles(targetFolder string) {
	extensions := map[string]string{
		".jpg":  "Images",
		".png":  "Images",
		".txt":  "TextFiles",
		".pdf":  "PDFs",
		".docx": "WordDocuments",
		".xlsx": "ExcelFiles",
	}

	files, err := ioutil.ReadDir(targetFolder)
	if err != nil {
		fmt.Println("Error reading target folder:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		filePath := filepath.Join(targetFolder, filename)
		extension := strings.ToLower(filepath.Ext(filename))

		folderName := extensions[extension]
		if folderName == "" {
			folderName = "OtherFiles"
		}

		destinationFolder := filepath.Join(targetFolder, folderName)
		err := os.MkdirAll(destinationFolder, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating destination folder:", err)
			continue
		}

		err = os.Rename(filePath, filepath.Join(destinationFolder, filename))
		if err != nil {
			fmt.Println("Error moving file:", err)
			continue
		}
	}

	fmt.Println("Files organized successfully")
}

func main() {
	var targetFolder string
	fmt.Print("Enter the target folder path: ")
	fmt.Scanln(&targetFolder)

	fileInfo, err := os.Stat(targetFolder)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("Invalid folder path.")
		return
	}

	organizeFiles(targetFolder)
}
