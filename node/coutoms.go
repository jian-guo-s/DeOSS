package node

import (
	"fmt"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func UploadFile(gin *gin.Context, dir string, fileId string) error {
	log.Println("start upload file----------------")
	file, err := gin.FormFile("file")
	if err != nil {
		fmt.Println("Get file failed:", err)
		return err
	}
	dirPath := strings.Replace(dir, "/file", "/catch", -1)
	_, err = os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, pattern.DirMode)
		if err != nil {
			fmt.Println("Error occurred:", err)
			return err
		}
	}

	cacheFilePath := filepath.Join(dirPath, fileId)
	f, err := os.Create(cacheFilePath)
	if err != nil {
		fmt.Println("create occurred:", err)
		return err
	}
	defer f.Close()
	src, err := file.Open()
	if err != nil {
		fmt.Println("Internal Server Error:", err)
		return err
	}
	defer src.Close()

	_, err = io.Copy(f, src)
	if err != nil {
		fmt.Println("Copy Internal Server Error:", err)
		return err
	}
	imageRegex := regexp.MustCompile(`(?i)\.(jpg|jpeg|png|gif|svg)$`)
	if imageRegex.MatchString(file.Filename) {
		showBaseDirPath := strings.Replace(dir, "/file", "/show", -1)
		showDirPath := filepath.Join(showBaseDirPath, fileId)
		_, err = os.Stat(showDirPath)
		if err != nil {
			err = os.MkdirAll(showDirPath, pattern.DirMode)
			if err != nil {
				fmt.Println("Error occurred:", err)
				return err
			}
		}
		showFilePath := filepath.Join(showDirPath, file.Filename)
		showOut, err := os.Create(showFilePath)
		if err != nil {
			fmt.Println("show file create occurred:", err)
			return err
		}
		defer showOut.Close()
		fileOut, err := file.Open()
		if err != nil {
			fmt.Println("Internal Server Error:", err)
			return err
		}
		defer fileOut.Close()
		log.Println("start upload show file")
		_, err = io.Copy(showOut, fileOut)
		if err != nil {
			fmt.Println("Copy show Internal Server Error:", err)
			return err
		}
	}
	return nil
}
