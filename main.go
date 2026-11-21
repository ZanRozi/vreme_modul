package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ZanRozi/vreme_modul/vreme"
	"github.com/urfave/cli/v3"
)

func izvedi(minTemp, maxTemp float64, stMeritev int) error {
	var postaje = make(map[string]vreme.Postaja)
	vreme.DodajPostajo("LJ01", "ZOO", "Ljubljana", []float64{}, postaje)
	vreme.DodajPostajo("LJ02", "BTC", "Ljubljana", []float64{}, postaje)
	vreme.DodajPostajo("KR01", "Center", "Kranj", []float64{}, postaje)
	vreme.DodajPostajo("MB01", "Pohorje", "Maribor", []float64{}, postaje)
	vreme.DodajPostajo("KP01", "Pristanisce", "Koper", []float64{}, postaje)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", 26.5)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", 27.0)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", 25.8)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", 28.1)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ02", 18.2)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ02", 19.5)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ02", 17.9)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ02", 20.3)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KR01", 8.4)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KR01", 9.1)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KR01", 7.6)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KR01", 10.2)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "MB01", 1.5)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "MB01", -2.0)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "MB01", 0.8)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "MB01", 3.2)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KP01", 24.5)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "KP01", 27.0)

	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", 99.9)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "LJ01", -80.0)
	vreme.DodajMeritev(minTemp, maxTemp, postaje, "XX99", 20.0)

	fmt.Println()
	vreme.IzpisMeritev(postaje)
	fmt.Println()

	vreme.IzpisPorocila(postaje, stMeritev)

	return nil
}

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Vremenske postaje",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Usage: "najnižja dovoljena temperatura v °C", Value: -50},
			&cli.FloatFlag{Name: "maxTemp", Usage: "najvišja dovoljena temperatura v °C", Value: 50},
			&cli.IntFlag{Name: "stMeritev", Usage: "najmanjše število meritev za statistiko", Value: 3},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minTemp := c.Float("minTemp")
			maxTemp := c.Float("maxTemp")
			stMeritev := int(c.Int("stMeritev"))

			return izvedi(minTemp, maxTemp, stMeritev)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
