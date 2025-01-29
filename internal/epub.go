package internal

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"regexp"
)

type Epub struct {
	Title string `xml:"metadata>title"`
}

func ReadEpub(path string) (Epub, error) {

	r, err := zip.OpenReader(path)
	if err != nil {
		return Epub{}, err
	}
	defer r.Close()

	pat := regexp.MustCompile(`.*\.opf`)
	for _, f := range r.File {
		if pat.MatchString(f.Name) {
			epub, err := readFromManifest(f)
			if err != nil {
				return Epub{}, err
			}

			return epub, nil
		}
	}

	return Epub{}, errors.New("epub manifest file not found")
}

func readFromManifest(f *zip.File) (Epub, error) {
	m, err := f.Open()
	if err != nil {
		return Epub{}, err
	}

	m_bytes, err := io.ReadAll(m)
	if err != nil && !errors.Is(err, io.EOF) {
		return Epub{}, err
	}
	m.Close()

	var epub Epub
	err = xml.Unmarshal(m_bytes, &epub)
	if err != nil {
		return Epub{}, err
	}

	return epub, nil
}
