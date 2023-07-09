package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var targetFolder string

func organizeFiles() error {
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
		return fmt.Errorf("error reading target folder: %v", err)
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
			log.Printf("error creating destination folder: %v", err)
			continue
		}

		destinationPath := filepath.Join(destinationFolder, filename)
		if _, err := os.Stat(destinationPath); err == nil {
			// Handle file name conflicts
			newFilename := fmt.Sprintf(
				"%s_%d%s",
				strings.TrimSuffix(filename, extension),
				file.ModTime().Unix(),
				extension,
			)
			destinationPath = filepath.Join(destinationFolder, newFilename)
		}

		err = os.Rename(filePath, destinationPath)
		if err != nil {
			log.Printf("error moving file: %v", err)
			continue
		}
	}

	return nil
}

func init() {
	flag.StringVar(&targetFolder, "target", "", "Target folder path")
	flag.Parse()
}

func main() {
	if targetFolder == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("error getting the current working directory: %v", err)
		}
		targetFolder = wd
	}

	err := organizeFiles()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files organized successfully")
}
