package responses

import "github.com/seamlessdocsdev/gonduit/entities"

// PHIDLookupResponse is the result of phid.lookup operations.
type PHIDLookupResponse map[string]*entities.PHIDResult
