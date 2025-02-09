package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Post recruitment/
// admin role

// CreateRecruitment create new recruitment
func CreateRecruitment(c *gin.Context) {
	var req request.CreateRecruitment
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, error2.RequestBodyError.WithDetail(err.Error()))
		return
	}
	if time.Now().After(req.Beginning) || req.Beginning.After(req.Deadline) || req.Deadline.After(req.End) {
		common.Error(c, error2.RequestBodyError.WithDetail("time set up wrong"))
		return
	}
	recruitmentId, err := models.CreateRecruitment(&req)
	if err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("recruitment"))
		return
	}
	common.Success(c, "Success create recruitment", map[string]interface{}{
		"rid": recruitmentId,
	})
}

// Post recruitment/:rid/schedule
// admin role

// UpdateRecruitment update recruitment details
func UpdateRecruitment(c *gin.Context) {
	recruitmentId := c.Param("rid")
	if recruitmentId == "" {
		common.Error(c, error2.RequestBodyError.WithDetail("recruitment id is null"))
		return
	}
	var req request.UpdateRecruitment
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, error2.RequestBodyError.WithDetail(err.Error()))
		return
	}
	if err := models.UpdateRecruitment(recruitmentId, &req); err != nil {
		common.Error(c, error2.UpdateDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success update recruitment", map[string]interface{}{
		"rid": recruitmentId,
	})
}

// Get recruitment/:rid

// GetRecruitmentById get recruitment details by id
// member can only get the recruitments' detail after join in
func GetRecruitmentById(c *gin.Context) {
	recruitmentId := c.Param("rid")
	if recruitmentId == "" {
		common.Error(c, error2.RequestBodyError.WithDetail("lost http query params [rid]"))
		return
	}
	userInfo, err := getUserInfoByUID(c, common.GetUID(c))
	if err != nil {
		common.Error(c, error2.CheckPermissionError.WithDetail(err.Error()))
		return
	}

	if utils.CheckRoleByUserDetail(userInfo, constants.Admin) {
		resp, err := models.GetRecruitmentById(recruitmentId, constants.Admin)
		if err != nil {
			common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
			return
		}
		common.Success(c, "Success get recruitment by admin role", resp)
	} else if utils.CheckRoleByUserDetail(userInfo, constants.MemberRole) {
		resp, err := models.GetRecruitmentById(recruitmentId, constants.MemberRole)
		if err != nil {
			common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
			return
		}

		//TODO(wwb)
		//Compare member join in time and recruitment time
		// if compareTime(resp.Beginning.String(), userInfo.) {

		// }

		common.Success(c, "Success get recruitment by member role", resp)
	} else {
		resp, err := models.GetRecruitmentById(recruitmentId, constants.CandidateRole)
		if err != nil {
			common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
			return
		}
		common.Success(c, "Success get recruitment by candidate role", resp)
	}
}

// Get recruitment/

// GetAllRecruitment get all recruitment details
func GetAllRecruitment(c *gin.Context) {
	// TODO(wwb)
	// compare member join in time and recruitment time
	resp, err := models.GetAllRecruitment()
	if err != nil {
		common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success get all recruitment", resp)
}

// Get recruitment/:rid/pending

// GetPendingRecruitment get pending recruitment details
func GetPendingRecruitment(c *gin.Context) {
	role, err := getUserRoleByUID(c, common.GetUID(c))
	if err != nil {
		common.Error(c, error2.CheckPermissionError.WithDetail(err.Error()))
		return
	}
	resp, err := models.GetPendingRecruitment(role)
	if err != nil {
		common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success get pending recruitment", resp)
}

// func compareTime(a string, b string) bool {
// 	ta := utils.TimeParse(a)
// 	tb := utils.TimeParse(b)
// 	return ta.After(tb)
// }
