package getorganization

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

const (
	ivConnection = "WebexConnection"

	ivorgid     = "orgID"
	ovMessageId = "responseBody"
)

type GetOrgActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &t{metadata: metadata}
}

func (a *t) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *t) Eval(context activity.Context) (done bool, err error) {

	if context.GetInput(ivorgid) == nil {
		return false, activity.NewError("Organization id is required ", "Webex-GETCUSTOMER-4002", nil)
	}

	//Read connection details
	connectionInfo, _ := data.CoerceToObject(context.GetInput(ivConnection))

	if connectionInfo == nil {
		return false, activity.NewError("Webex connection is not configured", "Webex-GETCUSTOMER-4001", nil)
	}

	var ConsumerKey string
	var ConsumerSecret string

	connectionSettings, _ := connectionInfo["settings"].([]interface{})
	if connectionSettings != nil {
		for _, v := range connectionSettings {
			setting, _ := data.CoerceToObject(v)
			if setting != nil {
				if setting["name"] == "ConsumerKey" {
					ConsumerKey, _ = data.CoerceToString(setting["value"])
				} else if setting["name"] == "ConsumerSecret" {
					ConsumerSecret, _ = data.CoerceToString(setting["value"])
				}

			}
		}
	}

	var orgid = context.GetInput(ivorgid).(string)

	url := "https://api.ciscospark.com/v1/organizations/" + orgid

	var mod = (ConsumerKey + ":" + ConsumerSecret)
	sEnc := base64.StdEncoding.EncodeToString([]byte(mod))

	var auth = "Basic " + sEnc

	urls := "https://api.ciscospark.com/"

	payload := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", urls, payload)

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	a1 := []rune(string(body))

	accesstoken := string(a1[17:53])

	var bearer = "Bearer " + accesstoken

	reqs, _ := http.NewRequest("GET", url, nil)

	reqs.Header.Add("Authorization", bearer)

	resp, _ := http.DefaultClient.Do(reqs)

	defer res.Body.Close()
	bodyre, _ := ioutil.ReadAll(resp.Body)

	//Set Message ID in the output

	context.SetOutput(ovMessageId, string(bodyre))

	return true, nil
}
