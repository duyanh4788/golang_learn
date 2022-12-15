package uploadawsservice

import (
	"github.com/gin-gonic/gin"
	"golang_01/component"
	"golang_01/modules/upload/transport/gin"
)

func UploadService(appCtx component.AppContext, router *gin.RouterGroup) error {
	router.POST("/upload-aws-s3", uploadgin.Upload(appCtx))
	return nil
}
