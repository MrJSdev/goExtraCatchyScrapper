package main

import (
	"github.com/MrJSdev/goExtraCatchyScrapper/helpers"
	"github.com/MrJSdev/goExtraCatchyScrapper/scrapper"
)

func init() {
	helpers.LoadEnvVariableFromFile()
}

func main() {
	var portfolioLink string = helpers.GetEnv("PORTFOLIO_LINK")

	scrapper.InitScrapping(portfolioLink)

}
