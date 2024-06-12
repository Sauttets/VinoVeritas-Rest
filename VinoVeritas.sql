CREATE TABLE IF NOT EXISTS Wine(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,
    country VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    imageURL VARCHAR(255) NOT NULL,
    volume DECIMAL(10, 2) NOT NULL,
    volAlc DECIMAL(4, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS Flavour(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS FitsTo(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Wine_Flavour(
    wine_id INTEGER PRIMARY KEY,
    flavour_id_1 INTEGER,
    flavour_id_2 INTEGER,
    flavour_id_3 INTEGER,
    FOREIGN KEY (flavour_id_1) REFERENCES Flavour(id) ON DELETE SET NULL,
    FOREIGN KEY (flavour_id_2) REFERENCES Flavour(id) ON DELETE SET NULL,
    FOREIGN KEY (flavour_id_3) REFERENCES Flavour(id) ON DELETE SET NULL,
    FOREIGN KEY (wine_id) REFERENCES Wine(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Wine_FitsTo(
    wine_id INTEGER PRIMARY KEY,
    fitsTo_id_1 INTEGER,
    fitsTo_id_2 INTEGER,
    fitsTo_id_3 INTEGER,
    FOREIGN KEY (fitsTo_id_1) REFERENCES FitsTo(id) ON DELETE SET NULL,
    FOREIGN KEY (fitsTo_id_2) REFERENCES FitsTo(id) ON DELETE SET NULL,
    FOREIGN KEY (fitsTo_id_3) REFERENCES FitsTo(id) ON DELETE SET NULL,
    FOREIGN KEY (wine_id) REFERENCES Wine(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS FavoriteWines(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    wine_id INTEGER NOT NULL,
    FOREIGN KEY (wine_id) REFERENCES Wine(id),
    UNIQUE (user_id, wine_id)
);

CREATE TABLE IF NOT EXISTS Supermarkets(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    city VARCHAR(255) NOT NULL,
    house_number VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS WineSupermarkets(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    wine_id INTEGER NOT NULL,
    supermarket_id INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (wine_id) REFERENCES Wine(id),
    FOREIGN KEY (supermarket_id) REFERENCES Supermarkets(id),
    UNIQUE (wine_id, supermarket_id)
);

CREATE TABLE IF NOT EXISTS WineFactOTD(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    fact TEXT NOT NULL
);


-- @block

INSERT INTO WineFactOTD (fact) VALUES ('Wein ist eines der ältesten alkoholischen Getränke der Welt und wurde vor etwa 8.000 Jahren erstmals in Georgien hergestellt.');
INSERT INTO WineFactOTD (fact) VALUES ('Der teuerste Wein der Welt ist eine Flasche Romanée-Conti 1945, die 2018 für 558.000 US-Dollar verkauft wurde.');
INSERT INTO WineFactOTD (fact) VALUES ('Ein „Jahrgangswein“ stammt aus Trauben, die in einem einzigen Jahr geerntet wurden. „Nichtjahrgangsweine“ bestehen aus Trauben mehrerer Jahre.');
INSERT INTO WineFactOTD (fact) VALUES ('Es gibt über 10.000 verschiedene Rebsorten, aber nur etwa 1.300 werden kommerziell zur Weinproduktion verwendet.');
INSERT INTO WineFactOTD (fact) VALUES ('Der ideale Lagerort für Wein ist ein kühler, dunkler Keller mit einer konstanten Temperatur von etwa 12-15 Grad Celsius.');
INSERT INTO WineFactOTD (fact) VALUES ('Das Dekantieren von Wein ermöglicht es, dass Sauerstoff mit dem Wein reagiert, was die Aromen und Geschmacksnoten intensivieren kann.');
INSERT INTO WineFactOTD (fact) VALUES ('Hochwertige Weine werden zunehmend mit Schraubverschlüssen verschlossen, um Korkschmecker zu vermeiden.');
INSERT INTO WineFactOTD (fact) VALUES ('Der Begriff "Terroir" bezieht sich auf die Umweltfaktoren, die den Charakter des Weins beeinflussen, einschließlich Boden, Klima und Topographie.');
INSERT INTO WineFactOTD (fact) VALUES ('In Maßen genossen, kann Rotwein das Herz-Kreislauf-System unterstützen, da er Antioxidantien wie Resveratrol enthält.');
INSERT INTO WineFactOTD (fact) VALUES ('In vielen Kulturen hat Wein eine wichtige rituelle und religiöse Bedeutung, wie in der christlichen Eucharistie.');
INSERT INTO WineFactOTD (fact) VALUES ('Schwefeldioxid wird häufig als Konservierungsmittel in Wein verwendet, um Oxidation und bakteriellen Verderb zu verhindern.');
INSERT INTO WineFactOTD (fact) VALUES ('Echter Champagner kann nur aus der Champagne-Region in Frankreich stammen und muss bestimmten Produktionsmethoden folgen.');
INSERT INTO WineFactOTD (fact) VALUES ('Roséwein wird hergestellt, indem die Schalen roter Trauben nur kurz mit dem Saft in Kontakt bleiben, was ihm seine rosa Farbe verleiht.');
INSERT INTO WineFactOTD (fact) VALUES ('Süßweine wie Port und Sherry entstehen durch Hinzufügen von Brandy, um die Gärung zu stoppen und Restzucker zu erhalten.');
INSERT INTO WineFactOTD (fact) VALUES ('Wein wird oft in Eichenfässern gelagert, was ihm zusätzliche Aromen wie Vanille, Karamell und Rauch verleiht.');
INSERT INTO WineFactOTD (fact) VALUES ('Die meisten Weine haben einen Alkoholgehalt zwischen 12% und 14%, obwohl einige Sorten, wie Dessertweine, höher liegen können.');
INSERT INTO WineFactOTD (fact) VALUES ('Die Form des Weinglases kann die Wahrnehmung des Weingeschmacks beeinflussen, daher gibt es verschiedene Gläser für unterschiedliche Weinsorten.');
INSERT INTO WineFactOTD (fact) VALUES ('Bekannte Weinregionen sind Bordeaux und Burgund in Frankreich, Toskana in Italien und Napa Valley in den USA.');
INSERT INTO WineFactOTD (fact) VALUES ('Die Traubenlese erfolgt meist im Herbst und kann von Hand oder maschinell durchgeführt werden.');
INSERT INTO WineFactOTD (fact) VALUES ('Biowein wird aus Trauben hergestellt, die ohne synthetische Pestizide oder Düngemittel angebaut werden.');
INSERT INTO WineFactOTD (fact) VALUES ('Orange Wine ist ein Weißwein, der wie ein Rotwein hergestellt wird, indem die Trauben mit den Schalen fermentiert werden.');
INSERT INTO WineFactOTD (fact) VALUES ('Edelfäule (Botrytis cinerea) ist ein Pilz, der Trauben befällt und hochkonzentrierte, süße Weine wie Sauternes erzeugt.');
INSERT INTO WineFactOTD (fact) VALUES ('Schaumweine wie Prosecco, Cava und Crémant sind Alternativen zu Champagner und haben ihre eigenen Herstellungsmethoden.');
INSERT INTO WineFactOTD (fact) VALUES ('Professionelle Weinproben bewerten Weine nach Aussehen, Aroma, Geschmack und Abgang.');
INSERT INTO WineFactOTD (fact) VALUES ('Weißweine werden gekühlt (8-12°C) serviert, während Rotweine bei Raumtemperatur (16-18°C) am besten schmecken.');
INSERT INTO WineFactOTD (fact) VALUES ('Ein Dekanter kann verwendet werden, um Weine zu belüften und Sedimente zu entfernen, besonders bei älteren Rotweinen.');
INSERT INTO WineFactOTD (fact) VALUES ('Die Kombination von Wein und Speisen kann die Aromen beider verbessern, daher gibt es spezielle Pairing-Regeln.');
INSERT INTO WineFactOTD (fact) VALUES ('Weinverkostungsnotizen beschreiben oft Aromen und Geschmacksrichtungen wie Früchte, Gewürze und erdige Töne.');
INSERT INTO WineFactOTD (fact) VALUES ('Flaschen sollten liegend gelagert werden, damit der Korken feucht bleibt und nicht austrocknet.');
INSERT INTO WineFactOTD (fact) VALUES ('Die Römer trugen erheblich zur Verbreitung des Weinbaus in ganz Europa bei, indem sie Weinreben und Weinherstellungstechniken einführten.');
