package sign

func SignIdcard(idcard string) string {
	cp := idcard
	leth := len(cp)
	return cp[0:4] + " **** **** " + cp[leth-4:]
}

func SignMobile(mobile string) string {
	cp := mobile
	leth := len(cp)
	return cp[0:3] + " **** " + cp[leth-4:]
}
