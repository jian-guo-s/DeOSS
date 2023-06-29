package node

import (
	"fmt"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
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

	return nil
}
