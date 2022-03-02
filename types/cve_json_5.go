package types

type CveJson5Response struct {
	Message string   `json:"message"`
	Created CveJson5 `json:"created"`
}

type CveJson5 struct {
	DataType    string              `json:"dataType"`
	DataVersion string              `json:"dataVersion"`
	CveMetadata CveJson5CveMetadata `json:"cveMetadata"`
	Containers  CveJson5Containers  `json:"containers"`
}

type CveJson5CveMetadata struct {
	Assigner      string `json:"assigner"`
	CveId         string `json:"id"`
	AssignerOrgId string `json:"assignerOrgId,omitempty"`
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
	Id string `json:"id"`
}

type CveJson5ProblemTypes struct {
	Descriptions []CveJson5ProblemTypeDescriptions `json:"descriptions"`
}

type CveJson5Descriptions struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type CveJson5ProblemTypeDescriptions struct {
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

type CveJson5Affected struct {
	Vendor        string             `json:"vendor"`
	Product       string             `json:"product"`
	Versions      []CveJson5Versions `json:"versions"`
	DefaultStatus string             `json:"defaultStatus,omitempty"`
}

type CveJson5Versions struct {
	Version         string `json:"version"`
	Status          string `json:"status"`
	LessThan        string `json:"lessThan,omitempty"`
	LessThanOrEqual string `json:"lessThanOrEqual,omitempty"`
	VersionType     string `json:"versionType,omitempty"`
}

type CveJson5References struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url"`
}
