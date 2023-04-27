package adapters

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/user"
)

func GetUser(context *events.APIGatewayProxyRequestContext) (*user.AuthorizedUser, error) {
	claims := context.Authorizer["claims"].(map[string]interface{})
	userId, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return nil, err
	}
	userEmail := claims["email"].(string)
	return &user.AuthorizedUser{
		Id:    userId,
		Email: userEmail,
	}, nil
}
