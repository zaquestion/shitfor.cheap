package rakuten

type MerchantResource struct {
	ID      float64 `json:"id"`
	Name    string  `json:"name"`
	Count   float64 `json:"count"`
	LogoUrl string  `json:"url"`
	Url     string  `json:"url"`
}

type BrandResource struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

type CategoryResource struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Order float64 `json:"order"`
}

type DealTypeResource struct {
	Count float64 `json:"count"`
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
}

type Parameter struct {
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Resources struct {
	Brands struct {
		Count float64         `json:"count"`
		Brand []BrandResource `json:"brand"`
	} `json:"brands"`
	Categories struct {
		Context struct {
			Category []CategoryResource `json:"category"`
		} `json:"context"`
		Matches struct {
			Category []CategoryResource `json:"category"`
		} `json:"matches"`
	} `json:"categories"`
	DealTypes struct {
		Count    float64            `json:"count"`
		DealType []DealTypeResource `json:"deal_type"`
	} `json:"deal_types"`
	Merchants struct {
		Count    float64            `json:"count"`
		Merchant []MerchantResource `json:"merchant"`
	} `json:"merchants"`
}

type Filter interface{}

type Product struct {
	Brand         float64 `json:"brand"`
	Category      float64 `json:"category"`
	Description   string  `json:"description"`
	ID            float64 `json:"id"`
	ImageUrlLarge string  `json:"image_url_large"`
	Name          string  `json:"name"`
	OfferCount    float64 `json:"offer_count"`
	Offers        struct {
		Count float64 `json:"count"`
		Offer []Offer `json:"offer"`
	} `json:"offers"`
	PriceMax         float64 `json:"price_max"`
	PriceMaxMerchant float64 `json:"price_max_merchant"`
	PriceMin         float64 `json:"price_min"`
	PriceMinMerchant float64 `json:"price_min_merchant"`
}

type Offer struct {
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
}

type FullRakutenResponse struct {
	Filters struct {
		Count  float64  `json:"count"`
		Filter []Filter `json:"filter"`
	} `json:"filters"`
	Message    string      `json:"message"`
	Parameters []Parameter `json:"parameters"`
	Resources  Resources   `json:"resources"`
	Results    struct {
		Deals struct {
			Count float64       `json:"count"`
			Deal  []interface{} `json:"deal"`
		} `json:"deals"`
		Products struct {
			Count   float64   `json:"count"`
			Product []Product `json:"product"`
		} `json:"products"`
	} `json:"results"`
	Session string  `json:"session"`
	Status  float64 `json:"status"`
}

var (
	a FullRakutenResponse
	b = a.Filters
)
