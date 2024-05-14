import random
from fastapi import FastAPI, HTTPException

app = FastAPI()

CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
APP_IDENTIFIER = ''.join(random.choices(CHARS, k=10))


@app.post('/convert')
async def convert_fahrenheit_to_celsius(fahrenheit: float):
    """
    Converts Fahrenheit temperature to Celsius.

    Receives JSON data with Fahrenheit temperature in the request body.
    Returns JSON response with the converted Celsius temperature.
    """
    if fahrenheit is None:
        raise HTTPException(status_code=400, detail='Fahrenheit temperature is required')

    celsius = (fahrenheit - 32) * 5.0 / 9.0
    print("Request will be served.")
    return {'celsius': celsius, 'app_identifier': APP_IDENTIFIER}


@app.get('/probe')
async def probes_test():
    """
    Satisfies K8S health, probe system.
    """
    return {'message': 'K8s Probe request'}


if __name__ == '__main__':
    import uvicorn
    print("Server listening on port 5000...")
    import fastapi

    print("Wersja FastAPI:", fastapi.__version__)
    print("Wersja uvicorn:", uvicorn.__version__)
    uvicorn.run(app, host="0.0.0.0", port=5000, debug=True)
