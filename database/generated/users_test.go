package generated

// user_test.go contains unit tests for the generated code related to the Users table.

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:    "klm123",
		Password:    "1xtdf",
		Email:       "7@7.com",
		AccessLevel: 3,
	}
	var user User
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
