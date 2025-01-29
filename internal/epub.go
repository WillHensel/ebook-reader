package internal

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"log"
)

type Epub struct {
	Title string `xml:"metadata>title"`
}

func ReadEpub(path string) (Epub, error) {

	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	m, err := r.Open("OEBPS/content.opf")
	if err != nil {
		log.Fatal(err)
	}

	info, err := m.Stat()
	if err != nil {
		log.Fatal(err)
	}

	m_bytes := make([]byte, info.Size())
	_, err = m.Read(m_bytes)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}
	m.Close()

	var epub Epub
	err = xml.Unmarshal(m_bytes, &epub)
	if err != nil {
		log.Fatal(err)
	}

	return epub, nil
}
