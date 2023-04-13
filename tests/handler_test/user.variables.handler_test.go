package handler_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"fmt"

	"github.com/bxcodec/faker/v4"
)

func generateUserForm() *forms.Signup {
	return &forms.Signup{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Password:    faker.Password(),
	}
}

func createAndRegisterUser(user *forms.Signup) *models.User {
	if user == nil {
		user = generateUserForm()
	}

	resultUser, err := userSrv.Add(user)
	if err != nil {
		panic(err)
	}

	return resultUser.User
}

func generateLoginForm(user *forms.Signup) *forms.Login {
	if user == nil {
		user = generateUserForm()

		_ = createAndRegisterUser(user)
	}

	return &forms.Login{
		Email:    user.Email,
		Password: user.Password,
	}
}

func loginUserAndGenerateAuth(login *forms.Login) string {
	if login == nil {
		login = generateLoginForm(nil)
	}

	auth, err := userSrv.Login(login)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Bearer %v", auth.AccessToken)
}
