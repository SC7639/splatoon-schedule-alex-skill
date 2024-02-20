package main

import (
	"github.com/SC7639/splatoon-schedule-alex-skill/internal/request"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	lambda.Start(request.HandleRequest)
}
