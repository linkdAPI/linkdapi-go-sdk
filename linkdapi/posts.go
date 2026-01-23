package linkdapi

// Posts Endpoints

// GetFeaturedPosts gets all featured posts for a given profile using its URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/posts/featured
func (c *Client) GetFeaturedPosts(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/posts/featured", params)
}

// GetAllPosts retrieves all posts for a given profile URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/posts/all
func (c *Client) GetAllPosts(urn string, cursor string, start int) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	stringParam(params, "cursor", cursor)
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/posts/all", params)
}

// GetPostInfo retrieves information about a specific post using its URN.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/posts/info
func (c *Client) GetPostInfo(urn string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	return c.sendRequest("GET", "api/v1/posts/info", params)
}

// GetPostComments gets comments for a specific LinkedIn post.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/posts/comments
func (c *Client) GetPostComments(urn string, start int, count int, cursor string) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	intParam(params, "start", start)
	intParam(params, "count", count)
	stringParam(params, "cursor", cursor)
	return c.sendRequest("GET", "api/v1/posts/comments", params)
}

// GetPostLikes retrieves all users who liked or reacted to a given post.
//
// Documentation: https://linkdapi.com/docs?endpoint=/api/v1/posts/likes
func (c *Client) GetPostLikes(urn string, start int) (map[string]any, error) {
	params := map[string]string{"urn": urn}
	intParam(params, "start", start)
	return c.sendRequest("GET", "api/v1/posts/likes", params)
}
