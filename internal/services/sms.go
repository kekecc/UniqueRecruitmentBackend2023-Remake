package services

import (
	"UniqueRecruitmentBackend/configs"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SMSType string

const (
	RegisterCode      = "RegisterCode"
	ResetPasswordCode = "ResetPasswordCode"
)

type SMSBody struct {
	Phone      string   `json:"phone_number"`
	TemplateID uint     `json:"template_id"`
	Params     []string `json:"template_param_set"`
}

func SendSMS(smsBody SMSBody) (*http.Response, error) {
	body, err := json.Marshal(smsBody)
	if err != nil {
		log.Println("marshal: ", err)
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://open.hustunique.com/sms/send_single",
		bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("AccessKey", configs.Config.SMS.Token)
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
