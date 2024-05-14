#! /bin/bash
curl -X POST http://localhost:5000/convert -H "Content-Type: application/json" -d '{"fahrenheit": 32}'

docker build -t openx_intern_task .

docker run -p 5000:5000 openx_intern_task

helm install openx-chart ./openx-flask-api
