// +acceptance

package test

import (
	"context"
	v1 "github.com/instinctG/protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(context.Background(), &v1.AddRocketRequest{
			Rocket: &v1.Rocket{
				Id:   "2030",
				Name: "everything",
				Type: "nothing",
			},
		})
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "2030", resp.Rocket.Id)
	})

}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
