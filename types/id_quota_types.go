package types

type IdQuotaResponse struct {
	IdQuota       int `json:"id_quota"`
	TotalReserved int `json:"total_reserved"`
	Available     int `json:"available"`
}
