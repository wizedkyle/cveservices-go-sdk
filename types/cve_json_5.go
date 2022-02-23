package types

type CveJson5Basic struct {
	DataType    string              `json:"dataType"`
	DataVersion string              `json:"dataVersion"`
	CveMetadata CveJson5CveMetadata `json:"cveMetadata"`
	Containers  CveJson5Containers  `json:"containers"`
}

type CveJson5CveMetadata struct {
	CveId         string `json:"cveId"`
	AssignerOrgId string `json:"assignerOrgId"`
	State         string `json:"state"`
}

type CveJson5Containers struct {
	Cna CveJson5Cna `json:"cna"`
}

type CveJson5Cna struct {
	ProviderMetadata CveJson5ProviderMetadata `json:"providerMetadata"`
	ProblemTypes     []CveJson5ProblemTypes   `json:"problemTypes"`
	Affected         []CveJson5Affected       `json:"affected"`
	Descriptions     []CveJson5Descriptions   `json:"descriptions"`
	References       []CveJson5References     `json:"references"`
}

type CveJson5ProviderMetadata struct {
	OrgId string `json:"orgId"`
}

type CveJson5ProblemTypes struct {
	Descriptions []CveJson5Descriptions `json:"descriptions"`
}

type CveJson5Descriptions struct {
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

type CveJson5Affected struct {
	Vendor        string             `json:"vendor"`
	Product       string             `json:"product"`
	Versions      []CveJson5Versions `json:"versions"`
	DefaultStatus string             `json:"defaultStatus"`
}

type CveJson5Versions struct {
	Version     string `json:"version"`
	Status      string `json:"status"`
	LessThan    string `json:"lessThan"`
	VersionType string `json:"versionType"`
}

type CveJson5References struct {
	Url string `json:"url"`
}
