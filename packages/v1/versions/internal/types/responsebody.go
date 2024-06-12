package types

type ResponseBody struct {
	// The semantic version of the API
	APIVersion string `json:"apiVersion" example:"1.1.0"`
	// The API environment
	Environment string `json:"environment" example:"production"`
}
