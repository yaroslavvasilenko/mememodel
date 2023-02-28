package meme_store_models

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(linkDownload string, idFile string) error {

	resp, err := http.Get(idFile)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(FilePath + idFile)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
