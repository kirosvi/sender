#!/bin/bash

template=test

curl \
    -H "Content-Type: application/json" \
    -X POST http://127.0.0.1:8080/alert/ppcore \
    -d '@alert_example'
#   -d '{"Name":"Alice","Body":"Hello","options":{"Time":1294706395881547000,"start":"opt","stop":"retry"}}'
curl \
    -H "Content-Type: application/json" \
    -X POST http://127.0.0.1:8080/alert/rundeck \
    -d '{"trigger": "failure","status": "failed","executionId": 5102,"execution": {"id": 5102,"href": "http://rundeck.infra.ppdev.ru/project/DevOps/execution/show/5102","permalink": null,"status": "failed","project": "DevOps","executionType": "user","user": "admin","date-started": {"unixtime": 1669198910364,"date": "2022-11-23T10:21:50Z"},"job": {"id": "6cd8db31-82a5-49d3-a6ce-ea147755045a","name": "test-tg","group": "test","project": "DevOps","description": "","href": "http://rundeck.infra.ppdev.ru/api/41/job/6cd8db31-82a5-49d3-a6ce-ea147755045a","permalink": "http://rundeck.infra.ppdev.ru/project/DevOps/job/show/6cd8db31-82a5-49d3-a6ce-ea147755045a"},"description": "exit 1","argstring": null,"serverUUID": "a14bc3e6-75e8-4fe4-a90d-a16dcc976bf6"}}'

curl \
    -H "Content-Type: application/json" \
    -X POST http://127.0.0.1:8080/alert/infra \
    -d '{"data":"test message"}'

curl \
    -H "Content-Type: application/json" \
    -X GET http://127.0.0.1:8080/ping
