package client

import (
	"encoding/json"

	"github.com/Axway/agent-sdk/pkg/apic/auth"
	"github.com/Axway/agent-sdk/pkg/cache"

	"github.com/sirupsen/logrus"

	wm "github.com/Axway/agent-sdk/pkg/watchmanager"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

// WatchClient - stream client for connecting to the Watch Controller
type WatchClient struct {
	config *Config
	logger logrus.FieldLogger
	wm     wm.Manager
}

type sequenceManager struct {
	seqCache cache.Cache
}

func (s *sequenceManager) GetSequence() int64 {
	cachedSeqID, err := s.seqCache.Get("watchSequenceID")
	if err == nil {
		if seqID, ok := cachedSeqID.(float64); ok {
			return int64(seqID)
		}
	}
	return 0
}

// Todo - To be updated after cache persistence story
func getSequenceManager() *sequenceManager {
	seqCache := cache.New()
	err := seqCache.Load("sample.sequence")
	if err != nil {
		seqCache.Set("watchSequenceID", int64(0))
		seqCache.Save("sample.sequence")
	}

	return &sequenceManager{seqCache: seqCache}
}

// NewWatchClient creates a WatchClient
func NewWatchClient(config *Config, logger logrus.FieldLogger) (*WatchClient, error) {
	entry := logger.WithField("package", "client")

	var watchOptions []wm.Option
	watchOptions = append(watchOptions, wm.WithLogger(entry))
	if config.Insecure {
		watchOptions = append(watchOptions, wm.WithTLSConfig(nil))
	}
	watchOptions = append(watchOptions, wm.WithSyncEvents(getSequenceManager()))

	ta := auth.NewTokenAuth(config.Auth, config.TenantID)

	cfg := &wm.Config{
		Host:        config.Host,
		Port:        config.Port,
		TenantID:    config.TenantID,
		TokenGetter: ta.GetToken,
	}

	w, err := wm.New(cfg, watchOptions...)
	if err != nil {
		return nil, err
	}

	return &WatchClient{
		config: config,
		logger: entry,
		wm:     w,
	}, nil
}

// Watch starts a two-way stream with the Watch Controller
func (w WatchClient) Watch() {
	log := w.logger
	log.Info("starting to watch events")

	eventChannel, errCh := make(chan *proto.Event), make(chan error)
	subscriptionID, err := w.wm.RegisterWatch(w.config.TopicSelfLink, eventChannel, errCh)
	if err != nil {
		log.Error(err)
		return
	}

	log = log.WithField("subscriptionId", subscriptionID)
	log.Infof("watch registered successfully")

	for {
		select {
		case err = <-errCh:
			log.Error(err)
			w.wm.CloseWatch(subscriptionID)
			return
		case event := <-eventChannel:
			bts, _ := json.MarshalIndent(event, "", "  ")
			log.Info(string(bts))
		}
	}
}