#! /bin/bash

# ============================================================================================================================================================
# REQUESTS TO TEST
curl -X POST -H "Content-Type: application/json" -d '{"fahrenheit": 32}' http://localhost:8080/convert
# Example minikube tunnel for service api
curl -X POST -H "Content-Type: application/json" -d '{"fahrenheit": 32}' http://127.0.0.1:55343/convert
curl -X GET http://localhost:8080/probe
# kubectl exec openx-api-684b578f57-nhll5 -- curl -X POST -H "Content-Type: application/json" -d '{"fahrenheit": 32}' http://localhost:5000/convert
# ============================================================================================================================================================

# ============================================================================================================================================================
# API STEPS
docker build -t openx_intern_task .
docker build -t openx_intern_task -f Dockerfile-go .
docker run -p 8080:8080 srpl/openx_intern_task:1.0.0
helm install openx-chart ./api
# Very important cause of minikube internal tunneling
minikube service openx-api
# ============================================================================================================================================================

# ============================================================================================================================================================
# LOCUST STEPS
locust -f locustfile.py --headless --host=http://localhost:8080
docker build -t openx_intern_task_locust -f Dockerfile-locust .
docker tag openx_intern_task_locust:latest srpl/openx_intern_task_locust:1.0.0
docker login
docker push openx_intern_task_locust:1.0.0
# ============================================================================================================================================================

# ============================================================================================================================================================
# TENSORFLOW SERVING STEPS
docker pull tensorflow/serving
docker run -t --rm -p 8501:8501 -v "/model/saved_model" -e MODEL_NAME=saved_model tensorflow/serving 
# ============================================================================================================================================================

