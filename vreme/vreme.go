// Package vreme omogoča upravljanje vremenskih postaj, beleženje meritev
// temperature in izpis statističnih poročil.
package vreme

import (
	f "fmt"
)

// Postaja opisuje vremensko postajo z imenom, krajem in seznamom meritev.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar doda novo postajo z dano oznako, imenom, krajem in začetnimi meritvami.
func DodajPostajo(sifra string, ime string, kraj string, meritve []float64, postaje map[string]Postaja) {
	postaje[sifra] = Postaja{ime, kraj, meritve}
}

// DodajMeritev postaji z dano oznako doda meritev, če leži med minTemp in maxTemp.
func DodajMeritev(minTemp, maxTemp float64, postaje map[string]Postaja, oznaka string, temperatura float64) {
	if temperatura >= minTemp && temperatura <= maxTemp {
		new, valid := postaje[oznaka]
		if !valid {
			f.Printf("Postaja ne obstaja.\n")
			return
		}
		new.meritve = append(new.meritve, temperatura)
		postaje[oznaka] = new
	} else {
		f.Printf("Neveljavna meritev.\n")
	}
}

// statistika je pomožna funkcija, ki izračuna min, max in povprečje meritev neke postaje, če je število meritev dovolj veliko
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	tabela, valid := postaje[oznaka]
	if len(tabela.meritve) < stMeritev || !valid {
		return 0, 0, 0, false
	}

	min := tabela.meritve[0]
	max := tabela.meritve[0]
	sum := 0.0

	for _, x := range tabela.meritve {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		sum += x
	}
	return min, max, (sum / (float64(len(tabela.meritve)))), true
}

// IzpisMeritev izpiše oznako, ime, kraj in meritve vseh postaj.
func IzpisMeritev(postaje map[string]Postaja) {
	f.Println("MERITVE:")
	for oznaka, x := range postaje {
		f.Printf("%s - %s (%s): %v\n", oznaka, x.ime, x.kraj, x.meritve)
	}
}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno
// temperaturo, če je meritev vsaj stMeritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, x := range postaje {
		min, max, avg, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			f.Printf("%s: premalo meritev\n", x.ime)
			continue
		}
		rez := ""
		switch {
		case avg >= 25:
			rez = "Vroče"
		case avg >= 15:
			rez = "Toplo"
		case avg >= 5:
			rez = "Hladno"
		default:
			rez = "Mrzlo"
		}
		f.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", x.ime, min, max, avg, rez)
	}
}
