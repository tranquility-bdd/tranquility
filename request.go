package tranquility

//Request provides an abstraction for a full block of PreAction, Action and Test
type Request struct {
	PreAction func(action *Action)
	Action    *Action
	Test      func(res *Response)
}

//Run perfoms the full request in the following order PreAction, Action and Test
func (req Request) Run() {
	req.PreAction(req.Action)
	res, err := req.Action.Run()
	if err != nil {
		panic(err)
	}
	req.Test(res)
}
