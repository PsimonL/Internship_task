"""
This module defines a Flask application for converting Fahrenheit temperature to Celsius.
"""

import random
from flask import Flask, request, jsonify

app = Flask(__name__)

chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
APP_IDENTIFIER = ''.join(random.choices(chars, k=10))


@app.route('/convert', methods=['POST'])
def convert_fahrenheit_to_celsius():
    """
    Converts Fahrenheit temperature to Celsius.

    Receives JSON data with Fahrenheit temperature in the request body.
    Returns JSON response with the converted Celsius temperature.
    """
    fahrenheit = request.json.get('fahrenheit')

    if fahrenheit is None:
        return jsonify({'error': 'Fahrenheit temperature is required'}), 400

    celsius = (fahrenheit - 32) * 5.0 / 9.0
    print("Request will be served.")
    return jsonify({'celsius': celsius, 'app_identifier': APP_IDENTIFIER})


if __name__ == '__main__':
    print("Server listening on port 8080...")
    app.run(debug=True, port=8080)
