package linkdapi

// Comments Endpoints

// GetAllComments retrieves all comments made by a profile using their URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/comments/all
func (c *Client) GetAllComments(urn string, cursor string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	stringParam(params, "cursor", cursor)
	return c.sendRequest("GET", "api/v1/comments/all", params)
}

// GetCommentLikes gets all users who reacted to one or more comment URNs.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/comments/likes
func (c *Client) GetCommentLikes(urns string, start int) (map[string]any, error) {
	params := map[string]string{"urn": urns}
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/comments/likes", params)
}
