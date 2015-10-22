package pdfmagic

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Convert(url string, output_dir string, page int, end_page int, format string) (string, error) {
	path, err := Download(url, output_dir)
	if err != nil {
		return "", err
	}

	imgs, err := ConvertToImgs(path, page, end_page, format)
	if err != nil {
		return "", err
	}

	return imgs, nil
}

func GetMimeTypeByFilename(base string) string {
	return mime.TypeByExtension(filepath.Ext(base))
}

func Download(url string, output_dir string) (string, error) {
	err := os.MkdirAll(output_dir, 0777)
	if err != nil {
		return "", err
	}

	base := filepath.Base(url)
	path := output_dir + "/" + base

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return path, nil
}

func mkImgsDir(input_path string) (string, string, error) {
	input_path_slice := strings.Split(input_path, "/")

	imgs_dir := input_path + randToken() + "-imgs"
	err := os.MkdirAll(imgs_dir, 0777)
	if err != nil {
		return "", "", err
	}


	filename := strings.Replace(input_path_slice[len(input_path_slice)-1], ".pdf", "", -1)
	return imgs_dir, filename, nil
}

func ConvertToImgs(input_path string, page int, end_page int, format string) (string, error) {
	imgs_dir, filename, err := mkImgsDir(input_path)
	if err != nil {
		return "", err
	}
	if page > 0 {
		page = page - 1
	}
	if end_page > 0 {

		end_page = end_page - 1
	}

	
	input_path = fmt.Sprint(input_path, "[", page, "-", end_page, "]")
	// Construct output path
	output_path := fmt.Sprint(imgs_dir, "/", filename, "-%d.", format)
	// set command
	// i.e. convert -density 300 -scene 1 img/bitcoin.pdf[0-1] -quality 100 test-%d.jpg
	cmd := exec.Command("convert", "-density", "300", "-scene", "1", input_path, "-quality", "100", output_path)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	// Run command
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	dir, err := os.Open(imgs_dir)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	filenames := make([]string, 0)

	fis, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	for i := 1; i <= len(fis); i++ {
		number := strconv.Itoa(i)
		filenames = append(filenames, imgs_dir+"/"+filename+"-" + number +"." + format)
	}

	filenames_joined := strings.Join(filenames, ",")
	return filenames_joined, nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
