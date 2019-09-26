package executioncluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInClusterGetTarget(t *testing.T) {
	cluster := InCluster{
		target: ExecutionTarget{
			ID: "t1",
		},
	}
	target, err := cluster.GetTarget(nil)
	assert.Nil(t, err)
	assert.Equal(t, "t1", target.ID)
}

func TestInClusterGetRemoteTarget(t *testing.T) {
	cluster := InCluster{
		target: ExecutionTarget{},
	}
	_, err := cluster.GetTarget(&ExecutionTargetSpec{TargetId: "t1"})
	assert.EqualError(t, err, "remote target t1 is not supported")
}

func TestInClusterGetAllValidTargets(t *testing.T) {
	cluster := InCluster{
		target: ExecutionTarget{
			ID: "t1",
		},
	}
	targets := cluster.GetAllValidTargets()
	assert.Equal(t, 1, len(targets))
	assert.Equal(t, "t1", targets[0].ID)

}