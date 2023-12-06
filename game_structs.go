package main

import "time"

// Game Data struct for the provided JSON data
type Game struct {
	Title                string  `json:"title"`
	ID                   string  `json:"id"`
	Namespace            string  `json:"namespace"`
	Description          string  `json:"description"`
	EffectiveDate        string  `json:"effectiveDate"`
	IsCodeRedemptionOnly bool    `json:"isCodeRedemptionOnly"`
	KeyImages            []Image `json:"keyImages"`
	CurrentPrice         int     `json:"currentPrice"`
	Seller               Seller  `json:"seller"`
	ProductSlug          string  `json:"productSlug"`
	URLSlug              string  `json:"urlSlug"`
	URL                  string  `json:"url"`
	Tags                 []Tag   `json:"tags"`
	Items                []struct {
		ID        string `json:"id"`
		Namespace string `json:"namespace"`
	} `json:"items"`
	CustomAttributes       []CustomAttribute `json:"customAttributes"`
	Categories             []Path            `json:"categories"`
	CatalogNs              CatalogNs         `json:"catalogNs"`
	OfferMappings          []Mapping         `json:"offerMappings"`
	DeveloperDisplayName   string            `json:"developerDisplayName"`
	PublisherDisplayName   string            `json:"publisherDisplayName"`
	Price                  Price             `json:"price"`
	PrePurchase            interface{}       `json:"prePurchase"`
	ReleaseDate            time.Time         `json:"releaseDate"`
	PCReleaseDate          time.Time         `json:"pcReleaseDate"`
	ViewableDate           time.Time         `json:"viewableDate"`
	ApproximateReleasePlan struct {
		Day             interface{} `json:"day"`
		Month           interface{} `json:"month"`
		Quarter         interface{} `json:"quarter"`
		Year            interface{} `json:"year"`
		ReleaseDateType string      `json:"releaseDateType"`
	} `json:"approximateReleasePlan"`
}

// Image struct for keyImages
type Image struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Seller struct for seller information
type Seller struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Tag struct for tags information
type Tag struct {
	ID string `json:"id"`
}

// CustomAttribute struct for custom attributes
type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Path struct for path information in categories
type Path struct {
	Path string `json:"path"`
}

// CatalogNs struct for catalog namespace information
type CatalogNs struct {
	Mappings []Mapping `json:"mappings"`
}

// Mapping struct for catalog mappings
type Mapping struct {
	PageSlug string `json:"pageSlug"`
	PageType string `json:"pageType"`
}

// Price struct for game pricing information
type Price struct {
	TotalPrice struct {
		DiscountPrice   int    `json:"discountPrice"`
		OriginalPrice   int    `json:"originalPrice"`
		VoucherDiscount int    `json:"voucherDiscount"`
		Discount        int    `json:"discount"`
		CurrencyCode    string `json:"currencyCode"`
		CurrencyInfo    struct {
			Decimals int `json:"decimals"`
		} `json:"currencyInfo"`
		FmtPrice struct {
			OriginalPrice     string `json:"originalPrice"`
			DiscountPrice     string `json:"discountPrice"`
			IntermediatePrice string `json:"intermediatePrice"`
		} `json:"fmtPrice"`
	} `json:"totalPrice"`
	LineOffers []struct {
		AppliedRules []interface{} `json:"appliedRules"`
	} `json:"lineOffers"`
}

type SearchStore struct {
	Elements []Game `json:"elements"`
}

type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

type Data struct {
	Catalog Catalog `json:"Catalog"`
}

type GamesJsonData struct {
	Data       Data        `json:"data"`
	Extensions interface{} `json:"extensions"`
}
