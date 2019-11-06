package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// isOldest Take the scouts name and print whether they are the oldest
func isOldest(scout string) {

	fmt.Printf("%s is the youngest\n", scout)

	fmt.Printf("%s is not the oldest\n", scout)

	fmt.Printf("%s is the oldest\n", scout)
}

// birthday Take the scouts name and print their birthday
func birthday(scout string) {

}

// age prints the age processed based on the command
func age(scout, action string) {

}

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "oldest",
			Aliases: []string{""},
			Usage:   "Pass the name of a scout and print yes if they are the oldest",
			Action: func(c *cli.Context) error {
				scout := c.Args().First()
				if scout == "" {
					return fmt.Errorf("invalid scout name")
				}

				isOldest(scout)

				return nil
			},
		},
		{
			Name:    "birthday",
			Aliases: []string{""},
			Usage:   "Pass the name of a scout and print yes if they are the oldest",
			Action: func(c *cli.Context) error {
				scout := c.Args().First()
				if scout == "" {
					return fmt.Errorf("invalid scout name")
				}

				birthday(scout)

				return nil
			},
		},
		{
			Name: "age",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "average, avg",
				},
				cli.BoolFlag{
					Name: "sum",
				},
				cli.BoolFlag{
					Name: "product",
				},
			},
			Aliases: []string{""},
			Usage:   "Pass the name of a scout and print yes if they are the oldest",
			Action: func(c *cli.Context) error {

				cmd := ""
				if c.Bool("average") {
					cmd = "average"
				} else if c.Bool("sum") {
					cmd = "sum"
				} else if c.Bool("product") {
					cmd = "product"
				}

				scout := c.Args().First()
				if scout == "" {
					return fmt.Errorf("Invalid scout name")
				}

				age(scout, cmd)

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
