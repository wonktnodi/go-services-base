package metrics

import "go-services-base/pkg/persistence/codec"

// MetricsInterface represents the metrics interface for all available providers
type Interface interface {
  Record(store, metric string, value float64)
  RecordFromCodec(codec codec.Interface)
}
