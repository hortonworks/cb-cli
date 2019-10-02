package freeipa

var listHeader = []string{"Name", "Status", "Domain"}

type freeIpaDetails struct {
	Name           string `json:"Name" yaml:"Name"`
	Status         string `json:"Status" yaml:"Status"`
	Domain         string `json:"Domain" yaml:"Domain"`
	CRN            string `json:"CRN" yaml:"CRN"`
	EnvironmentCrn string `json:"EnvironmentCrn" yaml:"EnvironmentCrn"`
}

func (ipa *freeIpaDetails) DataAsStringArray() []string {
	return []string{ipa.Name, ipa.Status, ipa.Domain}
}
