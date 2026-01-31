package linkdapi

import "fmt"

// Profile Endpoints

// GetProfileOverview gets basic profile information by username.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/overview
func (c *Client) GetProfileOverview(username string) (map[string]any, error) {
	params := map[string]string{"username": username}
	return c.sendRequest("GET", "api/v1/profile/overview", params)
}

// GetProfileDetails gets profile details information by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/details
func (c *Client) GetProfileDetails(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/details", params)
}

// GetContactInfo gets contact details for a profile by username.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/contact-info
func (c *Client) GetContactInfo(username string) (map[string]any, error) {
	params := map[string]string{"username": username}
	return c.sendRequest("GET", "api/v1/profile/contact-info", params)
}

// GetFullExperience gets complete work experience by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/full-experience
func (c *Client) GetFullExperience(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/full-experience", params)
}

// GetCertifications gets lists of professional certifications by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/certifications
func (c *Client) GetCertifications(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/certifications", params)
}

// GetEducation gets full education information by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/education
func (c *Client) GetEducation(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/education", params)
}

// GetSkills gets profile skills by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/skills
func (c *Client) GetSkills(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/skills", params)
}

// GetSocialMatrix gets social network metrics by username.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/social-matrix
func (c *Client) GetSocialMatrix(username string) (map[string]any, error) {
	params := map[string]string{"username": username}
	return c.sendRequest("GET", "api/v1/profile/social-matrix", params)
}

// GetRecommendations gets profile given and received recommendations by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/recommendations
func (c *Client) GetRecommendations(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/recommendations", params)
}

// GetSimilarProfiles gets similar profiles for a given profile using its URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/similar
func (c *Client) GetSimilarProfiles(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/similar", params)
}

// GetProfileAbout gets about this profile such as last update and verification info.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/about
func (c *Client) GetProfileAbout(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/about", params)
}

// GetProfileReactions gets all reactions for given profile by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/reactions
func (c *Client) GetProfileReactions(urn string, cursor string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	stringParam(params, "cursor", cursor)
	return c.sendRequest("GET", "api/v1/profile/reactions", params)
}

// GetProfileInterests gets profile interests by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/interests
func (c *Client) GetProfileInterests(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/interests", params)
}

// GetFullProfile gets full profile data in 1 request (everything included).
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/full
func (c *Client) GetFullProfile(username, urn string) (map[string]any, error) {
	if username == "" && urn == "" {
		return nil, fmt.Errorf("either username or urn must be provided")
	}

	params := make(map[string]string)
	stringParam(params, "username", username)
	stringParam(params, "urn", urn)

	return c.sendRequest("GET", "api/v1/profile/full", params)
}

// GetProfileServices gets profile services by URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/services
func (c *Client) GetProfileServices(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/profile/services", params)
}

// GetProfileURN gets profile URN from username.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/profile/username-to-urn
func (c *Client) GetProfileURN(username string) (map[string]any, error) {
	params := map[string]string{"username": username}
	return c.sendRequest("GET", "api/v1/profile/username-to-urn", params)
}

// GetProfilePostedJobs gets all jobs posted by a profile using its URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/jobs/posted-by-profile
func (c *Client) GetProfilePostedJobs(profileUrn string, start, count int) (map[string]any, error) {
	params := map[string]string{"profileUrn": profileUrn}
	intParam(params, "start", start)
	intParam(params, "count", count)
	return c.sendRequest("GET", "api/v1/jobs/posted-by-profile", params)
}
