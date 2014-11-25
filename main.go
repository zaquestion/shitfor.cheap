package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"

	"github.com/zaqthefreshman/Shitfor.Cheap/rakuten"
)

var apiString = fmt.Sprintf("http://api.popshops.com/v3/products.json?account=%s&catalog=%s", "ao7k0w59fbqag2stztxwdrod6", "db46yl7pq0tgy9iumgj88bfj7")

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[INFO] %s Request to /\n", r.Method)

		var data TemplateData

		if r.Method == "POST" {
			r.ParseForm()
			fmt.Println("tag:", r.Form["tag"])
			request := fmt.Sprintf("%s&keyword=%s&results_per_page=100&include_discounts=true", apiString, r.FormValue("tag"))
			fmt.Println(request)
			resp, err := http.Get(request)

			var apiResponse rakuten.FullRakutenResponse
			body, _ := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				fmt.Fprintf(w, "<html><head><title>Shit For Cheap!</title></head><body>Error:<br>%s</body></html>", err)
				return
			}
			RakutenToTemplate(&apiResponse, &data)
		} else if r.Method != "GET" {
			// Method Not Allowed
			w.WriteHeader(405)
			w.Write([]byte("<html><head><title>Shit For Cheap!</title></head><body>Error: Method Not Allowed</body></html>"))
			return
		}

		// Render template
		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		t.Execute(w, data)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

type TemplateData struct {
	Keyword string
	Offers  []Offer
}

type Offer struct {
	// product specific
	Brand              string
	Category           string
	PriceMax           float64
	PriceMaxMerchant   float64
	PriceMin           float64
	PriceMinMerchant   float64
	ProductDescription string
	ProductID          float64
	ProductImageUrl    string
	ProductName        string
	// best-offer specific
	Condition        string
	CurrencyIso      string
	Merchant         string
	PriceMerchant    float64
	PriceRetail      float64
	Sku              string
	URL              string
	OfferDescription string
	OfferID          float64
	OfferImageUrl    string
	OfferName        string
}

func RakutenToTemplate(rakutenProduct *rakuten.FullRakutenResponse, data *TemplateData) {
	//resources := &rakutenProduct.Resources
	for _, product := range rakutenProduct.Results.Products.Product {
		if product.Offers.Count > 0 {
			offer := product.Offers.Offer[0]
			data.Offers = append(data.Offers, Offer{
				ProductDescription: product.Description,
				ProductID:          product.ID,
				ProductImageUrl:    product.ImageUrlLarge,
				ProductName:        product.Name,
				Condition:          offer.Condition,
				PriceMerchant:      offer.PriceMerchant,
				URL:                offer.URL,
				PriceRetail:        offer.PriceRetail,
			})
		}
	}
}
