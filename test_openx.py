import unittest
from openx import app


class TestOpenx(unittest.TestCase):
    def setUp(self):
        self.app = app.test_client()

    def test_convert_fahrenheit_to_celsius_success(self):
        response = self.app.post('/convert', json={'fahrenheit': 32})
        data = response.get_json()
        self.assertEqual(response.status_code, 200)
        self.assertIn('celsius', data)
        self.assertIn('app_identifier', data)

    def test_convert_32F_to_0C(self):
        response = self.app.post('/convert', json={'fahrenheit': 32})
        data = response.get_json()
        self.assertEqual(data['celsius'], 0)


if __name__ == '__main__':
    unittest.main()
