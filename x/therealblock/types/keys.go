package types

const (
	// ModuleName defines the module name
	ModuleName = "therealblock"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_therealblock"

	// ProjectKey is the store keyprefix string for project
	ProjectKey = "Project/value/"

	// ProjectCountKey is the store key for the count of all projects
	ProjectCountKey = "Project/count/"

	// ProjectCreatedEvent is the event type for project creation
	ProjectCreatedEventType    = "project_created"

	// General event attributes
	ProjectEventProjectKey     = "id"
	ProjectEventProjectCreator = "creator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
