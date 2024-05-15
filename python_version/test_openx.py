"""
Unit tests for the openx module.
"""

import unittest
from openx import app


class TestOpenx(unittest.TestCase):
    """
    Test cases for the openx module.
    """

    def setUp(self):
        """
        Set up the test client.
        """
        self.app = app.test_client()

    def test_convert_fahrenheit_to_celsius_success(self):
        """
        Test conversion of Fahrenheit to Celsius with successful response.
        """
        response = self.app.post('/convert', json={'fahrenheit': 32})
        data = response.get_json()
        self.assertEqual(response.status_code, 200)
        self.assertIn('celsius', data)
        self.assertIn('app_identifier', data)

    def test_convert_32f_to_0c(self):
        """
        Test conversion of 32°F to 0°C.
        """
        response = self.app.post('/convert', json={'fahrenheit': 32})
        data = response.get_json()
        self.assertEqual(data['celsius'], 0)

    def test_convert_212f_to_100c(self):
        """
        Test conversion of 212°F to 100°C.
        """
        response = self.app.post('/convert', json={'fahrenheit': 212})
        data = response.get_json()
        self.assertEqual(data['celsius'], 100)


if __name__ == '__main__':
    unittest.main()