import random
import urllib.parse
import requests

# Example lists for generating random wine data
names = ["Eldoria Elixir", "Dragon's Breath", "Elven Nectar", "Dwarven Stout", "Mystic Mead", "Phoenix Fire", "Mermaid's Tears", "Unicorn Bliss", "Goblin Brew", "Fairy Fizz",
         "Wizard's Wine", "Sorcerer's Sigh", "Enchanted Essence", "Orcish Ale", "Gnome's Glee", "Vampire's Vino", "Gryphon Grape", "Witch's Whim", "Basilisk Berry", "Sphinx Spirit",
         "Troll's Tipple", "Chimera's Choice", "Harpy's Hooch", "Centaur Cellar", "Merlin's Magic", "Titan's Toast", "Dryad's Delight", "Frost Giant's Froth", "Kraken's Keg", "Hydra's Hooch",
         "Pegasus Potion", "Leprechaun's Libation", "Minotaur's Mead", "Nymph's Nectar", "Pixie Potion", "Satyr's Sip", "Selkie's Swirl", "Banshee's Beverage", "Cyclops' Cup", "Djinn's Drink",
         "Dryad Draught", "Elemental Elixir", "Gargoyle Grog", "Ghoul's Gulp", "Griffin's Gulp", "Harlequin's Harmony", "Incubus' Infusion", "Jester's Juice", "Kelpie's Kiss", "Lamia's Libation",
         "Manticore's Mead", "Naga's Nectar", "Ogre's Oath", "Phantom's Philter", "Pixie's Pilsner", "Quicksilver Quencher", "Raven's Rum", "Siren's Sangria", "Spectre's Spirit", "Sprite's Seltzer",
         "Sylph's Swig", "Thorn's Thirst", "Trickster's Tonic", "Undine's Unction", "Valkyrie's Vintage", "Will-o'-the-Wisp's Wine", "Wraith's Wine", "Yeti's Yawn", "Zephyr's Zinfandel", "Aether Ale",
         "Basilisk's Brew", "Centaure's Chardonnay", "Drake's Drink", "Efreet's Elixir", "Faun's Fizz", "Genie's Gin", "Hobgoblin's Hooch", "Icarus' Ichor", "Jinn's Juice", "Kelpie's Cooler",
         "Leviathan's Lager", "Mogwai's Mead", "Nereid's Nectar", "Oni's Ouzo", "Phoenix's Punch", "Qilin's Quench", "Roc's Rum", "Satyr's Sipper", "Troll's Toddy", "Undead's Updraft",
         "Vampyre's Vintage", "Wyvern's Wine", "Xorn's Xyloid", "Yokai's Yuzu", "Zephyr's Zinfandel", "Arcane Ale", "Berserker's Brew", "Changeling's Cider", "Djinn's Draught", "Elf's Elixir",
         "Frostbite Fizz", "Goblin's Grog", "Huldra's Honey", "Imp's Infusion", "Jackal's Juice", "Krampus' Keg", "Lich's Liquor", "Medusa's Mead", "Nymph's Nectar", "Ogre's Ouzo",
         "Phantom's Philter", "Quetzal's Quencher", "Raven's Red", "Sylvan Spirit", "Tengu's Tipple", "Ursa's Unction", "Void Vine", "Wendigo's Wine", "Xenon's Xyloid", "Yew's Yearn"]
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

print(len(names))

def generate_wine():
    name = random.choice(names)
    names.remove(name)  # Remove the name from the list to ensure uniqueness
    year = random.choice(years)
    country = random.choice(countries)
    wine_type = random.choice(types)
    description = random.choice(descriptions)
    image_url = base_image_url + f"{name.replace(' ', '_').lower()}.png"
    volume = random.choice(volumes)
    vol_alc = random.choice(vol_alcs)
    
    dryness = round(random.uniform(0, 0.87), 2)
    acidity = round(random.uniform(0, 0.87), 2)
    tanninLevel = round(random.uniform(0, 0.87), 2)
    
    query_params = {
        "name": name,
        "year": year,
        "country": country,
        "type": wine_type,
        "description": description,
        "imageURL": image_url,
        "volume": volume,
        "volAlc": vol_alc,
        "dryness": dryness,
        "acidity": acidity,
        "tanninLevel": tanninLevel
    }
    
    url = "http://localhost:8083/addWine?" + urllib.parse.urlencode(query_params)
    return url

# Generate and send 200 wine URLs
for _ in range(len(names)):
    wine_url = generate_wine()
    response = requests.post(wine_url, headers=headers)
    #print(f"Request URL: {wine_url}")
    print(f"Response Status Code: {response.status_code}")
    if response.status_code != 200:
        print(f"Error: {response.text}")
