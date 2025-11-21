package middleware

import (
	"context"
	"fmt"
	"hifi/types"

	"github.com/valkey-io/valkey-go"
)

func NewRouter(valkeyAddr string) (*types.Router, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{valkeyAddr},
	})
	if err != nil {
		return nil, err
	}

	return &types.Router{
		Valkey: client,
		Mem:    make(map[string]string),
	}, nil
}

// ------------------------- CLOUD -------------------------

func sendToCloud(action, key, value string) {
	fmt.Printf("[CLOUD] action=%s key=%s value=%s\n", action, key, value)
}

// ------------------------- GET -------------------------

func Get(r *types.Router, ctx context.Context, key string) (string, error) {
	if key == "cloud" {
		v, ok := r.Mem[key]
		if !ok {
			return "", fmt.Errorf("cloud key missing")
		}
		go sendToCloud("get", key, v)
		return v, nil
	}

	cmd := r.Valkey.B().Get().Key(key).Build()
	return r.Valkey.Do(ctx, cmd).ToString()
}

// ----------------------- SET -----------------------

func Set(r *types.Router, ctx context.Context, key, val string) error {
	if key == "cloud" {
		r.Mem[key] = val
		go sendToCloud("set", key, val)
		return nil
	}

	cmd := r.Valkey.B().Set().Key(key).Value(val).Build()
	return r.Valkey.Do(ctx, cmd).Error()
}

// ----------------------- DEL -----------------------

func Del(r *types.Router, ctx context.Context, key string) error {
	if key == "cloud" {
		delete(r.Mem, key)
		go sendToCloud("del", key, "")
		return nil
	}

	cmd := r.Valkey.B().Del().Key(key).Build()
	return r.Valkey.Do(ctx, cmd).Error()
}
