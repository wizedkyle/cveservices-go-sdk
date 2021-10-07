package types

type CveJson4 struct {
	DataType    *interface{}         `json:"data_type"`
	DataFormat  *interface{}         `json:"data_format"`
	DataVersion *interface{}         `json:"data_version"`
	CVEDataMeta *CveJson4CveDataMeta `json:"CVE_data_meta"`
	Affects     *CveJson4Affects     `json:"affects"`
	Problemtype *CveJson4Problemtype `json:"problemtype"`
	References  *CveJson4References  `json:"references"`
	Description *CveJson4Description `json:"description"`
}

type CveJson41 struct {
	DataType    *interface{}          `json:"data_type"`
	DataFormat  *interface{}          `json:"data_format"`
	DataVersion *interface{}          `json:"data_version"`
	CVEDataMeta *CveJson41CveDataMeta `json:"CVE_data_meta"`
	Affects     *interface{}          `json:"affects,omitempty"`
	Description *CveJson4Description  `json:"description"`
	Problemtype *interface{}          `json:"problemtype,omitempty"`
	References  *interface{}          `json:"references,omitempty"`
}

type CveJson41CveDataMeta struct {
	ID       string       `json:"ID"`
	ASSIGNER string       `json:"ASSIGNER"`
	STATE    *interface{} `json:"STATE,omitempty"`
}

type CveJson4Affects struct {
	Vendor *CveJson4AffectsVendor `json:"vendor"`
}

type CveJson4AffectsVendor struct {
	VendorData []CveJson4AffectsVendorVendorData `json:"vendor_data"`
}

type CveJson4AffectsVendorProduct struct {
	ProductData []Product `json:"product_data"`
}

type CveJson4AffectsVendorVendorData struct {
	VendorName string                        `json:"vendor_name"`
	Product    *CveJson4AffectsVendorProduct `json:"product"`
}

type CveJson4CveDataMeta struct {
	ID       string `json:"ID"`
	ASSIGNER string `json:"ASSIGNER"`
}

type CveJson4Description struct {
	DescriptionData []LangString `json:"description_data"`
}

type CveJson4Problemtype struct {
	ProblemtypeData []CveJson4ProblemtypeProblemtypeData `json:"problemtype_data"`
}

type CveJson4ProblemtypeProblemtypeData struct {
	Description []LangString `json:"description"`
}

type CveJson4References struct {
	ReferenceData []Reference `json:"reference_data"`
}

type LangString struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Product struct {
	ProductName string          `json:"product_name"`
	Version     *ProductVersion `json:"version"`
}

type ProductVersion struct {
	VersionData []ProductVersionVersionData `json:"version_data"`
}

type ProductVersionVersionData struct {
	VersionValue string `json:"version_value"`
}

type Reference struct {
	Url string `json:"url"`
}
