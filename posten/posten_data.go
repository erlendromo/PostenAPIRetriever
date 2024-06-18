package posten

import (
	"fmt"
	"time"
)

type PostenResponse struct {
	Metadata Metadata  `json:"metadata"`
	Adresses []Address `json:"adresser"`
}

func (pr *PostenResponse) GetMetadata() Metadata {
	return pr.Metadata
}

func (pr *PostenResponse) GetAddresses() []Address {
	return pr.Adresses
}

func (pr *PostenResponse) ExtractValuableData() (*ValuableData, error) {
	addresses := pr.GetAddresses()
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no addresses registered")
	}
	data := addresses[0]

	return &ValuableData{
		Coordinates: Coordinates{
			Latitude:  data.PointOfRepresentation.Latitude,
			Longitude: data.PointOfRepresentation.Longitude,
		},
		MunicipalityName: data.MunicipalityName,
		PostalPlace:      data.PostalPlace,
		PostalNumber:     data.PostalNumber,
	}, nil
}

type ValuableData struct {
	Coordinates      Coordinates `json:"coordinates"`
	MunicipalityName string      `json:"municipality_name"`
	PostalPlace      string      `json:"postal_place"`
	PostalNumber     string      `json:"postal_number"`
}

func (vd *ValuableData) GetCoordinates() Coordinates {
	return vd.Coordinates
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Metadata struct {
	Page            int    `json:"side"`
	HitsPerPage     int    `json:"treffPerSide"`
	ShowsTo         int    `json:"viserTil"`
	SokeStreng      string `json:"sokeStreng"`
	AsciiCompatible bool   `json:"asciiKompatibel"`
	ShowsFrom       int    `json:"viserFra"`
	TotalHits       int    `json:"totaltAntallTreff"`
}

type Address struct {
	AddressName                             string                `json:"addressenavn"`
	AddressText                             string                `json:"addressetekst"`
	AddressAdditionalName                   string                `json:"addressetilleggsnavn"`
	AddressCode                             int                   `json:"addressekode"`
	Number                                  int                   `json:"nummer"`
	Letter                                  string                `json:"bokstav"`
	MunicipalityNumber                      string                `json:"kommunenummer"`
	MunicipalityName                        string                `json:"kommunenavn"`
	YardNumber                              int                   `json:"gardsnummer"`
	UseNumber                               int                   `json:"bruksnummer"`
	AttachmentNumber                        int                   `json:"festenummer"`
	SubNumber                               int                   `json:"undernummer"`
	UtilityUnitNumber                       []string              `json:"bruksenhetsnummer"`
	ObjectType                              string                `json:"objtype"`
	PostalPlace                             string                `json:"poststed"`
	PostalNumber                            string                `json:"postnummer"`
	AddressTextWithoutAddressAdditionalName string                `json:"adressetekstutenadressetilleggsnavn"`
	LocationVerified                        bool                  `json:"stedfestingverifisert"`
	PointOfRepresentation                   PointOfRepresentation `json:"representasjonspunkt"`
	UpdateDate                              time.Time             `json:"oppdateringsdato"`
}

func (a *Address) GetPointOfRepresentation() PointOfRepresentation {
	return a.PointOfRepresentation
}

type PointOfRepresentation struct {
	EPSG      string  `json:"epsg"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
