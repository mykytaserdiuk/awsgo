package models_test

import (
	"testing"

	"github.com/mykytaserdiuk/aws-go/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestValidateSuccess(t *testing.T) {
	request := models.TodoIn{Topic: "Topic1", Description: "Descr1"}

	err := request.Validate()

	require.NoError(t, err)
	// NO change inside
	require.Equal(t, request.Topic, "Topic1")
	require.Equal(t, request.Description, "Descr1")
}

func TestValidateFailed(t *testing.T) {

	cases := []struct {
		name    string
		in      models.TodoIn
		expired error
	}{
		{
			name:    "bad_topic",
			in:      models.TodoIn{Description: ""},
			expired: models.ErrorUnvalidTopic,
		},
	}
	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()

			require.EqualError(t, err, tc.expired.Error())
		})
	}
}
