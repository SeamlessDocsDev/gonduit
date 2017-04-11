package responses

import "github.com/seamlessdocsdev/gonduit/entities"

// DifferentialQueryResponse is the response of calling differential.query.
type DifferentialQueryResponse []*entities.DifferentialRevision
