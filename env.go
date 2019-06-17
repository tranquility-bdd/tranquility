package tranquility

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type environment map[string]string

//Env is the global enviroment to be shared across actions
var Env = environment{}

// Set modifies/adds the new key - value in the global environment
func (env environment) Set(key, value string) {
	env[key] = value
}

// Get returns the value associated with key in case it exists, otherwise ""
func (env environment) Get(key string) string {
	return env[key]
}

// Del deletes the key and the value associated from the map
func (env environment) Del(key string) {
	delete(env, key)
}

// Init initializes the environment with the given map
func (env environment) Init(initMap map[string]string) {
	for k, v := range initMap {
		env.Set(k, v)
	}
}

// Representation of a Postman Environment json file
type postmanEnv struct {
	Values []postmanEnvVar `json:"values"`
}

// postmanEnvVar is the representation of a Postman environment value from a json
type postmanEnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ReadPostmanEnv initializes the environment with the given json format postman environment file
func (env environment) ReadPostmanEnv(file string) {
	envFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer envFile.Close()
	jsonBytes, err := ioutil.ReadAll(envFile)
	if err != nil {
		panic(err)
	}
	var loadedEnv postmanEnv
	json.Unmarshal(jsonBytes, &loadedEnv)

	for i := range loadedEnv.Values {
		env.Set(loadedEnv.Values[i].Key, loadedEnv.Values[i].Value)
	}
}
