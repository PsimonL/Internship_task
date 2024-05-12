FROM python:3.9-slim

WORKDIR /app

COPY openx.py /app
COPY requirements.txt /app

RUN pip install -r requirements.txt

EXPOSE 8080

ENV FLASK_APP=openx.py

CMD ["flask", "run", "--host=0.0.0.0"]