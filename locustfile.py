"""
This file contains a Locust performance testing script for simulating user behavior of converting temperature units.
"""
# TODO:

from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    """
    A Locust User class representing a simulated user performing api calls for temperature conversion tasks.
    """
    wait_time = between(1, 5)

    @task
    def convert_temperature(self):
        """
        Task method representing the user converting temperature from Fahrenheit to Celsius.
        """
        response = self.client.post("/convert", json={"fahrenheit": 32})

        if response.status_code == 200:
            print("Conversion successful")
        else:
            print("Conversion failed")