package exception


/**
	异常捕获
	@author Bill
 */
func Try(fun func(), handler func()) {
	defer func() {
		if err := recover(); err != nil {
			handler()
		}
	}()
	fun()
}