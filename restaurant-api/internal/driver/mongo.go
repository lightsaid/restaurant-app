package driver

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect 连接
func Connect() (*mongo.Client, error) {
	// 创建链接参数
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	opts.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	// 链接
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Error connecting: ", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Close 关闭
func Close(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// 关闭链接
	if err := client.Disconnect(ctx); err != nil {
		log.Panic(err)
	}
}
