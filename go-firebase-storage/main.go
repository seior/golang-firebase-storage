package main

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	projectId := "blog-2c112"

	opt := option.WithCredentialsFile(path + "/config/blog-2c112-firebase-adminsdk-4igq8-24dfa40f2c.json")
	config := &firebase.Config{ProjectID: projectId, StorageBucket: "blog-2c112.appspot.com"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(err)
	}

	firebaseStorage, err := app.Storage(context.Background())
	if err != nil {
		panic(err)
	}

	bucket, err := firebaseStorage.DefaultBucket()
	if err != nil {
		panic(err)
	}

	//uploadFile(bucket, "muse.jpg", []byte(""))
	file, err := getFile(bucket, "muse.jpg")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("muse.jpg", file, 0666)
	if err != nil {
		panic(err)
	}
}

func getFile(bucket *storage.BucketHandle, filename string) ([]byte, error) {
	object, err := bucket.Object(filename).NewReader(context.Background())
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(object)
	if err != nil {
		panic(err)
	}

	return bytes, err
}

func uploadFile(bucket *storage.BucketHandle, filename string, file []byte) {
	object := bucket.Object(filename).NewWriter(context.Background())
	object.ContentType = "image/jpeg"
	object.Metadata = map[string]string{
		"created-by": "raden",
		"created-at": "2020-01-01",
	}

	//image, err := ioutil.ReadFile(path + "/muse.jpeg")
	//if err != nil {
	//	panic(err)
	//}

	_, err := object.Write(file)
	if err != nil {
		panic(err)
	}

	defer object.Close()
}
