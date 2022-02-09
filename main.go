package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/gorilla/mux"
)

// Informacoes a serem recebias via json
type Alerts struct {
	Channel string `json:"Channel"`
	Msg     string `json:"Msg"`
}

// Function para estabelecer conexao com sqs
func newSQS(region, endpoint string) sqsiface.SQSAPI {
	cfg := aws.Config{
		Region: aws.String(region),
	}
	if endpoint != "http://localhost.localstack.cloud:4566" {
		cfg.Endpoint = aws.String(endpoint)
	}
	sess := session.Must(session.NewSession(&cfg))
	cliSQS := sqs.New(sess)
	return cliSQS
}

//Function para realziar o envio da menssagem para a fila
func sendMessage(sqsClient sqsiface.SQSAPI, msg, queueURL string) (*sqs.SendMessageOutput, error) {
	sqsMessage := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(msg),
	}

	output, err := sqsClient.SendMessage(sqsMessage)
	if err != nil {
		return nil, fmt.Errorf("nao foi possivel enivar msg para a fila %v: %v", queueURL, err)
	}

	return output, nil
}

// Function para gerar novo alerta (menssagem sqs)
func newAlert(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Alerts
	json.Unmarshal(reqBody, &post)

	json.NewEncoder(w).Encode(post)

	newData, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
		sqsClient := newSQS(endpoints.UsEast1RegionID, "localhost.localstack.cloud:4566")
		sendMessage(sqsClient, string(newData), "http://localhost.localstack.cloud:4566/000000000000/alert")
		resp, err := http.Get("http://worker:8082/recivesqs")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
	}
}

// Function para subir o server go e receber as informacoes em json via POST
func alertPost() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/postAlert", newAlert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	alertPost()
}
