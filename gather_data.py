"""
This file contains a Locust performance testing script for simulating user behavior of converting temperature units.
"""

import docker
import time

OUTPUT_FILE = "/output/tf_serving.txt"


def stream_logs(container):
    """
    Streams logs from the specified Docker container.
    """
    for log in container.logs(stream=True):
        yield log.decode('utf-8')


def write_logs_to_file(logs):
    """
    Writes logs to the specified output file.
    """
    with open(OUTPUT_FILE, "a") as f:
        for log in logs:
            f.write(log)


def main(container_name):
    """
    Main function to start the process.
    """
    client = docker.from_env()
    try:
        container = client.containers.get(container_name)
        while True:
            logs = stream_logs(container)
            write_logs_to_file(logs)
            time.sleep(3)
    except docker.errors.NotFound:
        print("Kontener o podanej nazwie nie został znaleziony.")


if __name__ == "__main__":
    main(container_name="tensorflow/serving")
