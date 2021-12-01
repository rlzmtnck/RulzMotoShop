package sellers_test

import (
	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business"
	"rulzmotoshop/business/sellers"
	_sellersMock "rulzmotoshop/business/sellers/mocks"
	"rulzmotoshop/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mocksellerRepository _sellersMock.Repository
	sellersService       sellers.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testauth123",
		ExpiresDuration: 2,
	}

	sellersService = sellers.NewServiceSeller(&mocksellerRepository, 2, jwtAuth)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     password,
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     "123456",
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		mocksellerRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := sellersService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("654321")
		outputDomain := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     password,
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     "123456",
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		mocksellerRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		resp, err := sellersService.Register(&inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("test case 3, invalid register duplicate", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     password,
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := sellers.Domain{
			Username:     "rulz",
			Email:        "rulz@mail.com",
			Password:     "123456",
			Name:         "Muhammad Khiyarus",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}

		mocksellerRepository.On("Register", mock.Anything).Return(sellers.Domain{}, business.ErrDuplicateData).Once()

		resp, err := sellersService.Register(&inputService)

		assert.NotNil(t, err)
		assert.NotEqual(t, outputDomain, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := sellers.Domain{
			Email:    "test@mail.com",
			Password: password,
		}
		inputService := sellers.Domain{
			Email:    "test@mail.com",
			Password: "123456",
		}
		mocksellerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := sellersService.Login(inputService.Email, inputService.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("654321")
		outputDomain := sellers.Domain{
			Email:    "test@mail.com",
			Password: password,
		}
		inputService := sellers.Domain{
			Email:    "teasst@mail.com",
			Password: "123456",
		}
		mocksellerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := sellersService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrEmailorPass)
	})

	t.Run("test case 4, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("6127376")
		outputDomain := sellers.Domain{
			Email:    "test@mail.com",
			Password: password,
		}
		inputService := sellers.Domain{
			Email:    "test@mail.com",
			Password: "123456",
		}
		mocksellerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := sellersService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrEmailorPass)
	})
}
