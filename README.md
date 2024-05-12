# Documentation

# Problem Description
## DevOps for the Data Science team intern – interview tasks  
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