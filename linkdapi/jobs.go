package linkdapi

// Jobs Endpoints

// SearchJobs searches for jobs with various filters.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/search
func (c *Client) SearchJobs(searchParams JobSearchParams) (map[string]any, error) {
	params := make(map[string]string)

	stringParam(params, "keyword", searchParams.Keyword)
	stringParam(params, "location", searchParams.Location)
	stringParam(params, "geoId", searchParams.GeoID)
	sliceParam(params, "companyIds", searchParams.CompanyIDs)
	sliceParam(params, "jobTypes", searchParams.JobTypes)
	sliceParam(params, "experience", searchParams.Experience)
	sliceParam(params, "regions", searchParams.Regions)
	stringParam(params, "timePosted", searchParams.TimePosted)
	stringParam(params, "salary", searchParams.Salary)
	sliceParam(params, "workArrangement", searchParams.WorkArrangement)
	intParam(params, "start", searchParams.Start)

	return c.sendRequest("GET", "api/v1/jobs/search", params)
}

// GetJobDetails gets job details by job ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/job/details
func (c *Client) GetJobDetails(jobID string) (map[string]any, error) {
	params := map[string]string{"jobId": jobID}
	return c.sendRequest("GET", "api/v1/jobs/job/details", params)
}

// GetSimilarJobs gets similar jobs by job ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/job/similar
func (c *Client) GetSimilarJobs(jobID string) (map[string]any, error) {
	params := map[string]string{"jobId": jobID}
	return c.sendRequest("GET", "api/v1/jobs/job/similar", params)
}

// GetPeopleAlsoViewedJobs gets related jobs that people also viewed.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/job/people-also-viewed
func (c *Client) GetPeopleAlsoViewedJobs(jobID string) (map[string]any, error) {
	params := map[string]string{"jobId": jobID}
	return c.sendRequest("GET", "api/v1/jobs/job/people-also-viewed", params)
}

// GetJobDetailsV2 gets job details V2 by job ID. This endpoint supports all job statuses
// (open, closed, expired, etc.) and provides detailed information about the job.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/job/details-v2
func (c *Client) GetJobDetailsV2(jobID string) (map[string]any, error) {
	params := map[string]string{"jobId": jobID}
	return c.sendRequest("GET", "api/v1/jobs/job/details-v2", params)
}

// SearchJobsV2 searches for jobs V2 with comprehensive filters (all filters available).
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/search/jobs
func (c *Client) SearchJobsV2(searchParams JobSearchV2Params) (map[string]any, error) {
	params := make(map[string]string)

	stringParam(params, "keyword", searchParams.Keyword)
	intParam(params, "start", searchParams.Start)
	stringParam(params, "sortBy", searchParams.SortBy)
	stringParam(params, "datePosted", searchParams.DatePosted)
	sliceParam(params, "experience", searchParams.Experience)
	sliceParam(params, "jobTypes", searchParams.JobTypes)
	sliceParam(params, "workplaceTypes", searchParams.WorkplaceTypes)
	stringParam(params, "salary", searchParams.Salary)
	sliceParam(params, "companies", searchParams.Companies)
	sliceParam(params, "industries", searchParams.Industries)
	sliceParam(params, "locations", searchParams.Locations)
	sliceParam(params, "functions", searchParams.Functions)
	sliceParam(params, "titles", searchParams.Titles)
	sliceParam(params, "Benefits", searchParams.Benefits)
	sliceParam(params, "commitments", searchParams.Commitments)
	boolParam(params, "easyApply", searchParams.EasyApply)
	boolParam(params, "verifiedJob", searchParams.VerifiedJob)
	boolParam(params, "under10Applicants", searchParams.Under10Applicants)
	boolParam(params, "fairChance", searchParams.FairChance)

	return c.sendRequest("GET", "api/v1/search/jobs", params)
}

// GetHiringTeam gets the hiring team for a given job by job ID.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/job/hiring-team
func (c *Client) GetHiringTeam(jobID string, start int) (map[string]any, error) {
	params := map[string]string{"jobId": jobID}
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/jobs/job/hiring-team", params)
}
