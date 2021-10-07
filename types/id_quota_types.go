package types

type IdQuotaResponse struct {
	IdQuota       int32 `json:"id_quota"`
	TotalReserved int32 `json:"total_reserved"`
	Available     int32 `json:"available"`
}
