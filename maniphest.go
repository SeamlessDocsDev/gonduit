package gonduit

import (
	"github.com/seamlessdocsdev/gonduit/entities"
	"github.com/seamlessdocsdev/gonduit/requests"
	"github.com/seamlessdocsdev/gonduit/responses"
)

// ManiphestQuery performs a call to maniphest.query.
func (c *Conn) ManiphestQuery(
	req requests.ManiphestQueryRequest,
) (*responses.ManiphestQueryResponse, error) {
	var res responses.ManiphestQueryResponse

	if err := c.Call("maniphest.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ManiphestCreateTask performs a call to maniphest.createtask.
func (c *Conn) ManiphestCreateTask(
	req requests.ManiphestCreateTaskRequest,
) (*entities.ManiphestTask, error) {
	var res entities.ManiphestTask

	if err := c.Call("maniphest.createtask", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Conn) ManiphestGetTaskTransactions(
	req requests.ManiphestGetTaskTransactions,
) (*responses.ManiphestGetTaskTransactionsResponse, error) {
	var res responses.ManiphestGetTaskTransactionsResponse

	if err := c.Call("maniphest.gettasktransactions", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
