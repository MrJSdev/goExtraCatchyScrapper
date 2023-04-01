package scrapper

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/MrJSdev/goExtraCatchyScrapper/entity"
	"github.com/gocolly/colly"
)

func InitScrapping(link string) {
	c := colly.NewCollector()

	var projects []entity.Project

	var idCounter int = 1

	c.OnHTML("div.portfolio-item", func(h *colly.HTMLElement) {
		var project entity.Project

		project.Id = idCounter
		project.Name = h.ChildText("h3")
		project.Link = h.ChildAttr("a", "href")
		project.Image = h.ChildAttr("img", "src")

		projects = append(projects, project)

		idCounter++
	})

	// Save file when scraping is finished
	c.OnScraped(func(r *colly.Response) {
		SaveProjects(projects)
	})

	c.Visit(link)
}

// Save array of projects to a json file
func SaveProjects(projects []entity.Project) {
	json, err := json.MarshalIndent(projects, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("projects.json", json, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
