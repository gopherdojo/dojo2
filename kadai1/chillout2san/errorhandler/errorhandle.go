package errorhandle

import "log"

/*
	errがあればlogに出力します。
*/
func AlertError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
