FROM golang:1.17-alpine

RUN adduser --disabled-password --uid 1001 remindme

WORKDIR /home/remindme

USER remindme
