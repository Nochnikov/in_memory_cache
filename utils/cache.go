package cache 


type cache struct {
	values map[string]interface{}
}



func  New() cache{
	return cache{
		values: make(map[string]interface{}),
	}
}


func (c *cache) Set(key string, value interface{}){
	c.values[key] = value
}

func (c cache) Get(key string) interface{}{
	return c.values[key]
}

func (c *cache) Delete(key string){
	delete(c.values, key)
}








