package cache 


type cache struct {
	key string 
	value interface{}
}



func New() cache{
	return cache{
		key: "",
		value: nil,
	}
}


func (c *cache) Set(key string, value interface{}){
	c.key = key
	c.value = value
}

func (c cache) Get(key string) interface{}{
	return cache{
		key: c.key, 
		value: c.value,
	}
}

func (c *cache) Delete(key string){
	c.key = ""
	c.value = nil

}








