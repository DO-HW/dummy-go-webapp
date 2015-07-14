package sayhello

import "testing"
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T){
  assert.Equal(t, HelloWorld(), "World", "name should be World");
}
