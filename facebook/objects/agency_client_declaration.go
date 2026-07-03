package objects

type AgencyClientDeclaration struct {
	AgencyRepresentingClient        *uint64 `json:"agency_representing_client,omitempty"`
	ClientBasedInFrance             *uint64 `json:"client_based_in_france,omitempty"`
	ClientCity                      *string `json:"client_city,omitempty"`
	ClientCountryCode               *string `json:"client_country_code,omitempty"`
	ClientEmailAddress              *string `json:"client_email_address,omitempty"`
	ClientName                      *string `json:"client_name,omitempty"`
	ClientPostalCode                *string `json:"client_postal_code,omitempty"`
	ClientProvince                  *string `json:"client_province,omitempty"`
	ClientStreet                    *string `json:"client_street,omitempty"`
	ClientStreet2                   *string `json:"client_street2,omitempty"`
	HasWrittenMandateFromAdvertiser *uint64 `json:"has_written_mandate_from_advertiser,omitempty"`
	IsClientPayingInvoices          *uint64 `json:"is_client_paying_invoices,omitempty"`
}

var AgencyClientDeclarationFields = struct {
	AgencyRepresentingClient        string
	ClientBasedInFrance             string
	ClientCity                      string
	ClientCountryCode               string
	ClientEmailAddress              string
	ClientName                      string
	ClientPostalCode                string
	ClientProvince                  string
	ClientStreet                    string
	ClientStreet2                   string
	HasWrittenMandateFromAdvertiser string
	IsClientPayingInvoices          string
}{
	AgencyRepresentingClient:        "agency_representing_client",
	ClientBasedInFrance:             "client_based_in_france",
	ClientCity:                      "client_city",
	ClientCountryCode:               "client_country_code",
	ClientEmailAddress:              "client_email_address",
	ClientName:                      "client_name",
	ClientPostalCode:                "client_postal_code",
	ClientProvince:                  "client_province",
	ClientStreet:                    "client_street",
	ClientStreet2:                   "client_street2",
	HasWrittenMandateFromAdvertiser: "has_written_mandate_from_advertiser",
	IsClientPayingInvoices:          "is_client_paying_invoices",
}
