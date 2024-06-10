# Documentation
DevOps for Data Science Team

## Repo description:
- Used and useful commands can be found at folder:
```
used_commands/
```
- To run Go API and Locust (Dockerfile-go, Dockerfile-locust):
```
helm install <DEPLOY_NAME> ./<FOLDER_NAME>
```
- To run TF Serving (Dockerile-tf-serving:
```
docker run -t --rm -p 8501:8501 -v ${PWD}/model/1:/models/saved_model -e MODEL_NAME=saved_model tensorflow/serving
```
- CI process will be automatically started after merge to **main branch**.

## Process description:
1. Basically api was provided in Golang, as well as in Flask. Whole process of CI can be found at **.github/CI_process.yaml**.
2. The problem with Flask was that it couldn't start at minikube as pod container. It worked fine run as file as well as image
pulled from Docker Hub. What's more it worked on real K8S cluster at my friends private network. And also k8s minikube cluster
infrastructure was troubleshot, it is properly configured.  
![problem-flask.png](/README_stuff/problem-flask.png)  
Proof that minikube actually worked for sample image, with same cluster config as for Flask - Nginx:  
![nginx_working.png](/README_stuff/nginx_working.png)  
Couldn't locate the exact reason of this issue.  
So I switched to Golang and it worked immediately:  
![api-working-1.png](/README_stuff/api-working-1.png)  
![api-working-2.png](/README_stuff/api-working-2.png)  
3. API was placed as pod according to task requirements with CPU scalling 1-5 pods. The service was properly configured 
with Port, TargetPort and NodePort so it can be reached outside of the minikube clsuster.
4. Whole thing was done through Helm Charts - folder **api/**.
5. It was similarly done for locust pod, but the main problem with locust - setup by Helm folder **locust/**:
```
locust-6b995667bb-nwgdt/ERROR/locust.main: Unknown User(s): locust  
```
No time to investigate.
6. Finally TensorFlow Serving, also some problems with config:
```
2024-05-16 20:18:07.668928: W tensorflow_serving/sources/storage_path/file_system_storage_path_source.cc:257] No versions of servable saved_model found under base path /models/saved_model. Did you forget to name your leaf directory as a number (eg. '/1/')?
```
Seems as related issue with weird volume mounting:  
https://stackoverflow.com/questions/45544928/tensorflow-serving-no-versions-of-servable-model-found-under-base-path  
![problem-tfs.png](/README_stuff/problem-tfs.png)

7. Folder **outputs/** should contain outputs from **gather_data.py** that listens and scrapes data from containers.
8. Folder **model** contains provided model with instruction.
9. Potentially working locust fulfills K8S architecture provided with instruction.
10. 

I would discourage using Docker and especially K8S minikube at Windows - my mistake. Virtualised host would be much better
choice with Linux OS.

## Stuff TODO / to fix / to finish:
- Locust - resolve issue **number 5** (https://docs.locust.io/en/stable/index.html). 
- Tensorflow Serving - resolve issue **number 6** (https://github.com/tensorflow/serving).
No time to investigate and fix both issues so couldn't produce report.
- PDF performance comparison

## Problem description
### DevOps for the Data Science team intern – interview tasks  
Your solution should be hosted as a new GitHub repository. Don’t send any files, only a link to
the repository.  
1. Create a web app backend (without frontend) using Python or Golang. The API call
should calculate Celsius temperature based on Fahrenheit temperature. The app should
use HTTP protocol. As a response send Celsius temperature and app identifier randomly
generated inside the app during startup.
2. Create a docker image to run the application from step 1.
3. Create a helm chart to deploy your app in the Kubernetes (k8s) environment. You can
use a locally hosted Minikube for that purpose. The app should scale automatically
based on the CPU usage metric between 1-5 replicas. The port exposed by the
kubernetes service should be configurable.
4. Use a locust as a client for your app. Build a docker image and a helm chart to run a
performance test at the k8s cluster. Run the client for 5 minutes. You can find
instructions here https://docs.locust.io/en/stable/running-in-docker.html
5. Build a continuous integration process (CI) for your project. It should contain common
actions like invoke linter to check language syntax, run tests, build docker image, push
docker image to public repo on docker hub, check helm chart syntax, etc.
6. Use TensorFlow Serving to serve a computational graph converting Fahrenheit
temperature to Celsius. The graph in a TF SavedModel format was provided along this
instruction. Repeat steps 2 – 4. Include the TF SavedModel inside a docker image.
7. Get the statistics files in cvs format and attach them to the repo under the folder named
results.
8. Compare the performance of the web app solution to the tensorflow solution. Prepare a
short report containing performance comparison between the two approaches (e.g.
which one is faster) in a PDF format. Upload the PDF to the same folder as the cvs files.

Expected architecture of your solution:  
![architecture-diagram.png](/README_stuff/architecture-diagram.png)  

We will score your solution based on these criteria:
- What % of the task you managed to implement
- Report of results
- Clean code
- Project structure
- Configurability
- Naming conventions
- Simplicity
- Comments
- Readme
