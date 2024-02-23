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

type CaptchaAssessmentService struct {
	ClientOpton             option.ClientOption
	ProjectId, RecaptchaKey string
}

func NewAssessmentService(clientOption option.ClientOption, projectId, recaptchaKey string) *CaptchaAssessmentService {
	return &CaptchaAssessmentService{ClientOpton: clientOption, ProjectId: projectId, RecaptchaKey: recaptchaKey}
}

func (service CaptchaAssessmentService) CreateAssessment(context context.Context, token string) (*recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, error) {
	client, err := recaptcha.NewService(context, service.ClientOpton)
	if err != nil {
		return nil, err
	}
	event := &recaptcha.GoogleCloudRecaptchaenterpriseV1Event{
		Token:   token,
		SiteKey: service.RecaptchaKey,
	}
	assessment := &recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment{
		Event: event,
	}
	call := client.Projects.Assessments.Create(fmt.Sprintf("projects/%s", service.ProjectId), assessment)
	return call.Do()
}

func (service CaptchaAssessmentService) ValidateAssessment(assessment *recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, eventAction string, minRiskScore float64) (err error) {
	if minRiskScore > 1 || minRiskScore < 0 {
		return ErrInvalidMinRisk
	}
	if !assessment.TokenProperties.Valid {
		return InvalidTokenError{Reason: assessment.TokenProperties.InvalidReason}
	}
	if eventAction != assessment.TokenProperties.Action {
		return ErrInvalidAction
	}
	if assessment.RiskAnalysis.Score < minRiskScore {
		return ErrLowScore
	}
	return
}

func (service CaptchaAssessmentService) CreateAndValidateAssessment(context context.Context, token,
	recaptchaAction string, minRiskScore float64) (assessment *recaptcha.GoogleCloudRecaptchaenterpriseV1Assessment, err error) {

	assessment, err = service.CreateAssessment(context, token)
	if err != nil {
		return
	}
	err = service.ValidateAssessment(assessment, recaptchaAction, minRiskScore)
	return
}
