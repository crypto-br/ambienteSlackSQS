#!/usr/bin/env bash

# Scripr para criacao de fila ao iniciar o LocalStack
awslocal --endpoint-url=http://0.0.0.0:4566 sqs create-queue --region us-east-1 --queue-name alert

