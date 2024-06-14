import random
import urllib.parse
import requests

# Example lists for generating random supermarket data
supermarket_names = ["Lidl", "Aldi", "Edeka", "Rewe", "Penny", "Netto", "Kaufland"]
streets = ["Hauptstraße", "Blumenweg", "Bahnhofstraße", "Bergstraße", "Gartenstraße", "Waldweg", "Mühlenweg", "Kirchstraße"]
postal_codes = ["78462", "78464", "78465", "78467"]
city = "Konstanz"
house_numbers = list(range(1, 101))

token = "token"
headers = {
    "Authorization": f"Bearer {token}"
}

def generate_supermarket():
    name = random.choice(supermarket_names)
    street = random.choice(streets)
    house_number = random.choice(house_numbers)
    postal_code = random.choice(postal_codes)
    
    query_params = {
        "name": name,
        "street": street,
        "postal_code": postal_code,
        "city": city,
        "houseNumber": house_number
    }
    
    url = "http://localhost:8083/addSupermarket?" + urllib.parse.urlencode(query_params)
    return url

# Generate and send 10 supermarket URLs
for _ in range(10):
    supermarket_url = generate_supermarket()
    response = requests.post(supermarket_url, headers=headers)
    print(f"Request URL: {supermarket_url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")
