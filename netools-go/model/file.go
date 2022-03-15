package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/config"
	"mime/multipart"
	"os"
)

type File struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`
	Filename  string `json:"filename"`
	Size      int64  `json:"size"`
}

func (f *File) Stringify() string {
	b, err := json.Marshal(f)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}
	return string(b)
}

func (f *File) Parse(s string) {
	json.Unmarshal([]byte(s), f)
}

func (f *File) Store(reader multipart.File) error {
	if reader == nil {
		return errors.New("nil reader")
	}
	out, err := os.OpenFile(config.TempFilePath+f.UUID, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, reader)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Remove() error {
	err := os.Remove(config.TempFilePath + f.UUID)
	return err
}
