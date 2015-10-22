package pdfmagic_test
import (
	pdfmagic "pdfmagic"
	"testing"
	"strconv"
	"os/exec"
)


func TestConvert(t *testing.T) {
	start_page := 1
	end_page := 3
	input_file := "http://www.theukulelereview.com/wp-content/uploads/2014/01/transcribed_corey_pathofwind.pdf"
	output_path := "./tmp"
	imgs, err := pdfmagic.Convert(input_file, output_path, start_page, end_page, "png")
	if err != nil {
		t.Errorf("Error", err)
	}

	expected_imgs_len := end_page - start_page + 1
	if len(imgs) != expected_imgs_len {
		t.Errorf("Wrong size")
		t.Errorf("Expected: ")
		t.Errorf(strconv.Itoa(expected_imgs_len))
		t.Errorf("Got: ")
		t.Errorf(strconv.Itoa(len(imgs)))
	}

	CleanUp(output_path)
}

func CleanUp(dir string) {
	cmd := exec.Command("rm", "-rf", dir)
	err := cmd.Run()
	if err != nil {
	}
}
