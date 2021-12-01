package admins_test

import (
	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business"
	"rulzmotoshop/business/admins"
	_adminMock "rulzmotoshop/business/admins/mocks"
	"rulzmotoshop/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockAdminRepository _adminMock.Repository
	adminService        admins.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testauth123",
		ExpiresDuration: 2,
	}

	adminService = admins.NewServiceAdmin(&mockAdminRepository, 2, jwtAuth)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := admins.Domain{
			Username: "khiyarus",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "khiyarus",
			Password: "123456",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := adminService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid test for register wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("654321")
		outputDomain := admins.Domain{
			Username: "khiyarus",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "khiyarus",
			Password: "123456",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		resp, err := adminService.Register(&inputService)
		assert.Equal(t, err, business.ErrInternalServer)
		assert.Empty(t, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test for login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := admins.Domain{
			Username: "khiyarus",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "khiyarus",
			Password: "123456",
		}

		mockAdminRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := adminService.Login(inputService.Username, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid test for login wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("23456")
		outputDomain := admins.Domain{
			Username: "khiyarus",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "khiyarus",
			Password: "123456",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		mockAdminRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := adminService.Login(inputService.Username, inputService.Password)
		assert.Equal(t, err, business.ErrEmailorPass)
		assert.Empty(t, resp)
	})
}
