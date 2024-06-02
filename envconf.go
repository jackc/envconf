package envconf

import (
	"cmp"
	"os"
	"slices"
)

type Item struct {
	Name        string
	Default     string
	Description string
}

// Config handles access to configuration stored in the environment.
type Config struct {
	items map[string]Item

	// LookupEnvFunc is used to fetch value of an environment variable.
	LookupEnvFunc func(string) (string, bool)
}

// New creates a new Config.
func New() *Config {
	return &Config{
		items:         make(map[string]Item),
		LookupEnvFunc: os.LookupEnv,
	}
}

// Register registers a configuration item.
func (c *Config) Register(item Item) {
	c.items[item.Name] = item
}

// Items returns all registered configuration items sorted by name alphabetically.
func (c *Config) Items() []Item {
	items := make([]Item, 0, len(c.items))
	for _, item := range c.items {
		items = append(items, item)
	}

	slices.SortFunc(items, func(a, b Item) int { return cmp.Compare(a.Name, b.Name) })

	return items
}

// Item returns a configuration item by name.
func (c *Config) Item(name string) (Item, bool) {
	item, found := c.items[name]
	return item, found
}

// MustItem returns a configuration item by name. It panics if the item is not found.
func (c *Config) MustItem(name string) Item {
	item, found := c.Item(name)
	if !found {
		panic("item not found: " + name)
	}

	return item
}

// Value returns the value of an environment variable. If the environment variable is not set or is the empty string, it
// returns the default value.
func (c *Config) Value(key string) string {
	if value, found := c.LookupEnvFunc(key); found {
		return value
	}

	if item, found := c.items[key]; found {
		return item.Default
	}

	return ""
}
