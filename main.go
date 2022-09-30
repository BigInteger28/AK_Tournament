package main

import (
	"fmt"
	"sort"
	"strconv"
	"os"
)

type Speler struct {
	id      int
	naam    string
	level   int
	positie int
	score   float64
}

func sorteerPerLevel(spelers []Speler) {
	sort.Slice(spelers, func(i, j int) bool {
		return (spelers[i].level) > (spelers[j].level)
	})
}

func sorteerPerPlaats(spelers []Speler) {
	sort.Slice(spelers, func(i, j int) bool {
		return (spelers[i].positie) < (spelers[j].positie)
	})
}

func sorteerGroep(spelers []Speler, groep int, sg int) {
	start := (groep - 1) * sg
	laatste := start + sg
	g := spelers[start:laatste]
	sort.Slice(g, func(i, j int) bool {
		return (g[i].score) > (g[j].score)
	})
	for i := start; i < laatste; i++ {
		spelers[i].positie = i + 1
	}
}

func stelGroepIn(spelers []Speler, sg int, groep int) {
	for i := 0; i < sg; i++ {
		plaats := i + ((groep - 1) * sg)
		fmt.Print(spelers[plaats].naam, "\tScore: ")
		fmt.Scanln(&spelers[plaats].score)
	}
	sorteerGroep(spelers, groep, sg)
}

func maakVolgendeRonde(spelers []Speler, sg int, aantalGroepen int) {
	var bot int = sg / 2
	var top int = (sg + 1) / 3
	for g := 1; g <= aantalGroepen; g++ {
		if g != 1 {
			for i := 0; i < top; i++ {
				spelers[i+((g-1)*sg)].positie -= top
			}
		}
		if g != aantalGroepen {
			for i := 0; i < bot; i++ {
				if sg%2 == 0 {
					spelers[i+((g-1)*sg)+bot].positie += bot
				} else {
					spelers[i+((g-1)*sg)+bot+1].positie += bot
				}
			}
		}
	}
	sorteerPerPlaats(spelers)
	for i := range spelers {
		spelers[i].score = 0
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exportToHtml(spelers []Speler, ronde int, aantalRonden int, aantalGroepen int, gamesPerGroep int, sg int) {
	var waarden = [4]int {ronde, aantalRonden, aantalGroepen, gamesPerGroep}
	swaarden := make([]string, 4)
	for i := 0; i < 4; i++ {
		swaarden[i] = strconv.Itoa(waarden[i])
	}
	var filename string
	var htmlcode string = "<!DOCTYPE html><html><head><style>td { text-align:center; width:150px; }</style></head><body>"
	if ronde != aantalRonden {
		filename = "Round_" + swaarden[0] + ".html"
		htmlcode += "<h1>Ronde "+swaarden[0]+"&emsp;/&emsp;" + swaarden[1]  + "</h1>"
	} else {
		filename = "EINDSTAND.html"
		htmlcode += "<h1>EINDSTAND</h1>"
	}
		htmlcode += "<br><br>Games per groep: " + swaarden[3]
	for g := 1; g <= aantalGroepen; g++ {
		htmlcode += "<br><br><h2>GROEP " + strconv.Itoa(g)
		htmlcode += "<table><tr><th>Plaats</th><th>Naam</th><th>Level</th><th>Score</th><th>Totaal Score</th></tr>"
		for s := 0; s < sg; s++ {
			plaats := s + ((g - 1) * sg)
			levelS := strconv.Itoa(spelers[plaats].level)
			scoreS := fmt.Sprintf("%.1f", spelers[plaats].score)
			htmlcode += "<tr><td>" + strconv.Itoa(spelers[plaats].positie) + "</td><td>" + spelers[plaats].naam + "</td><td>" +  levelS + "</td><td>" + scoreS + "</td><td>" + strconv.Itoa((sg-1)*gamesPerGroep) + "</td></tr>"
		}
		htmlcode += "</table>"
	}
	htmlcode += "</body></html>"	
	file, err := os.Create(filename)
	check(err)
	file.WriteString(string(htmlcode))
	file.Close()
	file.Sync()
}

func main() {
	var s int
	var sg int
	fmt.Print("Aantal spelers: ")
	fmt.Scanln(&s)
	fmt.Print("Spelers per groep: ")
	fmt.Scanln(&sg)
	spelers := make([]Speler, s)
	for i := 0; i < s; i++ {
		spelers[i].id = i
		fmt.Print("Naam: ")
		fmt.Scanln(&spelers[i].naam)
		fmt.Print(spelers[i].naam, " is level: ")
		fmt.Scanln(&spelers[i].level)
	}
	sorteerPerLevel(spelers)
	for i := range spelers {
		spelers[i].positie = i + 1
	}
	aantalGroepen := s / sg
	aantalRonden := aantalGroepen + 1
	gamesPerGroep := sg - 1

	for ronde := 1; ronde <= aantalRonden; ronde++ {
		fmt.Println("\nRONDE ", ronde, " / ", aantalRonden, "\t\tGames per groep: ", gamesPerGroep)
		sorteerPerPlaats(spelers)
		for groep := 1; groep <= aantalGroepen; groep++ {
			fmt.Println("\n\tGROEP ", groep)
			for s := 0; s < sg; s++ {
				plaats := s + ((groep - 1) * sg)
				fmt.Println("\t\t", spelers[plaats].positie, " ", spelers[plaats].naam, "\t\tLevel:", spelers[plaats].level, "\tScore: ", spelers[plaats].score, "/", (sg-1)*gamesPerGroep)
			}
		}
		keuze := -1
		kiesgroep := -1
		//zorg dat het programma niet crasht bij invullen verkeerd nummer of string
		for keuze != 1 {
			if ronde != aantalRonden {
				fmt.Println("1. Maak volgende ronde aan")
			} else {
				fmt.Println("1. TOON EINDSTAND")
			}
			fmt.Println("2. Vul scores in van een groep")
			fmt.Println("3. EXPORT CURRENT TO HTML")
			fmt.Print("Keuze: ")
			fmt.Scanln(&keuze)
			if keuze == 2 {
				fmt.Print("Scores invullen groep: ")
				fmt.Scanln(&kiesgroep)
				stelGroepIn(spelers, sg, kiesgroep)
				for groep := 1; groep <= aantalGroepen; groep++ {
					fmt.Println("\n\tGROEP ", groep)
					for s := 0; s < sg; s++ {
						plaats := s + ((groep - 1) * sg)
						fmt.Println("\t\t", spelers[plaats].positie, " ", spelers[plaats].naam, "\t\tLevel:", spelers[plaats].level, "\tScore: ", spelers[plaats].score, "/", (sg-1)*gamesPerGroep)
					}
				}
			}
			if keuze == 3 {
				exportToHtml(spelers, ronde, aantalRonden, aantalGroepen, gamesPerGroep, sg)
			}
		}
		if ronde != aantalRonden {
			maakVolgendeRonde(spelers, sg, aantalGroepen)
		} else {
			fmt.Println()
			for o := range spelers {
				fmt.Println("\t\t", spelers[o].positie, " ", spelers[o].naam, "\t\tLevel:", spelers[o].level)
			}
		}
	}
}
