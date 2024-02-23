# About package

This package implements google recaptcha code verification.

Usage
```go
assessment, err := gorecaptcha.CreateAndValidateAssesment(context.Background(), option.WithAPIKey("<api-key>"), appId, recaptchaToken, tokenFromFrontend, action, minScore)
```

Links to related documentation
* [Create an option](https://pkg.go.dev/google.golang.org/api@v0.166.0/option)
* [Get started with recaptcha](https://cloud.google.com/security/products/recaptcha-enterprise?_ga=2.243419779.-724191075.1706884456)
* [About recaptcha score](https://cloud.google.com/recaptcha-enterprise/docs/interpret-assessment-website#:~:text=reCAPTCHA%20Enterprise%20has%2011%20levels,risk%20and%20might%20be%20fraudulent.)
* [About recaptcha actions](https://cloud.google.com/recaptcha-enterprise/docs/actions-website)
