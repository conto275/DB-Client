package stations

import (
	"time"
)

// type StationsByIDResponse struct {
// 	Stations []Station `json:""`
// }
// type StationByIDResponse struct {
// 	Station Station `json:""`
// }

// Station представляет модель данных для станции.
type Station struct {
	StationID              string          `json:"stationID"`
	Names                  map[string]Name `json:"names"`
	Metropolis             Metropolis      `json:"metropolis"`
	MunicipalityKey        string          `json:"municipalityKey"`
	OpeningHours           string          `json:"openingHours"`
	Owner                  Owner           `json:"owner"`
	Position               Position        `json:"position"`
	Roofing                string          `json:"roofing"`
	State                  string          `json:"state"`
	StationCategory        string          `json:"stationCategory"`
	AvailableLocalServices []string        `json:"availableLocalServices"`
	AvailableTransports    []string        `json:"availableTransports"`
	CountryCode            string          `json:"countryCode"`
	TransportAssociations  []string        `json:"transportAssociations"`
	TimeZone               string          `json:"timeZone"`
	ValidFrom              time.Time       `json:"validFrom"`
	ValidTo                time.Time       `json:"validTo"`
	Address                Address         `json:"address"`
}

// Name представляет модель данных для имени станции.
type Name struct {
	ID int `json:"id"`
}

// Metropolis представляет модель данных для метрополии.
type Metropolis struct {
	ID int `json:"id"`
}

// Address представляет модель данных для адреса станции.
type Address struct {
	Street                string `json:"street"`
	HouseNumber           string `json:"houseNumber"`
	PostalCode            string `json:"postalCode"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Country               string `json:"country"`
	AdditionalInformation string `json:"additionalInformation"`
	Website               string `json:"website"`
}

// Owner представляет модель данных для владельца станции.
type Owner struct {
	Name               string             `json:"name"`
	OrganisationalUnit OrganisationalUnit `json:"organisationalUnit"`
}

// OrganisationalUnit представляет модель данных для организационной единицы.
type OrganisationalUnit struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	NameShort string `json:"nameShort"`
}

// Position представляет модель данных для координат станции.
type Position struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// func (s Station) Info() string {
// 	return fmt.Sprintf("[ID] %s | [Names] %+v | [Metropolis] %+v | [MunicipalityKey] %s | [OpeningHours] %s | [Owner] %+v | [Position] %+v | [Roofing] %s | [State] %s | [StationCategory] %s | [AvailableLocalServices] %+v | [AvailableTransports] %+v | [CountryCode] %s | [TransportAssociations] %+v | [TimeZone] %s | [ValidFrom] %s | [ValidTo] %s | [Address.Street] %s | [Address.HouseNumber] %s | [Address.PostalCode] %s | [Address.City] %s | [Address.State] %s | [Address.Country] %s | [Address.AdditionalInformation] %s | [Address.Website] %s",

// 		s.StationID,
// 		s.Names,
// 		s.Metropolis,
// 		s.MunicipalityKey,
// 		s.OpeningHours,
// 		s.Owner,
// 		s.Position,
// 		s.Roofing,
// 		s.State,
// 		s.StationCategory,
// 		s.AvailableLocalServices,
// 		s.AvailableTransports,
// 		s.CountryCode,
// 		s.TransportAssociations,
// 		s.TimeZone,
// 		s.ValidFrom,
// 		s.ValidTo,
// 		s.Address.Street,
// 		s.Address.HouseNumber,
// 		s.Address.PostalCode,
// 		s.Address.City,
// 		s.Address.State,
// 		s.Address.Country,
// 		s.Address.AdditionalInformation,
// 		s.Address.Website)
// }
