package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const minlon = -46.693419
const minlat = -23.568704
const maxlon = -46.641146
const maxlat = -23.546686

const url_all_data = "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
const url_sample = "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-sample.json"

type Property struct {
	Address       Address      `json:"address"`
	Bathrooms     int          `json:"bathrooms"`
	Bedrooms      int          `json:"bedrooms"`
	CreatedAt     time.Time    `json:"createdAt"`
	ID            string       `json:"id"`
	Images        []string     `json:"images"`
	ListingStatus string       `json:"listingStatus"`
	ListingType   string       `json:"listingType"`
	Owner         bool         `json:"owner"`
	ParkingSpaces int          `json:"parkingSpaces"`
	PricingInfos  PricingInfos `json:"pricingInfos"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	UsableAreas   int          `json:"usableAreas"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type GeoLocation struct {
	Location  Location `json:"location"`
	Precision string   `json:"precision"`
}
type Address struct {
	City         string      `json:"city"`
	GeoLocation  GeoLocation `json:"geoLocation"`
	Neighborhood string      `json:"neighborhood"`
}
type PricingInfos struct {
	BusinessType     string  `json:"businessType"`
	MonthlyCondoFee  string  `json:"monthlyCondoFee"`
	Price            string `json:"price"`
	RentalTotalPrice string `json:"rentalTotalPrice"`
	YearlyIptu       string  `json:"yearlyIptu"`
}

type Properties struct {
	all_properties      []Property
	zap_properties      []Property
	vivareal_properties []Property
}

func msg(log string)  {
	println(log)
}

func LoadProperties() *Properties {
	var properties_struct Properties
	msg("downloading data...")
	resp, err := http.Get(url_all_data)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	msg("reading data...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	msg("parsing data...")
	err = json.Unmarshal(body, &properties_struct.all_properties)
	if err != nil {
		log.Println(err.Error())
	}
	msg("done!")
	return &properties_struct
}

func (properties *Properties) GetZap() []Property {
	if properties.zap_properties != nil {
		return  properties.zap_properties
	}

	for _, property := range properties.all_properties {
		if property.Address.GeoLocation.Location.Lat != 0 && property.Address.GeoLocation.Location.Lon != 0 {
			if property.UsableAreas > 0 {
			rental_total_price, err := strconv.ParseFloat(property.PricingInfos.RentalTotalPrice, 64)
			if err != nil {
				continue
			}

			if property.PricingInfos.BusinessType == "RENTAL" && rental_total_price >= 3500.00 {
				properties.zap_properties = append(properties.zap_properties, property)
			} else if property.PricingInfos.BusinessType == "SALE" {
				sale_total_price, err := strconv.ParseFloat(property.PricingInfos.Price, 64)
				if err != nil {
					continue
				}

				if (property.Address.GeoLocation.Location.Lat >= minlat && property.Address.GeoLocation.Location.Lat <= maxlat) &&
					(property.Address.GeoLocation.Location.Lon >= minlon && property.Address.GeoLocation.Location.Lon <= maxlon) {
					if sale_total_price >= 540000.00 {
						if property.UsableAreas > 3500.00 {
							properties.zap_properties = append(properties.zap_properties, property)
						}
					}
				} else if sale_total_price >= 600000.00 {
					if property.UsableAreas > 3500.00 {
						properties.zap_properties = append(properties.zap_properties, property)
						}
					}
				}
			}
		}
	}

	return properties.zap_properties
}


func (properties *Properties) GetVivaReal() []Property {
	if properties.vivareal_properties != nil {
		return properties.vivareal_properties
	}

	for _, property := range properties.all_properties {
		if property.Address.GeoLocation.Location.Lat != 0 && property.Address.GeoLocation.Location.Lon != 0 {
			if property.PricingInfos.BusinessType == "RENTAL" {
				value_condo, err := strconv.ParseFloat(property.PricingInfos.MonthlyCondoFee, 64)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				rental_total_price, err := strconv.ParseFloat(property.PricingInfos.RentalTotalPrice, 64)
				if err != nil{
					continue
				}
				if (property.Address.GeoLocation.Location.Lat >= minlat && property.Address.GeoLocation.Location.Lat <= maxlat) &&
					(property.Address.GeoLocation.Location.Lon >= minlon && property.Address.GeoLocation.Location.Lon <= maxlon) {
					if rental_total_price <= 6000.00 {
						percent_value_30 := 1800.00
						if value_condo < percent_value_30 {
							properties.vivareal_properties = append(properties.vivareal_properties, property)
						}
					}
				} else if rental_total_price <= 4000.00 {
					PercentValue30 := 1200.00
					if value_condo < PercentValue30 {
						properties.vivareal_properties = append(properties.vivareal_properties, property)
					}
				}
			} else {
				sale_total_price, err := strconv.ParseFloat(property.PricingInfos.Price, 64)
				if err != nil {
					continue
				}
				if property.PricingInfos.BusinessType == "SALE" && sale_total_price <= 700000.00 {
					properties.vivareal_properties = append(properties.vivareal_properties, property)
				}
			}
		}
	}
	return properties.vivareal_properties
}
