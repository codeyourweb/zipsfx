package zipsfx

import (
	"archive/zip"
	"bytes"
	b64 "encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func BuildSFX(inputDirectory string, executeCommand string, outputSfxExe string) error {
	// compress inputDirectory into archive
	archive, err := recursiveCompressFolder(inputDirectory, executeCommand)
	if err != nil {
		return err
	}

	// embed winrar-zipsfx binary
	sDec, err := b64.StdEncoding.DecodeString(zipSfxBinary)
	if err != nil {
		return err
	}

	file, err := os.Create(outputSfxExe)
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write([]byte(sDec))
	file.Write(archive.Bytes())

	return nil
}

func recursiveCompressFolder(basePath string, executeCommand string) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	archive := zip.NewWriter(&buffer)
	err := filepath.Walk(basePath, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(basePath))
		sevenzipFile, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(sevenzipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})

	var config = "the comment below contains sfx script commands\r\n\r\n" +
		"Path=%temp%" + "\r\n" +
		"Setup=" + executeCommand + "\r\n" +
		"Silent=1" + "\r\n" +
		"Overwrite=1"

	archive.SetComment(config)

	if err != nil {
		return buffer, err
	}
	err = archive.Close()
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
