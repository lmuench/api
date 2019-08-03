package store

import (
	"encoding/json"
	"io/ioutil"
)

const filename = ".api_store"

func Get(key string) (string, bool) {
	kv := parseOrCreateMap()
	val, ok := kv[key]
	return val, ok
}

func Set(key, val string) error {
	kv := parseOrCreateMap()
	kv[key] = val

	b, err := json.Marshal(kv)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0600)
	return err
}

func Delete(key string) error {
	kv := parseOrCreateMap()
	delete(kv, key)

	b, err := json.Marshal(kv)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0600)
	return err
}

func parseOrCreateMap() map[string]string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return make(map[string]string)
	}

	var kv map[string]string
	err = json.Unmarshal(b, &kv)
	if err != nil {
		return make(map[string]string)
	}
	return kv
}
