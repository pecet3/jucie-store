package storage

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const MAX_SIZE = 1024 * 1024 * 10 // 10mb

type StorageServices interface {
	AddImage(file multipart.File, header *multipart.FileHeader, subpath string) (string, error)
}

type Services struct {
}

func (m Services) AddImage(file multipart.File, header *multipart.FileHeader, subpath string) (string, error) {
	log.Println("Image start")

	defer file.Close()

	// checking a file

	if header.Size > MAX_SIZE {
		return "", errors.New("buff overflow")
	}

	ext := filepath.Ext(header.Filename)

	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", nil
	}

	// temporary file

	tmpFile, err := os.CreateTemp("", "image-*.tmp")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	if _, err = io.Copy(tmpFile, file); err != nil {
		return "", err
	}
	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	// compressing an image

	var img image.Image

	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(tmpFile)
	case ".png":
		img, err = png.Decode(tmpFile)
	}
	if err != nil {
		return "", nil
	}

	fileName := "juice_" + uuid.NewString() + ext

	out, err := os.Create("./static/images" + subpath + fileName)
	if err != nil {
		return "", nil
	}
	defer out.Close()

	switch ext {
	case ".jpg", ".jpeg":
		options := &jpeg.Options{Quality: 100}
		err = jpeg.Encode(out, img, options)
	case ".png":
		err = png.Encode(out, img)
	}
	if err != nil {
		return "", nil
	}
	log.Println("Image upload success: ", fileName)
	return subpath + "/" + fileName, nil
}
