import requests
import urllib.parse

token = "token"
headers = {
    "Authorization": f"Bearer {token}"
}

flavours = ["Fruchtig", "Blumig", "Erdig", "Kräutrig", "Gewürzig", "Holzig", "Nussig", "Buttrig", "Süß", "Mineralisch"]
fits_to = ["Käse", "Fisch", "Rindfleisch", "Geflügel", "Schokolade", "Pasta", "Meeresfrüchte", "Lamm", "Desserts", "Vegetarisch"]

flavour_url = "http://localhost:8083/addFlavour"
fits_to_url = "http://localhost:8083/addFitsTo"

# Function to add flavour
def add_flavour(flavour):
    query_params = {"flavour": flavour}
    url = flavour_url + "?" + urllib.parse.urlencode(query_params)
    response = requests.post(url, headers=headers)
    print(f"Request URL: {url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")

# Function to add fits to
def add_fits_to(fit):
    query_params = {"fit_id": fit}
    url = fits_to_url + "?" + urllib.parse.urlencode(query_params)
    response = requests.post(url, headers=headers)
    print(f"Request URL: {url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")

# Add 12 flavours
for flavour in flavours:
    add_flavour(flavour)

# Add 12 fits to
for fit in fits_to:
    add_fits_to(fit)
