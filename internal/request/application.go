package request

import "mime/multipart"

type CreateApplicationRequest struct {
	Grade         string `form:"grade" json:"grade" binding:"required"`
	Institute     string `form:"institute" json:"institute" binding:"required"`
	Major         string `form:"major" json:"major" binding:"required"`
	Rank          string `form:"rank" json:"rank" binding:"required"`
	Group         string `form:"group" json:"group" binding:"required"`
	Intro         string `form:"intro" json:"intro" binding:"required"` //自我介绍
	RecruitmentID string `form:"recruitmentID" json:"recruitmentID" binding:"required"`
	Referrer      string `form:"referrer" json:"referrer"` //推荐人
	IsQuick       bool   `form:"isQuick" json:"isQuick"`   //速通

	Resume *multipart.FileHeader `form:"resume" json:"resume"` //简历
}
type UpApplicationRequest struct {
	Grade         string `form:"grade" json:"grade"`
	Institute     string `form:"institute" json:"institute"`
	Major         string `form:"major" json:"major"`
	Rank          string `form:"rank" json:"rank"`
	Group         string `form:"group" json:"group"`
	Intro         string `form:"intro" json:"intro"` //自我介绍
	RecruitmentID string `form:"recruitmentID" json:"recruitmentID" binding:"required"`
	Referrer      string `form:"referrer" json:"referrer"` //推荐人
	IsQuick       bool   `form:"isQuick" json:"isQuick"`   //速通

	Resume *multipart.FileHeader `form:"resume" json:"resume"` //简历
}