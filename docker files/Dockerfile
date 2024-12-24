FROM golang:1.23 as builder

WORKDIR /app

RUN ["go", "build", "-o", "tickets_api", "./..."]

