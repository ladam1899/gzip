package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

var zipFile = "zipfile.gz"

func main() {
	writeZip()
	readZip()
}

func writeZip() {
	handle, err := openFile(zipFile)
	if err != nil {
		fmt.Println("[ERROR] Opening file:", err)
	}

	zipWriter, err := gzip.NewWriterLevel(handle, 9)
	if err != nil {
		fmt.Println("[ERROR] New gzip writer:", err)
	}
	numberOfBytesWritten, err := zipWriter.Write([]byte("Hello, World!\n"))
	if err != nil {
		fmt.Println("[ERROR] Writing:", err)
	}
	err = zipWriter.Close()
	if err != nil {
		fmt.Println("[ERROR] Closing zip writer:", err)
	}
	fmt.Println("[INFO] Number of bytes written:", numberOfBytesWritten)

	closeFile(handle)
}

func readZip() {
	handle, err := openFile(zipFile)
	if err != nil {
		fmt.Println("[ERROR] Opening file:", err)
	}

	zipReader, err := gzip.NewReader(handle)
	if err != nil {
		fmt.Println("[ERROR] New gzip reader:", err)
	}
	defer zipReader.Close()

	fileContents, err := ioutil.ReadAll(zipReader)
	if err != nil {
		fmt.Println("[ERROR] ReadAll:", err)
	}

	fmt.Printf("[INFO] Uncompressed contents: %s\n", fileContents)

	// ** Another way of reading the file **
	//
	// fileInfo, _ := handle.Stat()
	// fileContents := make([]byte, fileInfo.Size())
	// bytesRead, err := zipReader.Read(fileContents)
	// if err != nil {
	//     fmt.Println("[ERROR] Reading gzip file:", err)
	// }
	// fmt.Println("[INFO] Number of bytes read from the file:", bytesRead)

	closeFile(handle)
}

func openFile(fileToOpen string) (*os.File, error) {
	return os.OpenFile(fileToOpen, openFileOptions, openFilePermissions)
}

func closeFile(handle *os.File) {
	if handle == nil {
		return
	}

	err := handle.Close()
	if err != nil {
		fmt.Println("[ERROR] Closing file:", err)
	}
}

const openFileOptions int = os.O_CREATE | os.O_RDWR
const openFilePermissions os.FileMode = 0660
