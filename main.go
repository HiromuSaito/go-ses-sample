package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/mail", handleMailRequest)

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
func handleMailRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		sendMail(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
	}
}

const (
    Sender = "sender@example.com"
    Recipient = "recipient@example.com"
    Subject = "件名"
    TextBody = "メール本文"
    CharSet = "UTF-8"
)

func sendMail(w http.ResponseWriter, r *http.Request) {
    sess, err := session.NewSession(&aws.Config{
        Region:aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_SES_HOST")),},
    )
    svc := ses.New(sess)
    
    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{
            },
            ToAddresses: []*string{
                aws.String(Recipient),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Text: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(TextBody),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(CharSet),
                Data:    aws.String(Subject),
            },
        },
        Source: aws.String(Sender),
    }

    result, err := svc.SendEmail(input)
    
    if err != nil {
		log.Printf("sendEmail error:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
	body, _ := json.MarshalIndent(&result, "", "  ")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}