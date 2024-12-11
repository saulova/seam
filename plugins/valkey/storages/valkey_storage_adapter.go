package storages

import (
	"context"
	"fmt"
	"time"

	"github.com/saulova/seam/plugins/valkey/configs"

	"github.com/valkey-io/valkey-go"
)

type ValkeyStorageAdapter struct {
	storageConfig *configs.ValkeyStorageConfig
	cacheClient   valkey.Client
}

func NewValkeyStorageAdapter(config interface{}) (*ValkeyStorageAdapter, error) {
	storageConfig, err := configs.NewValkeyStorageConfig(config)
	if err != nil {
		return nil, err
	}

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{fmt.Sprintf("%s:%d", storageConfig.Host, storageConfig.Port)},
	})
	if err != nil {
		return nil, err
	}

	instance := &ValkeyStorageAdapter{
		storageConfig: storageConfig,
		cacheClient:   client,
	}

	return instance, nil
}

func (v *ValkeyStorageAdapter) getSessionId(key string) string {
	return fmt.Sprintf("%s:%s", v.storageConfig.Prefix, key)
}

func (v *ValkeyStorageAdapter) Get(key string) ([]byte, error) {
	sessionID := v.getSessionId(key)

	val, err := v.cacheClient.Do(context.Background(), v.cacheClient.B().Get().Key(sessionID).Build()).AsBytes()
	if valkey.IsValkeyNil(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return val, nil
}

func (v *ValkeyStorageAdapter) Set(key string, val []byte, exp time.Duration) error {
	sessionID := v.getSessionId(key)

	return v.cacheClient.Do(context.Background(), v.cacheClient.B().Set().Key(sessionID).Value(string(val)).Ex(exp).Build()).Error()
}

func (v *ValkeyStorageAdapter) Delete(key string) error {
	return v.cacheClient.Do(context.Background(), v.cacheClient.B().Del().Key(key).Build()).Error()
}

func (v *ValkeyStorageAdapter) Reset() error {
	return v.cacheClient.Do(context.Background(), v.cacheClient.B().Flushall().Build()).Error()
}

func (v *ValkeyStorageAdapter) Close() error {
	v.cacheClient.Close()

	return nil
}
