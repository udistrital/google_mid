package correo

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
)

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("templates/%s", templateFileName))
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
