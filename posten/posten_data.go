package posten

import (
	"fmt"
)

type PostenResponse struct {
	Metadata struct {
		Page            int    `json:"side"`
		HitsPerPage     int    `json:"treffPerSide"`
		ShowsTo         int    `json:"viserTil"`
		SokeStreng      string `json:"sokeStreng"`
		AsciiCompatible bool   `json:"asciiKompatibel"`
		ShowsFrom       int    `json:"viserFra"`
		TotalHits       int    `json:"totaltAntallTreff"`
	} `json:"metadata"`
	Addresses []struct {
		AddressName                             string   `json:"addressenavn"`
		AddressText                             string   `json:"addressetekst"`
		AddressAdditionalName                   string   `json:"addressetilleggsnavn"`
		AddressCode                             int      `json:"addressekode"`
		Number                                  int      `json:"nummer"`
		Letter                                  string   `json:"bokstav"`
		MunicipalityNumber                      string   `json:"kommunenummer"`
		MunicipalityName                        string   `json:"kommunenavn"`
		YardNumber                              int      `json:"gardsnummer"`
		UseNumber                               int      `json:"bruksnummer"`
		AttachmentNumber                        int      `json:"festenummer"`
		SubNumber                               int      `json:"undernummer"`
		UtilityUnitNumber                       []string `json:"bruksenhetsnummer"`
		ObjectType                              string   `json:"objtype"`
		PostalPlace                             string   `json:"poststed"`
		PostalNumber                            string   `json:"postnummer"`
		AddressTextWithoutAddressAdditionalName string   `json:"adressetekstutenadressetilleggsnavn"`
		LocationVerified                        bool     `json:"stedfestingverifisert"`
		PointOfRepresentation                   struct {
			EPSG      string  `json:"epsg"`
			Latitude  float64 `json:"lat"`
			Longitude float64 `json:"lon"`
		} `json:"representasjonspunkt"`
		UpdateDate string `json:"oppdateringsdato"`
	} `json:"adresser"`
}

type DataResponse interface{}

type CompleteData struct {
	AddressName                             string   `json:"addressenavn"`
	AddressText                             string   `json:"addressetekst"`
	AddressAdditionalName                   string   `json:"addressetilleggsnavn"`
	AddressCode                             int      `json:"addressekode"`
	Number                                  int      `json:"nummer"`
	Letter                                  string   `json:"bokstav"`
	MunicipalityNumber                      string   `json:"kommunenummer"`
	MunicipalityName                        string   `json:"kommunenavn"`
	YardNumber                              int      `json:"gardsnummer"`
	UseNumber                               int      `json:"bruksnummer"`
	AttachmentNumber                        int      `json:"festenummer"`
	SubNumber                               int      `json:"undernummer"`
	UtilityUnitNumber                       []string `json:"bruksenhetsnummer"`
	ObjectType                              string   `json:"objtype"`
	PostalPlace                             string   `json:"poststed"`
	PostalNumber                            string   `json:"postnummer"`
	AddressTextWithoutAddressAdditionalName string   `json:"adressetekstutenadressetilleggsnavn"`
	LocationVerified                        bool     `json:"stedfestingverifisert"`
	PointOfRepresentation                   struct {
		EPSG      string  `json:"epsg"`
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"representasjonspunkt"`
	UpdateDate string `json:"oppdateringsdato"`
}

type ExtractedData struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	MunicipalityName string `json:"municipality_name"`
	PostalPlace      string `json:"postal_place"`
	PostalNumber     string `json:"postal_number"`
}

func (pr *PostenResponse) completeData() (DataResponse, error) {
	if len(pr.Addresses) == 0 {
		return CompleteData{}, fmt.Errorf("no addresses registered")
	}

	return pr.Addresses[0], nil
}

// Returns the following fields from the original Posten-API response:
//   - Coordinates (Latitude & Longitude)
//   - Municipality name
//   - Postalplace
//   - Postalnumber
func (pr *PostenResponse) extractedData() (DataResponse, error) {
	addresses := pr.Addresses
	if len(addresses) == 0 {
		return ExtractedData{}, fmt.Errorf("no addresses registered")
	}

	return ExtractedData{
		Coordinates: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude:  addresses[0].PointOfRepresentation.Latitude,
			Longitude: addresses[0].PointOfRepresentation.Longitude,
		},
		MunicipalityName: addresses[0].MunicipalityName,
		PostalPlace:      addresses[0].PostalPlace,
		PostalNumber:     addresses[0].PostalNumber,
	}, nil
}
