package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func currentWorkingPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func doesFileExists(pathToExam string) bool {
	info, err := os.Stat(pathToExam)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createFile(dirpath string, filename string, filesize int, overwrite bool) error {
	path := filepath.Join(dirpath, filename)

	if doesFileExists(path) && !overwrite {
		return fmt.Errorf("File to be created already exists.\n" +
			"To overwirte use the flag -overwrite\n" +
			"The -overwrite flag must always be the last parameter\n" +
			"Example:\n\t" +
			"$ touch-fake some_name.txt 1000 -overwrite")
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Can't create file on path: " + path)
	}
	defer file.Close()

	buffer := make([]byte, filesize)
	file.Write(buffer)
	file.Sync()

	return nil
}

func canOverwrite() bool {
	if len(os.Args) != 4 || os.Args[3] != "-overwrite" {
		return false
	}
	return true
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Pass the file size and file name.\n" +
			"Example: To create a file called file_name.txt with 1000 bytes size that don't overwrite any file\n\t" +
			"$ touch-fake file_name.txt 1000\n" +
			"To overwrite a file use the -overwrite flag, it must always be the last parameter\n\t" +
			"$ touch-fake file_name.txt 1000 -overwrite")
		os.Exit(1)
	}

	filesize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("The file bytes size isn't valid")
		os.Exit(1)
	}

	filepath, err := currentWorkingPath()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	filename := os.Args[1]

	err = createFile(filepath, filename, filesize, canOverwrite())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
