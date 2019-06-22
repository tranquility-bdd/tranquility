package tranquility

//Collection provides an abstraction for a group of Requests defined by a user
type Collection []*Request

//Run perfoms the full request in the following order PreAction, Action and Test
func (col Collection) Run() {
	for i := range col {
		col[i].Run()
	}
}
