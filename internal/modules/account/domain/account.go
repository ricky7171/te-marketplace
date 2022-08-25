package accountdom

import "time"

type Account struct {
	Id                *int
	Email             string
	Password          string
	EmailVerification *EmailVerification
	ForgotPassword    *ForgotPassword
	TimeStampLog      *TimeStampLog
}

type EmailVerification struct {
	EmailVerificationCode   string
	EmailVerificationSentAt *time.Time
}

type ForgotPassword struct {
	ForgotPasswordCode   string
	ForgotPasswordSentAt *time.Time
}

type TimeStampLog struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedBy int
	UpdatedBy int
	DeletedBy int
}

func NewAccount(id *int, email string, password string, emailVerification *EmailVerification, forgotPassword *ForgotPassword, timestampLog *TimeStampLog) *Account {
	return &Account{
		Id:                id,
		Email:             email,
		Password:          password,
		EmailVerification: emailVerification,
		ForgotPassword:    forgotPassword,
		TimeStampLog:      timestampLog,
	}
}

func NewEmailVerification(emailVerificationCode string, emailVerificationSentAt *time.Time) *EmailVerification {
	return &EmailVerification{
		EmailVerificationCode:   emailVerificationCode,
		EmailVerificationSentAt: emailVerificationSentAt,
	}
}

func NewForgotPassword(forgotPasswordCode string, forgotPasswordSentAt *time.Time) *ForgotPassword {
	return &ForgotPassword{
		ForgotPasswordCode:   forgotPasswordCode,
		ForgotPasswordSentAt: forgotPasswordSentAt,
	}
}

func NewTimeStampLog(createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time, createdBy int, updatedBy int, deletedBy int) *TimeStampLog {
	return &TimeStampLog{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
		DeletedBy: deletedBy,
	}
}
