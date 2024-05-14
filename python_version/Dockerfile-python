FROM python:3.9-slim

WORKDIR /app

COPY openx.py /app
COPY requirements/prod.txt /app

RUN apt-get update && apt-get install -y curl

RUN pip install -r prod.txt

EXPOSE 5000

ENV FLASK_APP=openx.py

CMD ["flask", "run", "--host=0.0.0.0"]