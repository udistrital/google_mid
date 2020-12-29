package correo

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
)

func parseTemplate(templateFileName string, data interface{}) (string, error) {

	fmt.Println("templatePath: \n", fmt.Sprintf("templates/%s", templateFileName)) //Sobra
	var textoValidacion = "templatePath: " + fmt.Sprintf("templates/%s", templateFileName) //Sobra

	templatePath, err := filepath.Abs(fmt.Sprintf("templates/%s", templateFileName))

	textoValidacion = textoValidacion + "\ntemplatePath: " +  templatePath //Sobra

	fmt.Println("templatePath: \n", templatePath) //Sobra

	if err != nil {
		fmt.Println("err: \n", err)
		// return "", errors.New("invalid template name")
		textoValidacion = textoValidacion + "\nerr: " + err.Error() //Sobra
		fmt.Println(textoValidacion) //Sobra
		return textoValidacion, errors.New("invalid template name") //Sobra
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err: \n", err)
		// return "", err
		textoValidacion = textoValidacion + "\nerr: " + err.Error() //Sobra
		fmt.Println(textoValidacion) //Sobra
		return textoValidacion, err //Sobra
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("err: \n", err)
		// return "", err
		textoValidacion = textoValidacion + "\nerr: " + err.Error() //Sobra
		fmt.Println(textoValidacion) //Sobra
		return textoValidacion, err //Sobra
	}

	body := buf.String()
	return body, nil
}
