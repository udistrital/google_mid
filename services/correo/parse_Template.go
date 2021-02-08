package correo

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"github.com/astaxie/beego"
)

func parseTemplate(templateFileName string, data interface{}) (string, error) {

	path := beego.AppConfig.String("TemplatePath")

	templatePath, err := filepath.Abs(path + "/" + templateFileName)

	if err != nil {
		fmt.Println("err: \n", err)
		return "", errors.New("invalid template name")
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err: \n", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("err: \n", err)
		return "", err
	}

	body := buf.String()
	return body, nil
}
