package mails

import (
	"api/domain/requests"
	"api/lib/structure"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

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
