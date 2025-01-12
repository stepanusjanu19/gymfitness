package cache

import (
	"api/lib/utils/formatstring"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type OTP struct {
	OTP       int       `json:"otp"`
	IsUsed    bool      `json:"is_used"`
	ExpiresAt time.Time `json:"expires_at"`
}

var otpCache = make(map[uuid.UUID]*OTP)
var mutex sync.RWMutex

const otpExpirationTime = 5 * time.Minute

func SaveOTP(userId uuid.UUID, otp int) {
	mutex.Lock()
	defer mutex.Unlock()

	otpCache[userId] = &OTP{
		OTP:       otp,
		IsUsed:    false,
		ExpiresAt: time.Now().Add(otpExpirationTime),
	}
}

func GetOTP(userId uuid.UUID) (*OTP, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	otp, exists := otpCache[userId]
	return otp, exists
}

func MarkOTPAsUsed(userId uuid.UUID) {
	mutex.Lock()
	defer mutex.Unlock()

	if otp, exists := otpCache[userId]; exists {
		otp.IsUsed = true
	}
}

func CleanupExpiredOTPs() {
	mutex.Lock()
	defer mutex.Unlock()

	for userId, otp := range otpCache {
		if time.Now().After(otp.ExpiresAt) {
			delete(otpCache, userId)
		}
	}
}

func ValidateOTP(userId uuid.UUID, otp int) (bool, error) {
	cachedOtp, exists := GetOTP(userId)
	if !exists {
		return false, formatstring.FormatStringError("otpnotfound")
	}
	if time.Now().After(cachedOtp.ExpiresAt) {
		CleanupExpiredOTPs()
		return false, formatstring.FormatStringError("otpexpired")
	}
	if cachedOtp.OTP != otp {
		return false, formatstring.FormatStringError("invalidotpinput")
	}
	if cachedOtp.IsUsed {
		return false, formatstring.FormatStringError("otpalreadyused")
	}
	MarkOTPAsUsed(userId)
	return true, nil
}
