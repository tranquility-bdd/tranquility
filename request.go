package tranquility

//Request provides an abstraction for a full block of PreAction, Action and Test
type Request struct {
	PreAction func(action *Action)
	Action    *Action
	Test      func(res *Response)
}

//verifyAction initializes parameters and headers if they are null to avoid errors with PreActions
func verifyAction(req *Request) {
	if req.Action.Parameters == nil {
		req.Action.Parameters = map[string]string{}
	}
	if req.Action.Headers == nil {
		req.Action.Headers = map[string]string{}
	}
}

//Run perfoms the full request in the following order PreAction, Action and Test
func (req Request) Run() {
	verifyAction(&req)
	if req.PreAction != nil {
		req.PreAction(req.Action)
	}
	res, err := req.Action.Run()
	if err != nil {
		panic(err)
	}
	if req.Test != nil {
		req.Test(res)
	}
}
