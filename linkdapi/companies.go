package linkdapi

import "fmt"

// Companies Endpoints

// CompanyNameLookup searches companies by name.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/name-lookup
func (c *Client) CompanyNameLookup(query string) (map[string]any, error) {
	params := map[string]string{"query": query}
	return c.sendRequest("GET", "api/v1/companies/name-lookup", params)
}

// GetCompanyInfo gets company details either by ID or name.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/info
func (c *Client) GetCompanyInfo(companyID, name string) (map[string]any, error) {
	if companyID == "" && name == "" {
		return nil, fmt.Errorf("either companyID or name must be provided")
	}

	params := make(map[string]string)
	stringParam(params, "id", companyID)
	stringParam(params, "name", name)

	return c.sendRequest("GET", "api/v1/companies/company/info", params)
}

// GetSimilarCompanies gets similar companies by ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/similar
func (c *Client) GetSimilarCompanies(companyID string) (map[string]any, error) {
	params := map[string]string{"id": companyID}
	return c.sendRequest("GET", "api/v1/companies/company/similar", params)
}

// GetCompanyEmployeesData gets company employees data by ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/employees-data
func (c *Client) GetCompanyEmployeesData(companyID string) (map[string]any, error) {
	params := map[string]string{"id": companyID}
	return c.sendRequest("GET", "api/v1/companies/company/employees-data", params)
}

// GetCompanyJobs gets available job listings for given companies by ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/jobs
func (c *Client) GetCompanyJobs(companyIDs []string, start int) (map[string]any, error) {
	params := make(map[string]string)
	sliceParam(params, "companyIDs", companyIDs)
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/companies/jobs", params)
}

// GetCompanyAffiliatedPages gets affiliated pages/subsidiaries of a company by ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/affiliated-pages
func (c *Client) GetCompanyAffiliatedPages(companyID string) (map[string]any, error) {
	params := map[string]string{"id": companyID}
	return c.sendRequest("GET", "api/v1/companies/company/affiliated-pages", params)
}

// GetCompanyPosts gets posts of a company by ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/posts
func (c *Client) GetCompanyPosts(companyID string, start int) (map[string]any, error) {
	params := map[string]string{"id": companyID}
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/companies/company/posts", params)
}

// GetCompanyID gets ID of a company by universal_name (username).
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/universal-name-to-id
func (c *Client) GetCompanyID(universalName string) (map[string]any, error) {
	params := map[string]string{"universalName": universalName}
	return c.sendRequest("GET", "api/v1/companies/company/universal-name-to-id", params)
}

// GetCompanyDetailsV2 gets company details V2 with extended information by company ID.
// This endpoint returns more information about the company including
// peopleAlsoFollow, affiliatedByJobs, etc.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/companies/company/info-v2
func (c *Client) GetCompanyDetailsV2(companyID string) (map[string]any, error) {
	params := map[string]string{"id": companyID}
	return c.sendRequest("GET", "api/v1/companies/company/info-v2", params)
}
