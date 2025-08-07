package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"strings"
)

func SetPasswd(fp, passwd string) error {
	passwpeded := strings.Replace(fp, ".pdf", "withPasswd.pdf", 1)
	passwpeded = strings.Replace(fp, ".PDF", "withPasswd.pdf", 1)
	m := model.NewAESConfiguration(passwd, passwd, 256)
	err := api.EncryptFile(fp, passwpeded, m)
	if err != nil {
		return err
	}
	return nil
}
