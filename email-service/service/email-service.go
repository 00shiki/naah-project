package service

import (
	"email-service/model"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"bytes"

	"gopkg.in/gomail.v2"
    "github.com/joho/godotenv"
)

type EmailService struct {
    SMTPHost     string
    SMTPPort     int
    SMTPUsername string
    SMTPPassword string
    FromEmail    string
}

func NewEmailService() (*EmailService, error) {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	port, _ := strconv.Atoi(os.Getenv("SMTPPort"))
    return &EmailService{
        SMTPHost:     os.Getenv("SMTPHost"),
        SMTPPort:     port,
        SMTPUsername: os.Getenv("SMTPUsername"),
        SMTPPassword: os.Getenv("SMTPPassword"),
        FromEmail:    os.Getenv("FromEmail"),
    }, nil
}

func (es *EmailService) SendEmail(payload model.EmailPayload) error {
    m := gomail.NewMessage()
    m.SetHeader("From", es.FromEmail)
    m.SetHeader("To", payload.To)
    m.SetHeader("Subject", payload.Subject)

    var body string

    // Choose template based on email type
    if payload.Type == "verification" {
        data := map[string]interface{}{
            "VerificationURL": payload.VerificationURL,
        }
        body = parseTemplate("templates/verification.html", data)
    } else if payload.Type == "receipt" {
        data := map[string]interface{}{
            "OrderID": payload.OrderID,
        }
        body = parseTemplate("templates/receipt.html", data)
    } else {
        return fmt.Errorf("unknown email type: %s", payload.Type)
    }

    m.SetBody("text/html", body)

    d := gomail.NewDialer(es.SMTPHost, es.SMTPPort, es.SMTPUsername, es.SMTPPassword)

    if err := d.DialAndSend(m); err != nil {
        return fmt.Errorf("failed to send email: %v", err)
    }
    return nil
}

func parseTemplate(templatePath string, data map[string]interface{}) string {
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Fatalf("Failed to parse template: %v", err)
    }

    var tplOutput bytes.Buffer
    err = tmpl.Execute(&tplOutput, data)
    if err != nil {
        log.Fatalf("Failed to execute template: %v", err)
    }

    return tplOutput.String()
}


