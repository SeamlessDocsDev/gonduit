package gonduit

import (
	"github.com/seamlessdocsdev/gonduit/requests"
	"github.com/seamlessdocsdev/gonduit/responses"
)

// ProjectQuery performs a call to project.query.
func (c *Conn) ProjectQuery(
	req requests.ProjectQueryRequest,
) (*responses.ProjectQueryResponse, error) {
	var res responses.ProjectQueryResponse

	if err := c.Call("project.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
