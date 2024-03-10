package risk_scanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/cloud-native-security-inspector/lib/log"
	"net/http"
)

// Client to post requests
type Client struct {
	conf *Config
}

// NewClient new client
func NewClient(conf *Config) *Client {
	if conf == nil {
		conf = DefaultConfig()
	}
	return &Client{conf: conf}
}

// IsAnalyzeRunning get analyze status
func (c *Client) IsAnalyzeRunning() (bool, error) {
	requestURL := fmt.Sprintf(fmt.Sprintf("%s/%s", c.conf.Server, RouterStatus))
	log.Infof("get to: %s \n", requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	var target Status

	err = json.NewDecoder(res.Body).Decode(&target)
	if err != nil {
		return false, err
	}

	log.Infof("analyze running: %v \n", target.IsRunning)

	return target.IsRunning, nil
}

// PostAnalyze ask server to analyze resources
func (c *Client) PostAnalyze(a AnalyzeOption) error {
	requestURL := fmt.Sprintf(fmt.Sprintf("%s/%s", c.conf.Server, RouterAnalyze))
	log.Infof("post to: %s \n", requestURL)

	if jsonData, err := json.Marshal(a); err == nil {
		var request *http.Request
		request, err = http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		var response *http.Response
		response, err = client.Do(request)
		if err != nil || response.StatusCode != http.StatusCreated {
			return err
		}

		response.Body.Close()
	}

	return nil
}

func (c *Client) PostResource(a interface{}) error {
	requestURL := fmt.Sprintf(fmt.Sprintf("%s/%s", c.conf.Server, RouterResource))
	log.Infof("post to: %s \n", requestURL)

	if jsonData, err := json.Marshal(a); err == nil {
		request, _ := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		var response *http.Response
		response, err = client.Do(request)
		if err != nil || response.StatusCode != http.StatusCreated {
			return err
		}

		response.Body.Close()
	} else {
		log.Infof("json marshal err: %v \n", err)
	}

	return nil
}

// SendExitInstruction send exit instruction
func (c *Client) SendExitInstruction() error {
	requestURL := fmt.Sprintf(fmt.Sprintf("%s/%s", c.conf.Server, RouterExit))
	log.Infof("get to: %s \n", requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
