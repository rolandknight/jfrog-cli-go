package services

import (
	"encoding/json"
	"errors"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/missioncontrol/utils"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/cliutils"
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/config"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/utils/errorutils"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/utils/log"
	"net/http"
	"github.com/jfrog/jfrog-cli-go/jfrog-client/httpclient"
)

func DetachLic(service_name string, flags *DetachLicFlags) error {
	bucketId := flags.BucketId
	postContent := utils.LicenseRequestContent{
		Name:   service_name}
	requestContent, err := json.Marshal(postContent)
	if err != nil {
		return errorutils.CheckError(errors.New("Failed to marshal json. " + cliutils.GetDocumentationMessage()))
	}
	missionControlUrl := flags.MissionControlDetails.Url + "api/v3/detach_lic/buckets/" + bucketId
	httpClientDetails := utils.GetMissionControlHttpClientDetails(flags.MissionControlDetails)
	client := httpclient.NewDefaultHttpClient()
	resp, body, err := client.SendDelete(missionControlUrl, requestContent, httpClientDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return errorutils.CheckError(errors.New(resp.Status + ". " + utils.ReadMissionControlHttpMessage(body)))
	}
	log.Debug("Mission Control response: " + resp.Status)
	return nil
}

type DetachLicFlags struct {
	MissionControlDetails *config.MissionControlDetails
	Interactive           bool
	BucketId              string
}
