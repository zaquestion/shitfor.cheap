package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Form)

		if r.Method == "GET" {
			http.ServeFile(w, r, "./static/index.html")
		} else {
			r.ParseForm()
			fmt.Println("tag:", r.Form["tag"])
			http.ServeFile(w, r, "./static/index.html")
			resp, err := http.Get("http://api.popshops.com/v3/products.json?account=ao7k0w59fbqag2stztxwdrod6&catalog=db46yl7pq0tgy9iumgj88bfj7&keyword=" + r.FormValue("tag"))

			var rakuten RakutenProduct
			body, err := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &rakuten)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(rakuten.Status)

		}

	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

type RakutenProduct struct {
	Filters struct {
		Filter []interface{} `json:"filter"`
	} `json:"filters"`
	Message    string `json:"message"`
	Parameters []struct {
		Kind  string `json:"kind"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"parameters"`
	Resources struct {
		Brands struct {
			Brand []interface{} `json:"brand"`
			Count float64       `json:"count"`
		} `json:"brands"`
		Categories struct {
			Context struct {
				Category []struct {
					ID    float64 `json:"id"`
					Name  string  `json:"name"`
					Order float64 `json:"order"`
				} `json:"category"`
			} `json:"context"`
			Matches struct {
				Category []struct {
					Count float64 `json:"count"`
					ID    float64 `json:"id"`
					Name  string  `json:"name"`
				} `json:"category"`
			} `json:"matches"`
		} `json:"categories"`
		DealTypes struct {
			Count    float64 `json:"count"`
			DealType []struct {
				Count float64 `json:"count"`
				ID    float64 `json:"id"`
				Name  string  `json:"name"`
			} `json:"deal_type"`
		} `json:"deal_types"`
		Merchants struct {
			Count    float64       `json:"count"`
			Merchant []interface{} `json:"merchant"`
		} `json:"merchants"`
	} `json:"resources"`
	Results struct {
		Deals    struct{} `json:"deals"`
		Products struct {
			Count   float64 `json:"count"`
			Product []struct {
				Brand         float64 `json:"brand"`
				Category      float64 `json:"category"`
				Description   string  `json:"description"`
				ID            float64 `json:"id"`
				ImageUrlLarge string  `json:"image_url_large"`
				Name          string  `json:"name"`
				OfferCount    float64 `json:"offer_count"`
				Offers        struct {
					Count float64 `json:"count"`
					Offer []struct {
						Condition     string  `json:"condition"`
						CurrencyIso   string  `json:"currency_iso"`
						Description   string  `json:"description"`
						ID            float64 `json:"id"`
						ImageUrlLarge string  `json:"image_url_large"`
						Merchant      float64 `json:"merchant"`
						Name          string  `json:"name"`
						PriceMerchant float64 `json:"price_merchant"`
						PriceRetail   float64 `json:"price_retail"`
						Sku           string  `json:"sku"`
						URL           string  `json:"url"`
					} `json:"offer"`
				} `json:"offers"`
				PriceMax         float64 `json:"price_max"`
				PriceMaxMerchant float64 `json:"price_max_merchant"`
				PriceMin         float64 `json:"price_min"`
				PriceMinMerchant float64 `json:"price_min_merchant"`
			} `json:"product"`
		} `json:"products"`
	} `json:"results"`
	Session string  `json:"session"`
	Status  float64 `json:"status"`
}
