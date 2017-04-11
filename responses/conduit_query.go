package responses

import "github.com/seamlessdocsdev/gonduit/entities"

// ConduitQueryResponse is the response of calling conduit.query.
type ConduitQueryResponse map[string]*entities.ConduitMethod
