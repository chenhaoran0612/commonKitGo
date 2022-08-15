package cache

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type session struct {
	db int
}

// Set a key/value  过期单位秒
// 设置过期需要注意下：https://segmentfault.com/a/1190000023299657
func (d *session) Set(key string, data interface{}, time int) error {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return err
	}
	conn := r.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time) //EXPIRE 单位为秒； PEXPIRE 单位为毫秒
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *session) SetStr(key string, value string, time int) error {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return err
	}
	conn := r.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time) //EXPIRE 单位为秒； PEXPIRE 单位为毫秒
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists check a key
func (d *session) Exists(key string) bool {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return false
	}
	conn := r.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}
func (d *session) SetByte(key string, value []byte, time int) error {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return err
	}
	conn := r.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time) //EXPIRE 单位为秒； PEXPIRE 单位为毫秒
	if err != nil {
		return err
	}
	return nil
}

func (d *session) SetByteNoTimeout(key string, value []byte) error {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return err
	}
	conn := r.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func (d *session) Get(key string) ([]byte, error) {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return nil, err
	}
	conn := r.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (d *session) GetExpire(key string) (int, error) {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return 0, err
	}
	conn := r.Get()
	defer conn.Close()
	reply, err := redis.Int(conn.Do("TTL", key))

	return reply, nil
}

// Delete delete a kye
func (d *session) Delete(key string) (bool, error) {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return false, err
	}
	conn := r.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func (d *session) LikeDeletes(key string) error {
	r, err := getRedisPoll(d.db)
	if err != nil {
		return err
	}
	conn := r.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = d.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
