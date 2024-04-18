package v17localresidentbin

import (
	_ "embed"
	"errors"

	"archive/zip"
	"bytes"
	"io"
	"os"
)

type EmbedZipFile []byte

var ErrZipFileIsEmpty = errors.New("zip file is empty")

func (EmbedZipFile) readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (e EmbedZipFile) ExtractFirst(filename string) (err error) {
	r, err := zip.NewReader(bytes.NewReader(e), int64(len(e)))
	if err != nil {
		return err
	}
	if len(r.File) == 0 {
		return ErrZipFileIsEmpty
	}

	if bb, err := e.readZipFile(r.File[0]); err != nil {
		return err
	} else {
		// filename = r.File[0].Name
		return os.WriteFile(filename, bb, 0644)
	}
}

func (e EmbedZipFile) Extract() error {
	r, err := zip.NewReader(bytes.NewReader(e), int64(len(e)))
	if err != nil {
		return err
	}
	for _, zippedFile := range r.File {
		unzippedBB, err := e.readZipFile(zippedFile)
		if err != nil {
			return err
		}
		if err := os.WriteFile(zippedFile.Name, unzippedBB, 0644); err != nil {
			return err
		}
	}
	return nil
}

//go:embed "stable/lrsrv-debian-x64.zip"
var LRSRV_DEB_X64_STABLE_ZIPPED EmbedZipFile

//go:embed "stable/lrsrv-windows-x64.exe.zip"
var LRSRV_WIN_X64_STABLE_ZIPPED EmbedZipFile
