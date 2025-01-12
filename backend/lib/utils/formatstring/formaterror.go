package formatstring

import (
	"errors"
	"fmt"
	"strings"
)

func FormatStringError(err string) error {
	switch {
	case strings.Contains(err, "emailalready"):
		return errors.New("Email already taken")
	case strings.Contains(err, "hashedPassword"):
		return errors.New("Incorrect password")
	case strings.Contains(err, "userfound"):
		return errors.New("User not found")
	case strings.Contains(err, "userinactive"):
		return errors.New("User its inactive")
	case strings.Contains(err, "transactionStartError"):
		return errors.New("Error starting database transaction")
	case strings.Contains(err, "userCreationError"):
		return errors.New("Error creating user")
	case strings.Contains(err, "otpGenerationError"):
		return errors.New("Error generating OTP")
	case strings.Contains(err, "otpEmailError"):
		return errors.New("Error sending OTP email")
	case strings.Contains(err, "otpSaveError"):
		return errors.New("Error saving OTP to verification table")
	case strings.Contains(err, "transactionCommitError"):
		return errors.New("Error committing database transaction")
	case strings.Contains(err, "ErrInvalidToken"):
		return errors.New("invalid token")
	case strings.Contains(err, "ErrInvalidSigningMethod"):
		return errors.New("invalid signing method")
	case strings.Contains(err, "ErrInvalidTokenSign"):
		return errors.New("Invalid token signature")
	case strings.Contains(err, "ErrTokenExpired"):
		return errors.New("Token expired")
	case strings.Contains(err, "ErrJsonEncode"):
		return errors.New("error marshaling body")
	case strings.Contains(err, "ErrRequestCreate"):
		return errors.New("error creating request")
	case strings.Contains(err, "ErrRequestSending"):
		return errors.New("error sending request")
	case strings.Contains(err, "ErrReadResponseBody"):
		return errors.New("error reading response body")
	case strings.Contains(err, "usernotverified"):
		return errors.New("User account is not verified. Please complete OTP verification.")
	case strings.Contains(err, "verificationCheckError"):
		return errors.New("Error checking user verification status.")
	case strings.Contains(err, "otpnotfound"):
		return errors.New("OTP not found in cache")
	case strings.Contains(err, "otpexpired"):
		return errors.New("OTP expired")
	case strings.Contains(err, "invalidotpinput"):
		return errors.New("invalid OTP")
	case strings.Contains(err, "otpalreadyused"):
		return errors.New("OTP already used")
	case strings.Contains(err, "otpusedatabase"):
		return errors.New("OTP not found or already used in the database")
	default:
		return errors.New("An unexpected error occurred")
	}
}

func FormatStringErrorWithDetails(err string, details error) error {
	baseError := FormatStringError(err)
	return fmt.Errorf("%w: %v", baseError, details)
}
