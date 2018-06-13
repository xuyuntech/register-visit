package manager

func (m *DefaultManager) ChainQuery(queryString string) (string, error) {
	return m.fabricSetup.Query(queryString)
}
func (m *DefaultManager) ChainSetupChannel() error {
	return m.fabricSetup.SetupChannel()
}
func (m *DefaultManager) InstallAndInstantiateCC() error {
	return m.fabricSetup.InstallAndInstantiateCC()
}
