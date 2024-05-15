from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 3) 

    @task
    def convert_temperature(self):
        response = self.client.post("/convert", json={"fahrenheit": 32})

        if response.status_code == 200:
            print("Conversion successful")
        else:
            print("Conversion failed")