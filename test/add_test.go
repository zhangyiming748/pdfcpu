package test

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"testing"
)

func TestAdd(t *testing.T) {
	fp := "C:\\Users\\zen\\Github\\pdfcpu\\test\\党员干部网络言行“十不准”.pdf"
	out := "C:\\Users\\zen\\Github\\pdfcpu\\test\\党员干部网络言行十不准.pdf"
	m := model.NewAESConfiguration("userPW", "ownerPW", 256)
	err := api.EncryptFile(fp, out, m)
	if err != nil {
		return
	}
}
