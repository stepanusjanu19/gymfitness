package requests

type MailTrapRequestTemplate struct {
	From              EmailAddress   `json:"from"`
	To                []EmailAddress `json:"to"`
	TemplateUUID      string         `json:"template_uuid"`
	TemplateVariables OTPCode        `json:"template_variables"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type OTPCode struct {
	OTPCode string `json:"part_code"`
}

type ValidateOTP struct {
	Email   string `json:"email" binding:"required,email"`
	OTPCode int    `json:"otp_code" binding:"required"`
}
