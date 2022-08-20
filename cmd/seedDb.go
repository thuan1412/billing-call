package main

import (
	"calling-bill/ent"
	"context"
	"fmt"
)

func SeedDb(client *ent.Client) {
	for i := 0; i < 10; i++ {
		ctx := context.Background()
		client.User.Create().SetUsername(fmt.Sprintf("user_%d", i)).SaveX(ctx)
	}
}
