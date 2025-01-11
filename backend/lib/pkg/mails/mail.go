package mails

import (
	"strconv"
	"log"
	"backend/domain/requests"
	"backend/lib/structure"
	"github.com/spf13/viper"
)

func SendMailerOTP(email string, otp int) error {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("MAILTRAP")
	viper.BindEnv("HOST")
	viper.BindEnv("TEMPLATE_ID")
	viper.BindEnv("TOKEN")

	viper.SetEnvPrefix("SMTP")
	viper.BindEnv("FROM_NAME")
	viper.BindEnv("FROM_EMAIL")

	mailtrapURL := viper.GetString("HOST") + "/api/send"

	emailData := requests.MailTrapRequestTemplate{
		From: requests.EmailAddress{
			Email: viper.GetString("FROM_EMAIL"),
			Name: viper.GetString("FROM_NAME"),
		},
		To: []requests.EmailAddress{
			{
				Email: email,
			},
		},
		TemplateUUID: viper.GetString("TEMPLATE_ID"),
		TemplateVariables: requests.OTPCode{
			OTPCode: strconv.Itoa(otp),
		},
	}

	methodData := structure.ConstructorInstance()

	headerValidation := make(map[string]string)
	headerValidation["Authorization"] = "Bearer" + viper.GetString("TOKEN")
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