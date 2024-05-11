from flask import Flask, request, jsonify
import random

app = Flask(__name__)

app_identifier = ''.join(random.choices('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890', k=10))

@app.route('/convert', methods=['POST'])
def convert_fahrenheit_to_celsius():
    fahrenheit = request.json.get('fahrenheit')
    
    if fahrenheit is None:
        return jsonify({'error': 'Fahrenheit temperature is required'}), 400
    
    celsius = (fahrenheit - 32) * 5.0/9.0
    print("Request will be served.")
    return jsonify({'celsius': celsius, 'app_identifier': app_identifier})

if __name__ == '__main__':
    print("Server listening on port 5000...")
    app.run(debug=True)
