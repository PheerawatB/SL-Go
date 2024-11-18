package main

import (
	"comsumer/repositories"
	"comsumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initialDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"), viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.name"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	// 	viper.GetString("db.username"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.host"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.name"))

	// dial := mysql.Open(dsn)

	// db, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	// if err != nil {
	// 	panic(err)
	// }
	return db
}

func main() {

	// fmt.Printf("print 1st %v:%v", viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"))
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	db := initialDatabase()
	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Printf("Acoount consumer started...\n")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}
}

// func main() {
// 	server := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(server, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer consumer.Close()
// 	partitionConsumer, err := consumer.ConsumePartition("frankhello", 0, sarama.OffsetOldest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partitionConsumer.Close()
// 	fmt.Println("Consumer started")
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Printf("Error: %v\n", err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Printf(string(msg.Value))
// 		}
// 	}
// }
