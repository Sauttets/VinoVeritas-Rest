import random
import requests
import urllib.parse

token = "token"
headers = {
    "Authorization": f"Bearer {token}"
}

set_wine_flavour_url = "http://localhost:8083/setWineFlavour"
set_wine_fit_url = "http://localhost:8083/setWineFit"

# Function to set wine flavours
def set_wine_flavours(wine_id, flavour_ids):
    query_params = {
        "wine_id": wine_id,
        "flavour1": flavour_ids[0],
        "flavour2": flavour_ids[1],
        "flavour3": flavour_ids[2]
    }
    url = set_wine_flavour_url + "?" + urllib.parse.urlencode(query_params)
    response = requests.post(url, headers=headers)
    print(f"Request URL: {url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")

# Function to set wine fits
def set_wine_fits(wine_id, fit_ids):
    query_params = {
        "wine_id": wine_id,
        "fitsTo1": fit_ids[0],
        "fitsTo2": fit_ids[1],
        "fitsTo3": fit_ids[2]
    }
    url = set_wine_fit_url + "?" + urllib.parse.urlencode(query_params)
    response = requests.post(url, headers=headers)
    print(f"Request URL: {url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")

# Adding flavours and fits to wines
for wine_id in range(1, 251):  # Assuming wine IDs are from 1 to 250
    flavour_ids = random.sample(range(1, 13), 3)  # Select 3 random flavours from 1 to 12
    fit_ids = random.sample(range(1, 13), 3)  # Select 3 random fits from 1 to 12
    set_wine_flavours(wine_id, flavour_ids)
    set_wine_fits(wine_id, fit_ids)
