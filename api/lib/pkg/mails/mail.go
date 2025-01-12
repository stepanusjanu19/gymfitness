package mails

import (
	"strconv"
	"log"
	"api/domain/requests"
	"api/lib/structure"
	"github.com/spf13/viper"
)

func SendMailerOTP(email string, otp int) error {

	viper.SetConfigFile("../.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	mailtrapURL := viper.GetString("MAIL_HOST") + "/api/send"

	emailData := requests.MailTrapRequestTemplate{
		From: requests.EmailAddress{
			Email: viper.GetString("MAIL_FROM_EMAIL2"),
			Name: viper.GetString("MAIL_FROM_NAME"),
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