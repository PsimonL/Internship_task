#! /bin/bash
curl -X POST http://localhost:5000/convert -H "Content-Type: application/json" -d '{"fahrenheit": 32}'