package yarn

func (p *YarnProvider) GetCredentialParameters(stringFinder func(string) string) (map[string]interface{}, error) {
	var credentialMap = make(map[string]interface{})
	credentialMap["yarnEndpoint"] = stringFinder("yarn-endpoint")
	return credentialMap, nil
}
