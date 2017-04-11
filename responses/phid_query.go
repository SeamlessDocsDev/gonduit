package responses

import "github.com/seamlessdocsdev/gonduit/entities"

// PHIDQueryResponse is the result of phid.query operations.
type PHIDQueryResponse map[string]*entities.PHIDResult
