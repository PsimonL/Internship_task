Invoke-RestMethod -Uri http://localhost:8080/convert -Method Post -ContentType "application/json" -Body '{"fahrenheit": 32}'
Invoke-RestMethod -Uri http://localhost:8080/probe -Method Get

docker build -t openx_intern_task .
docker build -t openx_intern_task -f Dockerfile-go .

docker run -p 8080:8080 openx_intern_task

helm install openx-chart ./openx-flask-api

minikube service openx-api
