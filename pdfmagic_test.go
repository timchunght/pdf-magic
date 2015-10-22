// package pdfmagic_test

// import (
// 	pdfmagic "github.com/timchunght/pdfmagic"
// 	"testing"
// )

// const OUTPUT = "./tmp/01guest.pdf-pngs/1.png,./tmp/01guest.pdf-pngs/2.png,./tmp/01guest.pdf-pngs/3.png,./tmp/01guest.pdf-pngs/4.png,./tmp/01guest.pdf-pngs/5.png,./tmp/01guest.pdf-pngs/6.png,./tmp/01guest.pdf-pngs/7.png,./tmp/01guest.pdf-pngs/8.png,./tmp/01guest.pdf-pngs/9.png,./tmp/01guest.pdf-pngs/10.png,./tmp/01guest.pdf-pngs/11.png,./tmp/01guest.pdf-pngs/12.png,./tmp/01guest.pdf-pngs/13.png"

// func TestConvert(t *testing.T) {
// 	pngs, err := pdfmagic.Convert("http://www.bramstoker.org/pdf/stories/03guest/01guest.pdf", "./tmp")
// 	if err != nil {
// 		t.Errorf("Error", err)
// 	}

// 	if pngs != OUTPUT {
// 		t.Errorf("Pngs output was wrong")
// 		t.Errorf("Expected: ")
// 		t.Errorf(OUTPUT)
// 		t.Errorf("Got: ")
// 		t.Errorf(pngs)
// 	}
// }

package pdfmagic_test
import (
	"fmt"
	pdfmagic "pdfmagic"
	// "os/exec"
)

func main() {
	pngs, err := pdfmagic.Convert("http://www.theukulelereview.com/wp-content/uploads/2014/01/transcribed_corey_pathofwind.pdf", "./img", 1, 3, "png")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pngs)
}


