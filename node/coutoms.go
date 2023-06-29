package node

import (
	"fmt"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(file multipart.File, dir string, fileId string) error {
	dirPath := strings.Replace(dir, "/file", "/catch", -1)
	_, err := os.Stat(dirPath)
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
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	// save form file
	var num int
	var buf = make([]byte, 4*1024*1024)
	for {
		num, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("create occurred:", err)
			return err
		}
		if num == 0 {
			continue
		}
		f.Write(buf[:num])
	}
	f.Sync()
	f.Close()
	f = nil
	return nil
}
