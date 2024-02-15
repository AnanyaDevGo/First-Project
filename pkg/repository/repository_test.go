package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCheckUserAvailability(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		stub func(mock sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "successful, user available",
			arg:  "ananya2@gmail.com",
			stub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("select count(*) from users where email='ananya2@gmail.com'")).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			want: true,
		},
	}
	for _, tt := range tests {
		mockDB, mockSql, _ := sqlmock.New()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		userRepository := NewUserRepository(DB)
		tt.stub(mockSql)

		result := userRepository.CheckUserAvailability(tt.arg)
		assert.Equal(t, tt.want, result)
	}
}
// func TestUserSignUp(t *testing.T){
// 	tests := []struct{
// 		name string
// 		args 
// 		stub func(mock sqlmock.Sqlmock)
// 		want 
// 	}
// }