package correo

import (
	"context"
	"encoding/base64"
	// "errors"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"github.com/udistrital/google_mid/models"
)

var GmailService *gmail.Service

func OAuthGmailService() {

	config := oauth2.Config{
		ClientID:     beego.AppConfig.String("GmailClientId"),
		ClientSecret: beego.AppConfig.String("GmailClientSecret"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  beego.AppConfig.String("GmailTokenAccess"),
		RefreshToken: beego.AppConfig.String("GmailTokenRefresh"),
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is initialized")
	}
}

func SendEmailOAUTH2_V2(datosNotificacion models.Notificacion) (bool, error) {
	fmt.Println("Email sending...")

	// Se arma la lista de destinatarios separados por coma
	to := ""
	separador := ""
	for _, v := range datosNotificacion.To {
		to += separador + v ;
		separador = ","
	}
	emailTo := "To: " + to + "\r\n"

	// Se arma la lista de correos que recibirán copia
	hasCC := false
	cc := ""
	separador = ""
	for _, v := range datosNotificacion.CC {
		cc += separador + v ;
		separador = ","
		hasCC = true
	}
	emailCc := ""
	if (hasCC == true){
		emailCc = "Cc: " + cc + "\r\n"
	}

	// Se arma la lista de correos que recibirán copia oculta
	hasBCC := false
	bcc := ""
	separador = ""
	for _, v := range datosNotificacion.BCC {
		bcc += separador + v ;
		separador = ","
		hasBCC = true
	}
	emailBcc := ""
	if (hasBCC == true){
		emailBcc = "bcc: " + bcc + "\r\n"
	}

	// Se ensambla el contenido
	emailBody, err := parseTemplate(datosNotificacion.TemplateName, datosNotificacion.TemplateData)
	// fmt.Println("Email body: ", emailBody)
	// fmt.Println("err: ", err.Error())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	subject := "Subject: " + datosNotificacion.Subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + emailCc + emailBcc + subject + mime + "\n" + emailBody)

	var message gmail.Message
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println("Email sent")

	return true, nil
}