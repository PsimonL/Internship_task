#! /bin/bash
curl -X POST http://localhost:8080/convert -H "Content-Type: application/json" -d '{"fahrenheit": 32}'

docker build -t openx_intern_task .

docker run -p 8080:8080 openx_intern_task:1.0.0

helm install openx-chart ./openx-flask-api
