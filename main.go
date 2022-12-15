package main

import (
	"golang_01/component/uploadprovider"
	"golang_01/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dns := os.Getenv("MYSQL")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("s3Region")
	s3Apikey := os.Getenv("s3Apikey")
	s3SecretKey := os.Getenv("s3SecretKey")
	s3SDomain := os.Getenv("s3SDomain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3Apikey, s3SecretKey, s3SDomain)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Can not connect DB")
	}

	log.Println("Connect to", db)

	if err := mainrouter.MainServices(db, s3Provider); err != nil {
		log.Println(err)
	}
}
