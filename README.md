# About package
This package implements google recaptcha code verification.

### Installation 
```bash
go get github.com/vloldik/gorecaptcha
```

### Usage
```go
// Create assessment service
assessmentService := gorecaptcha.NewAssessmentService(option, projectId, captchaKey)
// Create and validate assessment of given token, check if action is LOGIN and score is more than 0.4
assessment, err := assessmentService.CreateAndValidateAssessment(context, token, "LOGIN", 0.4)
```
An err will be nil if all conditions in the `AssessmentService.ValidateAssessment()` function are met.

### Create an option
```go
// With API key
option := option.WithAPIKey(apiKey)
// With Credentials.json (service credentials, etc.)
option := option.WithCredentialsFile(fileName)
```

### Links to related documentation
* [Create an option](https://pkg.go.dev/google.golang.org/api@v0.166.0/option)
* [Get started with recaptcha](https://cloud.google.com/security/products/recaptcha-enterprise?_ga=2.243419779.-724191075.1706884456)
* [About recaptcha score](https://cloud.google.com/recaptcha-enterprise/docs/interpret-assessment-website#:~:text=reCAPTCHA%20Enterprise%20has%2011%20levels,risk%20and%20might%20be%20fraudulent.)
* [About recaptcha actions](https://cloud.google.com/recaptcha-enterprise/docs/actions-website)
