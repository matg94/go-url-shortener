package redis

type RedisConnectionMock struct {
	ReturnErrorGet error
	ReturnErrorSet error
	ReturnErrorDel error
	ReturnValue    string
}

func (r RedisConnectionMock) GET(key string) (string, error) {
	if r.ReturnErrorGet != nil {
		return "", r.ReturnErrorGet
	}
	return r.ReturnValue, nil
}

func (r RedisConnectionMock) SET(key, value string) error {
	if r.ReturnErrorSet != nil {
		return r.ReturnErrorSet
	}
	return nil
}

func (r RedisConnectionMock) DEL(keys ...string) error {
	if r.ReturnErrorDel != nil {
		return r.ReturnErrorDel
	}
	return nil
}
