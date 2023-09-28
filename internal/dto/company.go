package dto

type CompanyDto struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Document string            `json:"document"`
	Address  CompanyAddressDto `json:"address"`
	Status   string            `json:"status"`
}

type CompanyAddressDto struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type AddressDto struct {
	ID           string `json:"id"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
	CompanyId    string `json:"company_id"`
}

type AddressResponseDto struct {
	ID           string `json:"id"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}
