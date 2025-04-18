package common

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type EmailService struct {
	sesClient *sesv2.Client
}

func NewEmailService() (*EmailService, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ap-south-1"),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	// Create SES client
	sesClient := sesv2.NewFromConfig(cfg)

	return &EmailService{
		sesClient: sesClient,
	}, nil
}

type EmailType struct {
	Body       string
	Subject    string
	Recepients []string
}

func (s *EmailService) SendEmail(email EmailType) error {
	charset := aws.String("UTF-8")
	// Create email input
	emailContent := &types.EmailContent{
		Simple: &types.Message{
			Subject: &types.Content{
				Data:    aws.String(email.Subject),
				Charset: charset,
			},
			Body: &types.Body{
				Html: &types.Content{
					Data:    aws.String(email.Body),
					Charset: charset,
				},
			},
		},
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String("sender@example.com"),
		Destination: &types.Destination{
			ToAddresses: email.Recepients,
		},
		Content: emailContent,
	}

	fmt.Print("---------------------------------------input")
	// fmt.Print(email.Body)
	fmt.Print(input)

	// Send email
	// _, err := s.sesClient.SendEmail(context.Background(), input)
	// if err != nil {
	// 	return fmt.Errorf("failed to send email: %v", err)
	// }

	return nil
}
