package pdfmagic

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"crypto/rand"
)

func Convert(url string, output_dir string, page int) (string, error) {
	path, err := Download(url, output_dir)
	if err != nil {
		return "", err
	}

	pngs, err := ConvertToPngs(path, page)
	if err != nil {
		return "", err
	}

	return pngs, nil
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

func mkPngsDir(input_path string) (string, error) {
	pngs_dir := input_path + randToken() + "-pngs"

	err := os.MkdirAll(pngs_dir, 0777)
	if err != nil {
		return "", err
	}

	return pngs_dir, nil
}

func ConvertToPngs(input_path string, page int) (string, error) {
	pngs_dir, err := mkPngsDir(input_path)
	if err != nil {
		return "", err
	}
	if page > 0 {
		page = page - 1
	}
	// convert -density 300 img/bitcoin.pdf[0] -quality 100 test.jpg
	input_path = fmt.Sprint(input_path, "[", page, "-", page+2, "]") // s will be "[age:23]"
	fmt.Println(input_path)
	output_path := fmt.Sprint(pngs_dir, "/", page+1, ".png")
	cmd := exec.Command("convert", "-density", "300", "-scene", "1", input_path, "-quality", "100", output_path)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	dir, err := os.Open(pngs_dir)
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
		filenames = append(filenames, pngs_dir+"/"+number+".png")
	}

	filenames_joined := strings.Join(filenames, ",")
	return filenames_joined, nil
}

func randToken() string {
    b := make([]byte, 8)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
