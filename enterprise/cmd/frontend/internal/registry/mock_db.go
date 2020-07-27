package registry

func resetMocks() {
	mocks = dbMocks{}
}

type dbMocks struct { /* all structs must go */ }

var mocks dbMocks
