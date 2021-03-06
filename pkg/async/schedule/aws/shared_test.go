package aws

import (
	"testing"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/magiconair/properties/assert"
)

func TestHashIdentifier(t *testing.T) {
	identifier := admin.NamedEntityIdentifier{
		Project: "project",
		Domain:  "domain",
		Name:    "name",
	}
	hashedValue := hashIdentifier(identifier)
	assert.Equal(t, uint64(16301494360130577061), hashedValue)
}
