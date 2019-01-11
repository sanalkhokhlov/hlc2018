package uploader

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/sanalkhokhlov/hlc2018/store"
	"github.com/sanalkhokhlov/hlc2018/store/service"
)

func Upload(dirPath string, ds *service.DataStore) error {
	log.Println("uploading data...")
	zipPath := fmt.Sprintf("%s/data.zip", dirPath)
	optionsPath := fmt.Sprintf("%s/options.txt", dirPath)
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		var result struct {
			Accounts []store.Account `json:"accounts"`
		}

		err = json.NewDecoder(rc).Decode(&result)
		if err != nil {
			return err
		}

		fmt.Printf("File: %s, Account: %v\n", f.Name, len(result.Accounts))
		err = ds.BulkCreate(result.Accounts)
		if err != nil {
			return err
		}

		runtime.GC()
	}

	file, err := os.Open(optionsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	ts := scanner.Text()
	log.Println(ts)

	log.Println("complete")

	return nil
}
