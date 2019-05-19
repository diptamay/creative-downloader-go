package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const defaultFPath = "links.csv"
const defaultDirectory = "download"

type Downloader interface {
	// Download returns the content from URL and saves it to specified file system
	Download(url string) (err error)
}

// ParseAndDownload uses downloader to download the content from various urls and save to specified filesystem
func ParseAndDownload(fpath string, downloader Downloader) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println(url)
		err := downloader.Download(url)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fptr := flag.String("fpath", defaultFPath, "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	cd := CreativeDownloader{defaultDirectory}
	ParseAndDownload(*fptr, cd)
}

type CreativeDownloader struct {
	dpath string
}

func (cd CreativeDownloader) Download(url string) error {
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		io.Copy(writer, body)
	}()

	file, err := os.Open(cd.dpath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	io.Copy(file, reader)
	reader.Close()

	return nil
}
