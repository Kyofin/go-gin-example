package util

import (
	"context"
	"fmt"
	"github.com/huzk8/go-gin-example/pkg/setting"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

func mino() {
	endpoint := setting.ENDPOINT
	accessKeyID := setting.ACCESS_KEY
	secretAccessKey := setting.ACCESS_SECRET
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	file, err := os.Open("C:\\Users\\10400\\GolandProjects\\go-gin-example\\conf\\app.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), "myjfs", "myobject", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
}
