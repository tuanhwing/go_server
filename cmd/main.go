package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goter.com.vn/server/config"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/infrastructure/repository"
	"goter.com.vn/server/pkg/metric"
	"goter.com.vn/server/usecase/user"
)

func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Invalid query")
	}
	return os.Args[1], nil
}

func main() {

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	//CLI
	appMetric := metric.NewCLI("search")
	appMetric.Started()
	query, err := handleParams()
	if err != nil {
		log.Fatal(err.Error())
	}

	dataSourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(dataSourceName).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserMongoDB(client)
	userService := user.NewService(userRepo)

	id, _ := entity.StringToID(query)
	user, err := userService.GetUser(id)
	fmt.Printf("\n\n****** ID | firstname | lastname | email | createAt | updateAt ******\n\n")
	if user != nil {
		fmt.Printf("------------------------------[START]User \"id\" == \"%s\"------------------------------\n", query)
		fmt.Printf("%s\n", user.String())
		fmt.Printf("------------------------------[ END ]User \"id\" == \"%s\"------------------------------\n", query)

	}

	all, err := userService.SearchUsers(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("------------------------------[START]User \"firtname\" || \"lastname\" || \"email\" == \"%s\"------------------------------\n", query)
	for _, j := range all {
		fmt.Printf("%s\n", j.String())
	}
	fmt.Printf("------------------------------[ END ]User \"lastname\" || \"firtname\" || \"email\" == \"%s\"------------------------------\n", query)

	appMetric.Finished()
	err = metricService.SaveCLI(appMetric)
	if err != nil {
		log.Fatal(err.Error())
	}
}
