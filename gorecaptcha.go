package gorecaptcha

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/api/option"
	recaptcha "google.golang.org/api/recaptchaenterprise/v1"
)

var (
	ErrInvalidMinRisk = errors.New("risk value must be in 0..1")
	ErrInvalidAction  = errors.New("the action attribute in your reCAPTCHA tag does not match the action you are expecting to score")
	ErrLowScore       = errors.New("captcha score is too low")
)

type InvalidTokenError struct {
	Reason string
}

func (err InvalidTokenError) Error() string {
	return fmt.Sprintf("invalid captcha token error, reason: %s", err.Reason)
}

func CreateAssessment(context context.Context, clientOption option.ClientOption, projectID, recaptchaKey, token string) (assessment *recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, err error) {

	ctx := context
	client, err := recaptcha.NewService(ctx, clientOption)

	if err != nil {
		return
	}

	event := &recaptcha.GoogleCloudRecaptchaenterpriseV1Event{
		Token:   token,
		SiteKey: recaptchaKey,
	}

	newAssesment := &recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment{
		Event: event,
	}

	call := client.Projects.Assessments.Create(fmt.Sprintf("projects/%s", projectID), newAssesment)

	return call.Do()
}

func ValidateAssesment(assesment *recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, eventAction string, minRiskScore float64) (err error) {
	if minRiskScore > 1 || minRiskScore < 0 {
		return ErrInvalidMinRisk
	}
	if !assesment.TokenProperties.Valid {
		return InvalidTokenError{Reason: assesment.TokenProperties.InvalidReason}
	}
	if eventAction != assesment.TokenProperties.Action {
		return ErrInvalidAction
	}
	if assesment.RiskAnalysis.Score < minRiskScore {
		return ErrLowScore
	}
	return
}

func CreateAndValidateAssesment(context context.Context, clientOption option.ClientOption, projectID, recaptchaKey, token,
	recaptchaAction string, minRiskScore float64) (assessment *recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, err error) {

	assessment, err = CreateAssessment(context, clientOption, projectID, recaptchaKey, token)
	if err != nil {
		return
	}
	err = ValidateAssesment(assessment, recaptchaAction, minRiskScore)
	return
}
