package main

import (
	"github.com/yuryanliang/go_playground/mock_best/mocks"
	"testing"
)

var databaseMock = new(mocks.DB)

func TestSendGreet(t *testing.T) {
	type args struct {
		database DB
		lang     string
	}

	argDatabaseMock := args{
		database: databaseMock,
		lang:     "English",
	}
	//databaseMock.On("FetchMessage","English").Return("mock return anything",nil)
	//databaseMock.On("FetchDefaultMessage",).Return("mock default return",nil)
	//databaseMock.On("FetchMessage",mock.Anything).Return("lah",nil)
	//databaseMock.On("FetchMessage", mock.MatchedBy(func(lang string) bool { return lang[0] == 'i' })).Return("bzzzz", nil) // all of these call FetchMessage("iii"), FetchMessage("i"), FetchMessage("in") will match
	databaseMock.On("FetchMessage", "English").Return("lah", nil)
	databaseMock.On("FetchDefaultMessage").Return("me", nil)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: argDatabaseMock,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendGreet(tt.args.database, tt.args.lang)
			databaseMock.AssertExpectations(t)

		})
	}
}
