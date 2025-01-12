package mails

import (
	"api/domain/requests"
	"api/lib/structure"
	"context"
	"log"
	"strconv"
	"time"

	"github.com/mailersend/mailersend-go"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

//# ===================== ===================== ===================== #

func SendMailerOTP(email string, otp int) error {

	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = "/go_project/src/gymfitness/api/.env"
	}
	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	mailtrapURL := viper.GetString("MAIL_HOST") + "/api/send"

	emailData := requests.MailTrapRequestTemplate{
		From: requests.EmailAddress{
			Email: viper.GetString("MAIL_FROM_EMAIL2"),
			Name:  viper.GetString("MAIL_FROM_NAME"),
		},
		To: []requests.EmailAddress{
			{
				Email: email,
			},
		},
		TemplateUUID: viper.GetString("MAIL_TEMPLATE_ID"),
		TemplateVariables: requests.OTPCode{
			OTPCode: strconv.Itoa(otp),
		},
	}

	methodData := structure.ConstructorInstance()

	headerValidation := make(map[string]string)
	headerValidation["Authorization"] = "Bearer " + viper.GetString("MAIL_TOKEN")
	headerValidation["Content-Type"] = "application/json"

	for key, value := range headerValidation {
		methodData.Headers[key] = value
	}

	response, err := structure.RESTApi(methodData.POST, mailtrapURL, methodData, emailData)
	if err != nil {
		return err
	}

	log.Println("POST Response:", response)
	return nil
}

//# ===================== ===================== ===================== #
//# ===================== ===================== ===================== #

func SendMailerOTPbyAPI(email, fullname string, otp int) error {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	tokenMailer := viper.GetString("MAIL_SEND_TOKEN")

	msObject := mailersend.NewMailersend(tokenMailer)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subjectText := "OTP Verification"

	from := mailersend.From{
		Name:  viper.GetString("MAIL_SEND_FROM_NAME"),
		Email: viper.GetString("MAIL_SEND_FROM_EMAIL"),
	}

	recipients := []mailersend.Recipient{
		{
			Name:  fullname,
			Email: email,
		},
	}

	personalization := []mailersend.Personalization{
		{
			Email: email,
			Data: map[string]interface{}{
				"part_code":     strconv.Itoa(otp),
				"support_email": viper.GetString("MAIL_SEND_SUPPORT_EMAIL"),
			},
		},
	}

	mailerTemplateId := viper.GetString("MAIL_SEND_TEMPLATE_ID")

	tag := []string{}

	messageSender := msObject.Email.NewMessage()

	messageSender.SetFrom(from)
	messageSender.SetRecipients(recipients)
	messageSender.SetSubject(subjectText)
	messageSender.SetTemplateID(mailerTemplateId)
	messageSender.SetPersonalization(personalization)

	messageSender.SetTags(tag)

	res, _ := msObject.Email.Send(ctx, messageSender)

	log.Println(res.Header.Get("X-Message-Id"))
	return nil
}

//# ===================== ===================== ===================== #
