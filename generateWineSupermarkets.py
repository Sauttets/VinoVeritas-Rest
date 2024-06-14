import random
import urllib.parse
import requests

token = "token"
headers = {
    "Authorization": f"Bearer {token}"
}

wine_check_url = "http://localhost:8083/getWine?id="
set_wine_supermarket_url = "http://localhost:8083/setWineSupermarket"

# Weighted random price generator
def generate_price():
    price_bins = [
        (2.99, 10.00, 50),   # 50% probability for prices between 2.99 and 10.00
        (10.01, 20.00, 30),  # 30% probability for prices between 10.01 and 20.00
        (20.01, 35.00, 15),  # 15% probability for prices between 20.01 and 35.00
        (35.01, 49.99, 5)    # 5% probability for prices between 35.01 and 49.99
    ]
    
    bins, weights = zip(*[(bin_range, weight) for bin_range, _, weight in price_bins])
    chosen_bin = random.choices(price_bins, weights)[0]
    return round(random.uniform(chosen_bin[0], chosen_bin[1]), 2)

# Function to check if wine exists
def wine_exists(wine_id):
    response = requests.get(wine_check_url + str(wine_id), headers=headers)
    return response.status_code == 200

# Function to add wine to supermarket
def add_wine_to_supermarket(wine_id, supermarket_id, price):
    query_params = {
        "wine_id": wine_id,
        "supermarket_id": supermarket_id,
        "price": price
    }
    url = set_wine_supermarket_url + "?" + urllib.parse.urlencode(query_params)
    response = requests.post(url, headers=headers)
    print(f"Request URL: {url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")

# Adding wines to up to 3 supermarkets
for wine_id in range(1, 300):  # Assuming wine IDs are from 1 to 100
    if wine_exists(wine_id):
        supermarket_ids = random.sample(range(1, 12), k=random.randint(1, 3))  # Select 1 to 3 random supermarkets
        for supermarket_id in supermarket_ids:
            price = generate_price()
            add_wine_to_supermarket(wine_id, supermarket_id, price)
