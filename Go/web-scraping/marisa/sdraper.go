// TESTE WEBSCRAPING
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

type item struct {
	Name	string `json: "name"`
	Price	string `json: "price"`
	ImgUrl	string `json: "imgurl"`
}

func main(){
	c := colly.NewCollector(
		colly.AllowedDomains("marisa.com.br"),
	)

	var items []item

	c.OnHTML("div.product-grid-component ", func (h *colly.HTMLElement)  {
		item := item{
			Name: h.ChildText("span.title"),
			Price: h.ChildText("div.ProductBox_price__rs1yB"),
			ImgUrl: h.ChildAttr("img", "src"),

		}
		items = append(items, item)
	})

	c.OnHTML("div.Pagination_next__BZ82d", func (h *colly.HTMLElement)  {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func (r*colly.Request)  {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://www.lojasrenner.com.br/lista?s_icid=202103_HOME_MENU_BASICOS_BLUSAS_FEM&filter=true")

	content, err := json.Marshal(items)
	
	if err != nil{
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)



}