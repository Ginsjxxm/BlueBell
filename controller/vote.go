package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParams, errData)
		return
	}
	UserID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	err = logic.VoteForPost(UserID, p)
	if err != nil {
		ResponseErrorWithMsg(c, 0, err.Error())
		return
	}
	ResponseSuccess(c, nil)
	return
}
