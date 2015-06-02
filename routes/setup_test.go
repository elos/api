package routes_test

type TestParams map[string]string

func (tp TestParams) ByName(name string) string {
	return tp[name]
}
