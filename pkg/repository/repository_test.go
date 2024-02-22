package repository

import (
	//"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/models"
	"errors"
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

// func TestUserSignUp(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		user   models.UserDetails
// 		stubs  []func(mock sqlmock.Sqlmock)
// 		result models.UserDetailsResponse
// 		err    error
// 	}{
// 		{
// 			name: "successful signup",
// 			user: models.UserDetails{
// 				Name:     "Ananya",
// 				Email:    "ananya@gmail.com",
// 				Password: "password123",
// 				Phone:    "1234567890",
// 			},
// 			stubs: []func(mock sqlmock.Sqlmock){
// 				func(mock sqlmock.Sqlmock) {
// 					mock.ExpectBegin()
// 					mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users")).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "Ananya", "ananya@gmail.com", "1234567890"))
// 					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO wallets")).WillReturnResult(sqlmock.NewResult(1, 1))
// 					mock.ExpectCommit()
// 				},
// 			},
// 			result: models.UserDetailsResponse{
// 				Name:  "Ananya",
// 				Email: "ananya@gmail.com",
// 				Phone: "1234567890",
// 			},
// 			err: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		mockDB, mock, _ := sqlmock.New()

// 		DB, _ := gorm.Open(postgres.New(postgres.Config{
// 			Conn: mockDB,
// 		}), &gorm.Config{})

// 		userRepository := NewUserRepository(DB)

// 		for _, stub := range tt.stubs {
// 			stub(mock)
// 		}

// 		result, err := userRepository.UserSignUp(tt.user)
// 		assert.Equal(t, tt.err, err)
// 		assert.Equal(t, tt.result, result)
// 	}
// }

func TestFindUserByEmail(t *testing.T) {
	tests := []struct {
		name        string
		userLogin   models.UserLogin
		stub        func(mock sqlmock.Sqlmock)
		expected    models.UserSignInResponse
		expectedErr error
	}{
		{
			name: "user does not exist",
			userLogin: models.UserLogin{
				Email:    "ananya@gmail.com",
				Password: "password123",
			},
			stub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"})
				mock.ExpectQuery("SELECT \\* FROM users where email = 'ananya@gmail.com' and blocked = false").WillReturnRows(rows)

				mock.ExpectQuery("SELECT \\* FROM users where email = \\? and blocked = false").WithArgs("ananya@gmail.com").WillReturnError(errors.New("database error"))
			},
			expected:    models.UserSignInResponse{},
			expectedErr: errors.New("error checking user details"),
		},
	}
	for _, tt := range tests {
		mockDB, mock, _ := sqlmock.New()
		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})
		c := userDataBase{DB}

		tt.stub(mock)

		result, err := c.FindUserByEmail(tt.userLogin)

		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, tt.expected, result)
	}
}

func TestUserBlockStatus(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		stub        func(mock sqlmock.Sqlmock)
		expected    bool
		expectedErr error
	}{
		{
			name:  "user not blocked",
			email: "anu@example.com",
			stub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"blocked"}).AddRow(false)
				mock.ExpectQuery("select blocked from users where email = ?").WithArgs("anu@example.com").WillReturnRows(rows)
			},
			expected:    false,
			expectedErr: nil,
		},
		{
			name:  "user blocked",
			email: "anu@example.com",
			stub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"blocked"}).AddRow(true)
				mock.ExpectQuery("select blocked from users where email = ?").WithArgs("anu@example.com").WillReturnRows(rows)
			},
			expected:    true,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		mockDB, mock, _ := sqlmock.New()
		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})
		cr := userDataBase{DB}

		tt.stub(mock)

		result, err := cr.UserBlockStatus(tt.email)

		assert.Equal(t, tt.expectedErr, err)
		assert.Equal(t, tt.expected, result)
	}
}
