package envconf_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/jackc/envconf"
)

func TestConfigItems(t *testing.T) {
	config := envconf.New()

	fooItem := envconf.Item{
		Name:        "FOO",
		Default:     "",
		Description: "the foo",
	}
	config.Register(fooItem)

	barItem := envconf.Item{
		Name:        "BAR",
		Default:     "",
		Description: "the bar",
	}
	config.Register(barItem)

	items := config.Items()
	if len(items) != 2 {
		t.Fatalf("expected 2 item, got %d", len(items))
	}

	expectedItems := []envconf.Item{barItem, fooItem}

	slices.SortFunc(items, func(a, b envconf.Item) int { return cmp.Compare(a.Name, b.Name) })

	for i := range items {
		if items[i] != expectedItems[i] {
			t.Fatalf("expected item %d to be %v, got %v", i, expectedItems[i], items[i])
		}
	}
}

func TestConfigItem(t *testing.T) {
	config := envconf.New()

	fooItem := envconf.Item{
		Name:        "FOO",
		Default:     "",
		Description: "the foo",
	}
	config.Register(fooItem)

	barItem := envconf.Item{
		Name:        "BAR",
		Default:     "",
		Description: "the bar",
	}
	config.Register(barItem)

	item, found := config.Item("BAR")
	if !found {
		t.Fatalf("expected item to be found")
	}
	if item != barItem {
		t.Fatalf("expected item to be %v, got %v", barItem, item)
	}

	_, found = config.Item("BAZ")
	if found {
		t.Fatalf("expected item to not be found")
	}
}

func TestConfigValue(t *testing.T) {
	config := envconf.New()

	fooItem := envconf.Item{
		Name:        "FOO",
		Default:     "default-foo",
		Description: "the foo",
	}
	config.Register(fooItem)

	barItem := envconf.Item{
		Name:        "BAR",
		Default:     "default-bar",
		Description: "the bar",
	}
	config.Register(barItem)

	config.LookupEnvFunc = func(key string) (string, bool) {
		return "env-value", true
	}

	value := config.Value("FOO")
	if value != "env-value" {
		t.Fatalf("expected value to be 'env-value', got %s", value)
	}

	value = config.Value("BAR")
	if value != "env-value" {
		t.Fatalf("expected value to be 'env-value', got %s", value)
	}
}

func TestConfigValueDefault(t *testing.T) {
	config := envconf.New()

	fooItem := envconf.Item{
		Name:        "FOO",
		Default:     "default-foo",
		Description: "the foo",
	}
	config.Register(fooItem)

	barItem := envconf.Item{
		Name:        "BAR",
		Default:     "default-bar",
		Description: "the bar",
	}
	config.Register(barItem)

	config.LookupEnvFunc = func(key string) (string, bool) {
		return "", false
	}

	value := config.Value("FOO")
	if value != "default-foo" {
		t.Fatalf("expected value to be 'default-foo', got %s", value)
	}

	value = config.Value("BAR")
	if value != "default-bar" {
		t.Fatalf("expected value to be 'default-bar', got %s", value)
	}
}

func Example() {
	config := envconf.New()

	// Register configuration items.
	config.Register(envconf.Item{
		Name:        "FOO",
		Default:     "default-foo",
		Description: "the foo",
	})

	config.Register(envconf.Item{
		Name:        "BAR",
		Default:     "default-bar",
		Description: "the bar",
	})

	// You can override the source of environment variables by setting LookupEnvFunc. Defaults to os.LookupEnv.
	config.LookupEnvFunc = func(key string) (string, bool) {
		if key == "FOO" {
			return "foo from environment", true
		}
		return "", false
	}

	// You can use a config to create help output.
	fmt.Println("Configuration items:")
	items := config.Items()
	slices.SortFunc(items, func(a, b envconf.Item) int { return cmp.Compare(a.Name, b.Name) })
	for _, item := range items {
		fmt.Printf("%s: %s (default: %s)\n", item.Name, item.Description, item.Default)
	}

	fmt.Println()

	value := config.Value("FOO")
	fmt.Println("FOO", value)

	value = config.Value("BAR")
	fmt.Println("BAR", value)

	// Output:
	// Configuration items:
	// BAR: the bar (default: default-bar)
	// FOO: the foo (default: default-foo)
	//
	// FOO foo from environment
	// BAR default-bar
}
