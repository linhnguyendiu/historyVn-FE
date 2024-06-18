package controller

import "github.com/gin-gonic/gin"

type CourseController interface {
	Create(ctx *gin.Context)
	GetByType(ctx *gin.Context)
	GetByAuthorId(ctx *gin.Context)
	GetByUserId(ctx *gin.Context)
	GetByCategory(ctx *gin.Context)
	GetTop3Course(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	UserEnrolled(ctx *gin.Context)
	GetByKeyword(ctx *gin.Context)
	GetExamScore(ctx *gin.Context)
	GetByTypeAndCategory(ctx *gin.Context)
}
