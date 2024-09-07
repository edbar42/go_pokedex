package api

import (
	"testing"
	"time"
)

func TestCache_AddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Second)

	regions := Regions{
		Count: 1,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "Region1", URL: "http://example.com/region1"},
		},
	}

	mapCache := NewMapCache(regions)
	url := "http://example.com"
	cache.Add(url, mapCache)

	retrievedCache, ok := cache.Get(url)
	if !ok {
		t.Fatalf("Expected cache to contain entry for %s", url)
	}

	if retrievedCache.Data.Count != regions.Count {
		t.Errorf("Expected count %d, got %d", regions.Count, retrievedCache.Data.Count)
	}
	if len(retrievedCache.Data.Results) != len(regions.Results) {
		t.Errorf("Expected %d results, got %d", len(regions.Results), len(retrievedCache.Data.Results))
	}
	if retrievedCache.Data.Results[0].Name != regions.Results[0].Name {
		t.Errorf("Expected name %s, got %s", regions.Results[0].Name, retrievedCache.Data.Results[0].Name)
	}
}

func TestCache_ReapLoop(t *testing.T) {
	cacheTimeout := 100 * time.Millisecond
	cache := NewCache(cacheTimeout)

	url := "http://example.com"
	mapCache := NewMapCache(Regions{})
	cache.Add(url, mapCache)

	if _, ok := cache.Get(url); !ok {
		t.Fatalf("Expected cache to contain entry for %s", url)
	}

	time.Sleep(cacheTimeout + 10*time.Millisecond)

	if _, ok := cache.Get(url); ok {
		t.Fatalf("Expected cache to not contain entry for %s after expiration", url)
	}
}
