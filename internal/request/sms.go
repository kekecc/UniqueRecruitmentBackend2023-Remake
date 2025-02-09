package request

import (
	"UniqueRecruitmentBackend/internal/constants"
)

type SendSMS struct {
	Type      constants.SMSType `json:"type" binding:"required"`    // the candidate status : Pass or Fail
	Current   string            `json:"current" binding:"required"` // the application current step
	Next      string            `json:"next" binding:"required"`    // the application next step
	Time      string            `json:"time"`                       // the next step(interview/test) time
	Place     string            `json:"place"`                      // the next step(interview/test) place
	MeetingId string            `json:"meetingId"`
	Rest      string            `json:"rest"`
	Aids      []string          `json:"aids"` // the applications will be sent sms
}
