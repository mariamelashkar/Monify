package utils

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinioClient() {
	MinioEndPoint, _ := beego.AppConfig.String("minio_endpoint")
	MinioAccessId, _ := beego.AppConfig.String("minio_accessid")
	MinioAccessKey, _ := beego.AppConfig.String("minio_accesskey")
	useSSL := false

	var err error
	MinioClient, err = minio.New(MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessId, MinioAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize Minio client: %v", err)
	}
}
