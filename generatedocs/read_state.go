package generate_docs

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//WhichCloudState uses a case statement to evaluate providers and pull in remote terraform state from the appropriate location
func WhichCloudState(s string, st StateProviders) (CloudState, error) {
	var state CloudState
	switch s {
	case "azure":
		az, err := GetAzState(st)
		if err != nil {
			return az, fmt.Errorf(err.Error())
		}
		return az, nil
	case "aws":
		amazon, err := GetAwsState(st)
		if err != nil {
			return amazon, fmt.Errorf(err.Error())
		}
		return amazon, nil
	case "gcp":
		gcp, err := GetGcpState(st)
		if err != nil {
			return gcp, fmt.Errorf(err.Error())
		}
		return gcp, nil
	case "tfc":
		tfc, err := GetTfcState(st)
		if err != nil {
			return tfc, fmt.Errorf(err.Error())
		}
		return tfc, nil
	}
	return state, nil
}
//GetTfcState takes a TFC api token as well as the workspace name and organization name and returns the current version of the state file
func GetTfcState(s StateProviders) (CloudState,error) {
	var state TfcState
	var n CloudState

	u := fmt.Sprintf("https://app.terraform.io/api/v2/workspaces/%s/current-state-version", s.TerraformCloud.WorkspaceId)

	token := s.TerraformCloud.ApiToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return n, fmt.Errorf("failed to create GET request for http.NewRequest make sure API token for workspace with id: %s is not expired", s.TerraformCloud.WorkspaceId)
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return n, fmt.Errorf("failed to do request GET for %s check url", request.RequestURI)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return n, fmt.Errorf("failed to read body for %s", resp.Request.URL)
	}

	err = json.Unmarshal(response, &state)
	if err != nil {
		return n, fmt.Errorf("failed to unmarshal JSON for %s response, please check type State for formatting", resp.Request.URL)
	}

	n, err = ManipulateTfcState(state)
	if err != nil {
		return n, fmt.Errorf("failed to unmarshal JSON for %s response, please check type State for formatting", resp.Request.URL)
	}

	return n, nil
}

//ManipulateTfcState takes an argument of TfcState and returns cloud state by doing a GET request on the hosted state download url object
func ManipulateTfcState(t TfcState)(CloudState,error){
	var s CloudState
	u := fmt.Sprintf("%s", t.Data.Attributes.HostedStateDownloadUrl)

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return s, fmt.Errorf(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return s, fmt.Errorf(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(response, &s)
	if err != nil {
		return s, fmt.Errorf("failed to unmarshal JSON for %s response, please check type State for formatting", resp.Request.URL)
	}

	return s, nil
}

//GetAzState takes inputs including storage account name, container name, blob name, and SAS token to download a blob(tfstate file) for use later on
func GetAzState(s StateProviders) (CloudState, error) {
	var state CloudState

	u := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s?%s", s.Azure.StorageAccountName, s.Azure.ContainerName, s.Azure.BlobName, s.Azure.SasToken)


	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return state, fmt.Errorf("failed to create GET request for http.NewRequest make sure SAS token for blob with name  %s is not expired", s.Azure.BlobName)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return state, fmt.Errorf("failed to do request GET for %s check url", request.RequestURI)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return state, fmt.Errorf("failed to read body for %s", resp.Request.URL)
	}

	err = json.Unmarshal(response, &state)
	if err != nil {
		return state, fmt.Errorf("failed to unmarshal JSON for %s response, please check type State for formatting", resp.Request.URL)
	}

	return state, nil
}

//GetAwsState takes inputs of type StateProviders, and utilizes the AWS Go SDK to download state file from s3 buckets and then remove the state file once complete
func GetAwsState(s StateProviders) (CloudState,error) {
	var state CloudState
	err := os.Setenv("AWS_ACCESS_KEY_ID", s.Aws.AccessKey)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}
	err = os.Setenv("AWS_SECRET_ACCESS_KEY", s.Aws.SecretAccessKey)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}
	err = os.Setenv("AWS_SESSION_TOKEN", s.Aws.SessionToken)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}


	bucket := s.Aws.BucketName
	item := s.Aws.Object

	file, err := os.Create(item)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	sess, _ := session.NewSession(&aws.Config{Region: aws.String(s.Aws.Region)})
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}

	f , err := ioutil.ReadFile(item)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}

	errState := json.Unmarshal(f, &state)
	if errState != nil {
		return state, fmt.Errorf(err.Error())
	}

	err = os.Remove(item)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}

	return state, nil
}

//GetGcpState takes inputs of type StateProviders, and utilizes the Google Cloud Go SDK to download state file from s3 buckets and then remove the state file once complete
func GetGcpState(s StateProviders) (CloudState, error) {
	var state CloudState
	u := fmt.Sprintf("https://storage.googleapis.com/storage/v1/b/%s/o/%s?alt=media", s.GoogleCloud.BucketName, s.GoogleCloud.ObjectName)

	bearer := "Bearer " + s.GoogleCloud.Oauth2Token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return state, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(response, &state)

	return state, nil
}

