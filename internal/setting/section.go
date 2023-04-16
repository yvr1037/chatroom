package setting

import "time"

type ServerSetting struct {
	Runmode      string
	HttpPort     string
	ReadTimeout  time.Duration  // 读超时时间
	WriteTimeout time.Duration  // 写超时时间
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    string
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
