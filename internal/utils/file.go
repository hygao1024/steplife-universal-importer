package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateCSVFile
//
//	@Description: 		CSV 文件创建
//	@param csvFilePath
//	@return bool		文件曾是否存在
//	@return error
func CreateCSVFile(csvFilePath string) (bool, error) {
	existed := true
	if _, err := os.Stat(csvFilePath); os.IsNotExist(err) {
		file, err := os.Create(csvFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return false, err
		}
		defer file.Close()

		fmt.Printf("Created new CSV file: %s\n", csvFilePath)
		existed = false
	}
	return existed, nil
}

// GetAllFilePath
//
//	@Description:
//	@param path					需要获取的路径
//	@return map[string][]string	文件路径。key 为文件夹名，value 为文件路径
//	@return error
func GetAllFilePath(path string) (map[string][]string, error) {

	sourceFiles := make(map[string][]string)

	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			// 构建完整文件路径
			files, err := os.ReadDir(filepath.Join(path, dir.Name()))
			if err != nil {
				return nil, err
			}
			for _, file := range files {
				if !file.IsDir() && file.Name() != ".DS_Store" {
					sourceFiles[dir.Name()] = append(
						sourceFiles[dir.Name()],
						filepath.Join(path, dir.Name(), file.Name()),
					)
				}
			}
		}
	}

	return sourceFiles, nil
}

// ReadFile
//
//	@Description: 	读文件
//	@param filePath
//	@return []byte
//	@return error
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

// WriteCSV
//
//	@Description: 	写入 CSV 文件
//	@param filePath
//	@param rows
//	@return error
func WriteCSV(filePath string, rows [][]string) error {
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.WriteAll(rows)
	if err != nil {
		return err
	}
	csvWriter.Flush()
	return nil
}
