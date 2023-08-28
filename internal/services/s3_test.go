package services

import (
	"encoding/json"
	"fireteam-core/internal/testhelper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockS3Repository struct {
	PutFunc func(fileName string, data []byte) error
	GetFunc func(fileName string) ([]byte, error)
}

func (m *MockS3Repository) Put(fileName string, data []byte) error {
	if m.PutFunc != nil {
		return m.PutFunc(fileName, data)
	}
	return nil
}

func (m *MockS3Repository) Get(fileName string) ([]byte, error) {
	if m.GetFunc != nil {
		return m.GetFunc(fileName)
	}
	return nil, nil
}

func TestPutData_FailedToPut_ShouldReturnError(t *testing.T) {
	testCasesInJson, err := testhelper.LoadJsonFromFile("./testcases/savefile_sample_file.json")
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	var testCases []struct {
		Name     string `json:"name"`
		FileName string `json:"file_name"`
		FileSize int    `json:"file_size"`
		FileType string `json:"file_type"`
		Data     []byte `json:"file_data"`
	}
	err = json.Unmarshal(testCasesInJson, &testCases)
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	mockRepo := &MockS3Repository{
		PutFunc: func(fileName string, data []byte) error {
			return fmt.Errorf("error when writing object to S3")
		},
		GetFunc: func(fileName string) ([]byte, error) {
			return nil, nil
		},
	}
	s3Service, _ := NewS3ServiceImpl(mockRepo)
	err = s3Service.Put(testCases[0].FileName, testCases[0].Data)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	assert.EqualError(t, err, fmt.Sprintf("error saving data file - %s: error when writing object to S3", testCases[0].FileName))
}

func TestGetData_FailedToGet_ShouldReturnError(t *testing.T) {
	testCasesInJson, err := testhelper.LoadJsonFromFile("./testcases/savefile_sample_file.json")
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	var testCases []struct {
		Name     string `json:"name"`
		FileName string `json:"file_name"`
		FileSize int    `json:"file_size"`
		FileType string `json:"file_type"`
		Data     []byte `json:"file_data"`
	}
	err = json.Unmarshal(testCasesInJson, &testCases)
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}

	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	mockRepo := &MockS3Repository{
		PutFunc: func(fileName string, data []byte) error {
			return nil
		},
		GetFunc: func(fileName string) ([]byte, error) {
			return nil, fmt.Errorf("error when getting object from S3")
		},
	}
	s3Service, _ := NewS3ServiceImpl(mockRepo)
	_, err = s3Service.Get(testCases[0].FileName)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	assert.EqualError(t, err, fmt.Sprintf("error getting data for file - %s: error when getting object from S3", testCases[0].FileName))
}

func TestS3_NoRepositoryPassedToService_ShouldReturnError(t *testing.T) {
	_, err := NewS3ServiceImpl(nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	assert.EqualError(t, err, "file repository is required to initilize s3 service")
}
