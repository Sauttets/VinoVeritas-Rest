package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var WineFacts = []string{
	"Wein ist eines der ältesten alkoholischen Getränke der Welt und wurde vor etwa 8.000 Jahren erstmals in Georgien hergestellt.",
	"Der teuerste Wein der Welt ist eine Flasche Romanée-Conti 1945, die 2018 für 558.000 US-Dollar verkauft wurde.",
	"Ein „Jahrgangswein“ stammt aus Trauben, die in einem einzigen Jahr geerntet wurden. „Nichtjahrgangsweine“ bestehen aus Trauben mehrerer Jahre.",
	"Es gibt über 10.000 verschiedene Rebsorten, aber nur etwa 1.300 werden kommerziell zur Weinproduktion verwendet.",
	"Der ideale Lagerort für Wein ist ein kühler, dunkler Keller mit einer konstanten Temperatur von etwa 12-15 Grad Celsius.",
	"Das Dekantieren von Wein ermöglicht es, dass Sauerstoff mit dem Wein reagiert, was die Aromen und Geschmacksnoten intensivieren kann.",
	"Hochwertige Weine werden zunehmend mit Schraubverschlüssen verschlossen, um Korkschmecker zu vermeiden.",
	"Der Begriff \"Terroir\" bezieht sich auf die Umweltfaktoren, die den Charakter des Weins beeinflussen, einschließlich Boden, Klima und Topographie.",
	"In Maßen genossen, kann Rotwein das Herz-Kreislauf-System unterstützen, da er Antioxidantien wie Resveratrol enthält.",
	"In vielen Kulturen hat Wein eine wichtige rituelle und religiöse Bedeutung, wie in der christlichen Eucharistie.",
	"Schwefeldioxid wird häufig als Konservierungsmittel in Wein verwendet, um Oxidation und bakteriellen Verderb zu verhindern.",
	"Echter Champagner kann nur aus der Champagne-Region in Frankreich stammen und muss bestimmten Produktionsmethoden folgen.",
	"Roséwein wird hergestellt, indem die Schalen roter Trauben nur kurz mit dem Saft in Kontakt bleiben, was ihm seine rosa Farbe verleiht.",
	"Süßweine wie Port und Sherry entstehen durch Hinzufügen von Brandy, um die Gärung zu stoppen und Restzucker zu erhalten.",
	"Wein wird oft in Eichenfässern gelagert, was ihm zusätzliche Aromen wie Vanille, Karamell und Rauch verleiht.",
	"Die meisten Weine haben einen Alkoholgehalt zwischen 12% und 14%, obwohl einige Sorten, wie Dessertweine, höher liegen können.",
	"Die Form des Weinglases kann die Wahrnehmung des Weingeschmacks beeinflussen, daher gibt es verschiedene Gläser für unterschiedliche Weinsorten.",
	"Bekannte Weinregionen sind Bordeaux und Burgund in Frankreich, Toskana in Italien und Napa Valley in den USA.",
	"Die Traubenlese erfolgt meist im Herbst und kann von Hand oder maschinell durchgeführt werden.",
	"Biowein wird aus Trauben hergestellt, die ohne synthetische Pestizide oder Düngemittel angebaut werden.",
	"Orange Wine ist ein Weißwein, der wie ein Rotwein hergestellt wird, indem die Trauben mit den Schalen fermentiert werden.",
	"Edelfäule (Botrytis cinerea) ist ein Pilz, der Trauben befällt und hochkonzentrierte, süße Weine wie Sauternes erzeugt.",
	"Schaumweine wie Prosecco, Cava und Crémant sind Alternativen zu Champagner und haben ihre eigenen Herstellungsmethoden.",
	"Professionelle Weinproben bewerten Weine nach Aussehen, Aroma, Geschmack und Abgang.",
	"Weißweine werden gekühlt (8-12°C) serviert, während Rotweine bei Raumtemperatur (16-18°C) am besten schmecken.",
	"Ein Dekanter kann verwendet werden, um Weine zu belüften und Sedimente zu entfernen, besonders bei älteren Rotweinen.",
	"Die Kombination von Wein und Speisen kann die Aromen beider verbessern, daher gibt es spezielle Pairing-Regeln.",
	"Weinverkostungsnotizen beschreiben oft Aromen und Geschmacksrichtungen wie Früchte, Gewürze und erdige Töne.",
	"Flaschen sollten liegend gelagert werden, damit der Korken feucht bleibt und nicht austrocknet.",
	"Die Römer trugen erheblich zur Verbreitung des Weinbaus in ganz Europa bei, indem sie Weinreben und Weinherstellungstechniken einführten.",
}

func GetWineFactOTD(c *gin.Context) {
	currentDay := time.Now().Day()
	index := (currentDay - 1) % len(WineFacts)
	c.JSON(http.StatusOK, gin.H{
		"fact": WineFacts[index],
	})
}
