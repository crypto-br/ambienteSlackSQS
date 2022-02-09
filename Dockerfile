FROM golang:1.12.0-alpine3.9

WORKDIR /app

ADD main.go ./

# Instalando git
RUN apk add git 

# Instalando AWS-CLI (em contas aws seria usado roles)
RUN apk add --no-cache \
        python3 \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install \
        awscli \
    && rm -rf /var/cache/apk/*

# Setando configuracoes do aws-cli
RUN aws --profile default configure set aws_access_key_id mykey 
RUN aws --profile default configure set aws_secret_access_key mysecretkey
RUN aws --profile default configure set region us-east-1
RUN aws --profile default configure set output json

#Instalando deps golang
RUN go get github.com/gorilla/mux
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/sqs
RUN go get github.com/aws/aws-sdk-go/service/sqs/sqsiface

RUN go build -o main .

EXPOSE 8081

CMD ["/app/main"]

