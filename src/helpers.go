package dataparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ReadJSONFile reads a JSON file and returns its contents as a map.
func ReadJSONFile(filePath string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var fileData map[string]interface{}
	err = json.Unmarshal(data, &fileData)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

// DownloadFileFromS3 downloads a file from S3 and saves it locally.
func DownloadFileFromS3(s3Session *session.Session, bucket string, objectKey string, filePath string) error {
	_, err := s3.New(s3Session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
outfile, err := filepath.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
_, err = io.Copy(outfile, resp.Body)
outfile.Close()
return err
}

// HttpClient makes an HTTP GET request and returns the response body.
func HttpClient(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// ReadCSVFile reads a CSV file and returns its contents as a map.
func ReadCSVFile(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := csv.NewReader(file)
	rows, err := scanner.ReadAll()
	if err != nil {
		return nil, err
	}
	csvData := map[string][]string{}
	for i, row := range rows {
		if i == 0 {
			csvData["columns"] = row
			continue
		}
		csvData[strings.Join(row[:], ",")] = row
	}
	return csvData, nil
}

// UploadFileToS3 uploads a local file to S3.
func UploadFileToS3(s3Session *session.Session, bucket string, filePath string, objectKey string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = s3.New(s3Session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(data),
	})
	return err
}