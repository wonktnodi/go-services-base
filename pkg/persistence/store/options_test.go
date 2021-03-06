package store

import (
  "testing"
  "time"
  
  "github.com/stretchr/testify/assert"
)

func TestOptionsCostValue(t *testing.T) {
  // Given
  options := Options{
    Cost: 7,
  }
  
  // When - Then
  assert.Equal(t, int64(7), options.CostValue())
}

func TestOptionsExpirationValue(t *testing.T) {
  // Given
  options := Options{
    Expiration: 25 * time.Second,
  }
  
  // When - Then
  assert.Equal(t, 25*time.Second, options.ExpirationValue())
}

func TestOptionsTagsValue(t *testing.T) {
  // Given
  options := Options{
    Tags: []string{"tag1", "tag2", "tag3"},
  }
  
  // When - Then
  assert.Equal(t, []string{"tag1", "tag2", "tag3"}, options.TagsValue())
}
