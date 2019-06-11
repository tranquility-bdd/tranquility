package tranquility

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
	delete(env,key)
}

// Init initializes the environment with the given map
func (env environment) Init(initMap map[string]string) {
	for k,v := range initMap {
		env.Set(k,v)
	}
}
