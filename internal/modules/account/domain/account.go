package accountdom

import "time"

type Account struct {
	id                *int
	email             string
	password          string
	emailVerification *EmailVerification
	forgotPassword    *ForgotPassword
	timeStampLog      *TimeStampLog
}

type EmailVerification struct {
	emailVerificationCode   string
	emailVerificationSentAt *time.Time
}

type ForgotPassword struct {
	forgotPasswordCode   string
	forgotPasswordSentAt *time.Time
}

type TimeStampLog struct {
	createdAt *time.Time
	updatedAt *time.Time
	deletedAt *time.Time
	createdBy int
	updatedBy int
	deletedBy int
}

func NewAccount(id *int, email string, password string, emailVerification *EmailVerification, forgotPassword *ForgotPassword, timestampLog *TimeStampLog) *Account {
	return &Account{
		id:                id,
		email:             email,
		password:          password,
		emailVerification: emailVerification,
		forgotPassword:    forgotPassword,
		timeStampLog:      timestampLog,
	}
}

func NewEmailVerification(emailVerificationCode string, emailVerificationSentAt *time.Time) *EmailVerification {
	return &EmailVerification{
		emailVerificationCode:   emailVerificationCode,
		emailVerificationSentAt: emailVerificationSentAt,
	}
}

func NewForgotPassword(forgotPasswordCode string, forgotPasswordSentAt *time.Time) *ForgotPassword {
	return &ForgotPassword{
		forgotPasswordCode:   forgotPasswordCode,
		forgotPasswordSentAt: forgotPasswordSentAt,
	}
}

func NewTimeStampLog(createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time, createdBy int, updatedBy int, deletedBy int) *TimeStampLog {
	return &TimeStampLog{
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
		createdBy: createdBy,
		updatedBy: updatedBy,
		deletedBy: deletedBy,
	}
}
