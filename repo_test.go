package main

import (
	"go-post-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepo_Save(t *testing.T) {
	tests := []struct {
		name           string
		args           model.User
		wantUsers      []model.User
		wantUsersCount int
	}{
		{
			name: "successfully add user :POS",
			args: model.User{
				Name:   "User1",
				PAN:    "ABCDE3333F",
				Mobile: "1234432112",
				Email:  "user1@yahoo.com",
			},
			wantUsers: []model.User{
				{
					Name:   "User1",
					PAN:    "ABCDE3333F",
					Mobile: "1234432112",
					Email:  "user1@yahoo.com",
				},
			},
			wantUsersCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()

			err := repo.CreateUser(tt.args)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantUsersCount, len(repo.users), "expect users count to match")
			assert.Equal(t, tt.wantUsers, repo.users, "expect users to match")
		})
	}
}
