package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func mockDBRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dbService := MyDBService{DB: mockDB}

	testTable := []struct {
		names       []string
		errExpected error
	}{
		{names: []string{"Ivan", "Vychek"}},
		{names: nil, errExpected: errors.New("ExpectedError")},
	}

	for i, test := range testTable {
		mock.
			ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDBRows(test.names)).
			WillReturnError(test.errExpected)

		names, err := dbService.GetNames()
		if test.errExpected != nil {
			require.ErrorIs(t, err, test.errExpected, "row: %d, expected error:%w, actual error: %w", i, test.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, test.names, names, "row: %d, expected names: %s, actual names: %s", i, test.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type args struct {
		columnName string
		tableName  string
	}

	testTable := []struct {
		dbValues       []string
		valuesExpected []string
		args           args
		errExpected    error
	}{
		{
			valuesExpected: []string{"Ivan", "Vychek", "Grisha"},
			args:           args{"names", "users"},
		},
		{
			dbValues:    nil,
			errExpected: errors.New("ExpectedError"),
			args:        args{"names", "users"},
		},
	}

	dbService := MyDBService{DB: mockDB}

	for i, test := range testTable {
		mock.
			ExpectQuery("SELECT DISTINCT " + test.args.columnName + " FROM " + test.args.tableName).
			WillReturnRows(mockDBRows(test.valuesExpected)).
			WillReturnError(test.errExpected)

		names, err := dbService.SelectUniqueValues(test.args.columnName, test.args.tableName)
		if test.errExpected != nil {
			require.ErrorIs(t, err, test.errExpected, "row: %d, expected error:%w, actual error: %w", i, test.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, test.valuesExpected, names, "row: %d, expected value: %s, actual value: %s", i, test.valuesExpected, names)
	}
}

func TestNew(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type args struct {
		db Database
	}
	tests := []struct {
		name     string
		args     args
		expected MyDBService
	}{
		{
			name:     "test with not empty service",
			args:     args{mockDB},
			expected: MyDBService{mockDB},
		},
		{
			name:     "test with empty service",
			args:     args{nil},
			expected: MyDBService{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("New() = %v, want %v", got, tt.expected)
			}
		})
	}
}
