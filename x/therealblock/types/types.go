package types

const (
	ProjectStateDraft = "draft"

	ProjectStateActive = "active"

	ProjectStatePending = "pending"

	ProjectStateFunded = "funded"

	ProjectStateCompleted = "completed"

	ProjectStateCancelled = "cancelled"
)

func IsValidState(state string) error {
	var set = map[string]bool{
		ProjectStateDraft:     true,
		ProjectStateActive:    true,
		ProjectStateFunded:    true,
		ProjectStateCompleted: true,
		ProjectStateCancelled: true,
	}
	_, found := set[state]
	if !found {
		return ErrInvalidState
	}
	return nil
}
