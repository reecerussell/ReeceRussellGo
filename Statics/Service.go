package Statics

import (
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Service ... OOP object for static file service
type Service struct {
}

// CreateFile ... creates file
func (service *Service) CreateFile(w http.ResponseWriter, r *http.Request) (fileName string, status int, err error) {
	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, 2560000)
	if err := r.ParseMultipartForm(2560000); err != nil {
		return "", 400, errors.New("FILE_TOO_BIG")
	}

	// parse and validate file and post parameters
	file, _, err := r.FormFile("file")
	if err != nil {
		return "", 400, errors.New("INVALID_FILE")
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", 400, errors.New("INVALID_FILE")
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		break
	default:
		return "", 400, errors.New("INVALID_FILE_TYPE")
	}
	fileName = uuid.New().String()
	fileEndings, err := mime.ExtensionsByType(filetype)
	filePath := fileName + fileEndings[0]
	newPath := filepath.Join("./files", filePath)
	if err != nil {
		return "", 500, errors.New("CANT_READ_FILE_TYPE")
	}

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		return "", 500, errors.New("CANT_WRITE_FILE")
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", 500, errors.New("CANT_WRITE_FILE")
	}

	return filePath, 200, nil
}
