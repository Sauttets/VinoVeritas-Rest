import random
import urllib.parse
import requests

# Example lists for generating random wine data
names = ["Cabernet Sauvignon", "Merlot", "Chardonnay", "Pinot Noir", "Sauvignon Blanc", "Syrah", "Zinfandel", "Malbec", "Riesling", "Prosecco"]
years = list(range(1990, 2023))
countries = ["France", "Italy", "Spain", "Germany", "USA", "Australia", "Argentina", "Chile", "South Africa", "New Zealand"]
types = ["red", "white", "rose", "sparkling", "dessert"]
descriptions = [
    "A rich, full-bodied wine with notes of blackberry and oak.",
    "A crisp and refreshing white wine with hints of citrus and green apple.",
    "A light and fruity rose perfect for summer sipping.",
    "A sparkling wine with delicate bubbles and a hint of brioche.",
    "A sweet dessert wine with flavors of honey and apricot."
]
base_image_url = "http://gargelkarg.com/images/"
volumes = [500, 750, 800, 850, 1000]
vol_alcs = [10.0, 11.5, 12.0, 12.5, 13.0, 13.5, 14.0, 14.5, 15.0]

token = "token"
headers = {
    "Authorization": f"Bearer {token}"
}

def generate_wine():
    name = random.choice(names)
    year = random.choice(years)
    country = random.choice(countries)
    wine_type = random.choice(types)
    description = random.choice(descriptions)
    image_url = base_image_url + f"{name.replace(' ', '_').lower()}.png"
    volume = random.choice(volumes)
    vol_alc = random.choice(vol_alcs)
    
    query_params = {
        "name": name,
        "year": year,
        "country": country,
        "type": wine_type,
        "description": description,
        "imageURL": image_url,
        "volume": volume,
        "volAlc": vol_alc
    }
    
    url = "http://localhost:8083/addWine?" + urllib.parse.urlencode(query_params)
    return url

# Generate and send 100 wine URLs
for _ in range(250):
    wine_url = generate_wine()
    response = requests.post(wine_url, headers=headers)
    #print(f"Request URL: {wine_url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")
