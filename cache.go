package main

// Model for application wide cache
type Cache struct {
	NextMap *string
	PrevMap *string
}

func newCache() Cache {
	return Cache{}
}
