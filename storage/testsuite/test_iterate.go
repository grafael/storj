// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package testsuite

import (
	"context"
	"math/rand"
	"testing"

	"storj.io/storj/storage"
)

var ctx = context.Background() // test context

func testIterate(t *testing.T, store storage.KeyValueStore) {
	items := storage.Items{
		newItem("a", "a", false),
		newItem("b/1", "b/1", false),
		newItem("b/2", "b/2", false),
		newItem("b/3", "b/3", false),
		newItem("c", "c", false),
		newItem("c/", "c/", false),
		newItem("c//", "c//", false),
		newItem("c/1", "c/1", false),
		newItem("g", "g", false),
		newItem("h", "h", false),
	}
	rand.Shuffle(len(items), items.Swap)
	defer cleanupItems(store, items)
	if err := storage.PutAll(ctx, store, items...); err != nil {
		t.Fatalf("failed to setup: %v", err)
	}

	testIterations(t, store, []iterationTest{
		{"no limits",
			storage.IterateOptions{}, storage.Items{
				newItem("a", "a", false),
				newItem("b/", "", true),
				newItem("c", "c", false),
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},

		{"at a",
			storage.IterateOptions{
				First: storage.Key("a"),
			}, storage.Items{
				newItem("a", "a", false),
				newItem("b/", "", true),
				newItem("c", "c", false),
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},

		{"after a",
			storage.IterateOptions{
				First: storage.NextKey(storage.Key("a")),
			}, storage.Items{
				newItem("b/", "", true),
				newItem("c", "c", false),
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"at b",
			storage.IterateOptions{
				First: storage.Key("b"),
			}, storage.Items{
				newItem("b/", "", true),
				newItem("c", "c", false),
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"after b",
			storage.IterateOptions{
				First: storage.NextKey(storage.Key("b")),
			}, storage.Items{
				newItem("b/", "", true),
				newItem("c", "c", false),
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"after c",
			storage.IterateOptions{
				First: storage.NextKey(storage.Key("c")),
			}, storage.Items{
				newItem("c/", "", true),
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"at e",
			storage.IterateOptions{
				First: storage.Key("e"),
			}, storage.Items{
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"after e",
			storage.IterateOptions{
				First: storage.NextKey(storage.Key("e")),
			}, storage.Items{
				newItem("g", "g", false),
				newItem("h", "h", false),
			}},
		{"prefix b slash",
			storage.IterateOptions{
				Prefix: storage.Key("b/"),
			}, storage.Items{
				newItem("b/1", "b/1", false),
				newItem("b/2", "b/2", false),
				newItem("b/3", "b/3", false),
			}},
		{"prefix c slash",
			storage.IterateOptions{
				Prefix: storage.Key("c/"),
			}, storage.Items{
				newItem("c/", "c/", false),
				newItem("c//", "", true),
				newItem("c/1", "c/1", false),
			}},
		{"prefix c slash slash",
			storage.IterateOptions{
				Prefix: storage.Key("c//"),
			}, storage.Items{
				newItem("c//", "c//", false),
			}},
	})
}
