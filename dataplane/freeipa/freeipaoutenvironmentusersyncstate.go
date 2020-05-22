package freeipa

var environmentUserSyncStateHeader = []string{"State", "LastUserSyncOperationID"}

type freeIpaOutEnvironmentUserSyncState struct {
	State                   string `json:"State" yaml:"State"`
	LastUserSyncOperationID string `json:"LastUserSyncOperationID" yaml:"LastUserSyncOperationID"`
}

func (s *freeIpaOutEnvironmentUserSyncState) DataAsStringArray() []string {
	return []string{s.State, s.LastUserSyncOperationID}
}
