package tlsloader

import "io"

func ReadFrom(s Storage, cb func(r io.Reader) error) (err error) {
	var r io.ReadCloser
	if r, err = s.Reader(); err != nil {
		return
	}
	func() {
		defer func() {
			if err == nil {
				err = r.Close()
			} else {
				r.Close()
			}
		}()
		err = cb(r)
	}()
	return
}
