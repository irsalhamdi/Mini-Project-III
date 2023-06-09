package customer

import (
	"customer-relationship-management/entity"
	"customer-relationship-management/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	// Create an instance of the mock CustomerRepoInterface
	mockRepo := new(mocks.CustomerRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := customerUseCaseStruct{
		customerRepository: mockRepo,
	}

	// Define the input parameter
	customer := CustomerBody{
		FirstName: "Hirsch",
		LastName:  "Surguine",
		Email:     "hsurguinex@ebay.co.uk",
		Avatar:    "http://dummyimage.com/211x100.png/ff4444/ffffff",
	}

	expectedCustomer := entity.Customer{
		FirstName: "Hirsch",
		LastName:  "Surguine",
		Email:     "hsurguinex@ebay.co.uk",
		Avatar:    "http://dummyimage.com/211x100.png/ff4444/ffffff",
	}

	// Set up the mock repository's behavior
	mockRepo.On("CreateCustomer", mock.AnythingOfType("*entity.Customer")).Return(expectedCustomer, nil)

	// Call the CreateCustomer function on the use case
	createdCustomer, err := uc.CreateCustomer(customer)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, createdCustomer)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "CreateCustomer", mock.AnythingOfType("*entity.Customer"))
}

func TestGetCustomerById(t *testing.T) {
	// Create an instance of the mock CustomerRepoInterface
	mockRepo := new(mocks.CustomerRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := customerUseCaseStruct{
		customerRepository: mockRepo,
	}

	// Define the input parameter
	id := uint(1)

	expectedCustomer := entity.Customer{
		ID:        1,
		FirstName: "Hirsch",
		LastName:  "Surguine",
		Email:     "hsurguinex@ebay.co.uk",
		Avatar:    "http://dummyimage.com/211x100.png/ff4444/ffffff",
	}

	// Set up the mock repository's behavior
	mockRepo.On("GetCustomerById", id).Return(expectedCustomer, nil)

	// Call the GetCustomerById function on the use case
	customer, err := uc.GetCustomerById(id)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "GetCustomerById", id)
}

func TestGetAllCustomer(t *testing.T) {
	// Create an instance of the mock CustomerRepoInterface
	mockRepo := new(mocks.CustomerRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := customerUseCaseStruct{
		customerRepository: mockRepo,
	}

	// Define the input parameters
	page := uint(1)
	username := "john"

	expectedPage := uint(1)
	expectedPerPage := uint(10)
	expectedTotal := 100
	expectedTotalPages := uint(10)
	expectedCustomers := []entity.Customer{
		{
			ID:        1,
			FirstName: "Hirsch",
			LastName:  "Surguine",
			Email:     "hsurguinex@ebay.co.uk",
			Avatar:    "http://dummyimage.com/211x100.png/ff4444/ffffff",
		},
		{
			ID:        2,
			FirstName: "Barny",
			LastName:  "Benneyworth",
			Email:     "bbenneyworthv@bizjournals.com",
			Avatar:    "http://dummyimage.com/220x100.png/ff4444/ffffff",
		},
	}

	// Set up the mock repository's behavior
	mockRepo.On("GetAllCustomer", page, username).Return(expectedPage, expectedPerPage, expectedTotal, expectedTotalPages, expectedCustomers, nil)

	// Call the GetAllCustomer function on the use case
	p, pp, total, totalPages, customers, err := uc.GetAllCustomer(page, username)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedPage, p)
	assert.Equal(t, expectedPerPage, pp)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedTotalPages, totalPages)
	assert.Equal(t, expectedCustomers, customers)

	// Assert that the expected repository method was called with the correct arguments
	mockRepo.AssertCalled(t, "GetAllCustomer", page, username)
}
func TestUpdateCustomerById(t *testing.T) {
	// Create an instance of the mock CustomerRepoInterface
	mockRepo := new(mocks.CustomerRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := customerUseCaseStruct{
		customerRepository: mockRepo,
	}

	// Define the input parameters
	id := uint(1)
	customer := UpdateCustomerBody{
		FirstName: "Hirsch",
		LastName:  "Surguine",
		Avatar:    "http://dummyimage.com/211x100.png/ff4444/ffffff",
	}

	expectedCustomer := entity.Customer{
		ID:        1,
		FirstName: "Fabe",
		LastName:  "Smitherman",
		Avatar:    "fsmithermanw@rakuten.co.jp",
	}

	// Set up the mock repository's behavior
	mockRepo.On("UpdateCustomerById", id, mock.AnythingOfType("*entity.Customer")).Return(expectedCustomer, nil)

	// Call the UpdateCustomerById function on the use case
	updatedCustomer, err := uc.UpdateCustomerById(id, customer)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, updatedCustomer)

	// Assert that the expected repository method was called with the correct arguments
	mockRepo.AssertCalled(t, "UpdateCustomerById", id, mock.AnythingOfType("*entity.Customer"))
}

func TestDeleteCustomerById(t *testing.T) {
	// Create an instance of the mock CustomerRepoInterface
	mockRepo := new(mocks.CustomerRepoInterface)

	// Create an instance of the use case with the mock repository
	uc := customerUseCaseStruct{
		customerRepository: mockRepo,
	}

	// Define the input parameter
	id := uint(1)

	// Set up the mock repository's behavior
	mockRepo.On("DeleteCustomerById", id).Return(nil)

	// Call the DeleteCustomerById function on the use case
	err := uc.DeleteCustomerById(id)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected repository method was called with the correct argument
	mockRepo.AssertCalled(t, "DeleteCustomerById", id)
}
