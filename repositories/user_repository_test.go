package repositories

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-flow/template-api/db"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// UserRepositorySuite represents test suite for UserRepository interface
type UserRepositorySuite struct {
	suite.Suite
	DB   db.Store
	mock sqlmock.Sqlmock

	repository UserRepository
}

//TestUserRepositorySuite is ure repository test suite runner
func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

// SetupSuite configures suite for unit testing
func (s *UserRepositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	//s.DB.LogMode(true)

	s.repository = &userRepository{
		Store: s.DB,
	}
}

// AfterTest ensures that all Test Suite expectations were met
func (s *UserRepositorySuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepositorySuite) Test_GetByID() {

	insertTime := time.Now()
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "profile_image", "email", "is_email_verified", "bio", "phone_number", "is_phone_verified", "country", "state", "area", "city", "address", "post_code", "is_active", "birth_date", "tos_accepted", "invited_by_user_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Sedin", "Dugum", "/path/to/image", "sedin@mop.ba", true, "bio", "+123", true, "BA", "FBiH", "Sarajevo Canton", "Sarajevo", "Zmaja od Bosne 7a", "71000", true, nil, true, nil, insertTime, insertTime, nil)

	query := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((id = ?)) ORDER BY `users`.`id` ASC LIMIT 1"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows).WithArgs(1)

	user, err := s.repository.GetByID(1)

	if err != nil {
		s.Errorf(err, "ups got error")
		return
	}

	assert.NotNil(s.T(), user, "Expected model to be an object, got %v", user)
	assert.Equal(s.T(), uint64(1), user.ID, "Expected user.ID to be 1, got `%v`", user.ID)
	assert.Equal(s.T(), "Sedin", user.FirstName, "Expected user.FirstName to be `Sedin`, got `%v`", user.FirstName)
	assert.Equal(s.T(), "Dugum", user.LastName, "Expected user.LastName to be `Dugum`, got `%v`", user.LastName)
	assert.Equal(s.T(), "/path/to/image", user.ProfileImage, "Expected user.ProfileImage to be `/path/to/image`, got `%v`", user.ProfileImage)
	assert.Equal(s.T(), "sedin@mop.ba", user.Email, "Expected user.Email to be `sedin@mop.ba`, got `%v`", user.Email)
	assert.Equal(s.T(), true, user.IsEmailVerified, "Expected user.IsEmailVerified to be `true`, got `%v`", user.IsEmailVerified)
	assert.Equal(s.T(), "bio", user.Bio, "Expected user.Bio to be `bio`, got `%v`", user.Bio)
	assert.Equal(s.T(), "+123", user.PhoneNumber, "Expected user.PhoneNumber to be `+123`, got `%v`", user.PhoneNumber)
	assert.Equal(s.T(), true, user.IsPhoneVerified, "Expected user.IsPhoneVerified to be `true`, got `%v`", user.IsPhoneVerified)
	assert.Equal(s.T(), "BA", user.Country, "Expected user.Country to be `BA`, got `%v`", user.Country)
	assert.Equal(s.T(), "FBiH", user.State, "Expected user.State to be `FBiH`, got `%v`", user.State)
	assert.Equal(s.T(), "Sarajevo Canton", user.Area, "Expected user.Area to be `Sarajevo Canton`, got `%v`", user.Area)
	assert.Equal(s.T(), "Sarajevo", user.City, "Expected user.City to be `Sarajevo`, got `%v`", user.City)
	assert.Equal(s.T(), "Zmaja od Bosne 7a", user.Address, "Expected user.Address to be `Zmaja od Bosne 7a`, got `%v`", user.Address)
	assert.Equal(s.T(), "71000", user.PostCode, "Expected user.State to be `71000`, got `%v`", user.PostCode)
	assert.Equal(s.T(), true, user.IsActive, "Expected user.IsActive to be `true`, got `%v`", user.IsActive)
	assert.Nil(s.T(), user.BirthDate, "Expected user.BirthDate to be `nil` got `%v`", user.BirthDate)
	assert.Equal(s.T(), true, user.TosAccepted, "Expected user.TosAccepted to be `true` got `%v`", user.TosAccepted)
	assert.Nil(s.T(), user.InvitedByUserID, "Expected user.InvitedByUserID to be `nil` got `%v`", user.InvitedByUserID)
	assert.Equal(s.T(), user.CreatedAt, user.UpdatedAt, "Expected user.CreatedAt to be equal to user.UpdatedAt, got `%v`!= `%v`", user.CreatedAt, user.UpdatedAt)
	assert.Nil(s.T(), user.DeletedAt, "Expected user.DeletedAt to be `nil` got %v", user.DeletedAt)
}

func (s *UserRepositorySuite) Test_GetByEmail() {

	insertTime := time.Now()
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "profile_image", "email", "is_email_verified", "bio", "phone_number", "is_phone_verified", "country", "state", "area", "city", "address", "post_code", "is_active", "birth_date", "tos_accepted", "invited_by_user_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Sedin", "Dugum", "/path/to/image", "sedin@mop.ba", true, "bio", "+123", true, "BA", "FBiH", "Sarajevo Canton", "Sarajevo", "Zmaja od Bosne 7a", "71000", true, nil, true, nil, insertTime, insertTime, nil)

	query := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?)) ORDER BY `users`.`id` ASC LIMIT 1"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows).WithArgs("sedin@mop.ba")

	user, err := s.repository.GetByEmail("sedin@mop.ba")

	if err != nil {
		s.Errorf(err, "ups got error")
		return
	}

	assert.NotNil(s.T(), user, "Expected model to be an object, got %v", user)
	assert.Equal(s.T(), uint64(1), user.ID, "Expected user.ID to be 1, got `%v`", user.ID)
	assert.Equal(s.T(), "Sedin", user.FirstName, "Expected user.FirstName to be `Sedin`, got `%v`", user.FirstName)
	assert.Equal(s.T(), "Dugum", user.LastName, "Expected user.LastName to be `Dugum`, got `%v`", user.LastName)
	assert.Equal(s.T(), "/path/to/image", user.ProfileImage, "Expected user.ProfileImage to be `/path/to/image`, got `%v`", user.ProfileImage)
	assert.Equal(s.T(), "sedin@mop.ba", user.Email, "Expected user.Email to be `sedin@mop.ba`, got `%v`", user.Email)
	assert.Equal(s.T(), true, user.IsEmailVerified, "Expected user.IsEmailVerified to be `true`, got `%v`", user.IsEmailVerified)
	assert.Equal(s.T(), "bio", user.Bio, "Expected user.Bio to be `bio`, got `%v`", user.Bio)
	assert.Equal(s.T(), "+123", user.PhoneNumber, "Expected user.PhoneNumber to be `+123`, got `%v`", user.PhoneNumber)
	assert.Equal(s.T(), true, user.IsPhoneVerified, "Expected user.IsPhoneVerified to be `true`, got `%v`", user.IsPhoneVerified)
	assert.Equal(s.T(), "BA", user.Country, "Expected user.Country to be `BA`, got `%v`", user.Country)
	assert.Equal(s.T(), "FBiH", user.State, "Expected user.State to be `FBiH`, got `%v`", user.State)
	assert.Equal(s.T(), "Sarajevo Canton", user.Area, "Expected user.Area to be `Sarajevo Canton`, got `%v`", user.Area)
	assert.Equal(s.T(), "Sarajevo", user.City, "Expected user.City to be `Sarajevo`, got `%v`", user.City)
	assert.Equal(s.T(), "Zmaja od Bosne 7a", user.Address, "Expected user.Address to be `Zmaja od Bosne 7a`, got `%v`", user.Address)
	assert.Equal(s.T(), "71000", user.PostCode, "Expected user.State to be `71000`, got `%v`", user.PostCode)
	assert.Equal(s.T(), true, user.IsActive, "Expected user.IsActive to be `true`, got `%v`", user.IsActive)
	assert.Nil(s.T(), user.BirthDate, "Expected user.BirthDate to be `nil` got `%v`", user.BirthDate)
	assert.Equal(s.T(), true, user.TosAccepted, "Expected user.TosAccepted to be `true` got `%v`", user.TosAccepted)
	assert.Nil(s.T(), user.InvitedByUserID, "Expected user.InvitedByUserID to be `nil` got `%v`", user.InvitedByUserID)
	assert.Equal(s.T(), user.CreatedAt, user.UpdatedAt, "Expected user.CreatedAt to be equal to user.UpdatedAt, got `%v`!= `%v`", user.CreatedAt, user.UpdatedAt)
	assert.Nil(s.T(), user.DeletedAt, "Expected user.DeletedAt to be `nil` got %v", user.DeletedAt)
}

func (s *UserRepositorySuite) Test_GetAll() {

	insertTime := time.Now()
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "profile_image", "email", "is_email_verified", "bio", "phone_number", "is_phone_verified", "country", "state", "area", "city", "address", "post_code", "is_active", "birth_date", "tos_accepted", "invited_by_user_id", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Sedin", "Dugum", "/path/to/image", "sedin@mop.ba", true, "bio", "+123", true, "BA", "FBiH", "Sarajevo Canton", "Sarajevo", "Zmaja od Bosne 7a", "71000", true, nil, true, nil, insertTime, insertTime, nil).
		AddRow(2, "Sedin", "Dugum", "/path/to/image", "sedin+1@mop.ba", true, "bio", "+123", true, "BA", "FBiH", "Sarajevo Canton", "Sarajevo", "Zmaja od Bosne 7a", "71000", true, nil, true, nil, insertTime, insertTime, nil).
		AddRow(3, "Sedin", "Dugum", "/path/to/image", "sedin+2@mop.ba", true, "bio", "+123", true, "BA", "FBiH", "Sarajevo Canton", "Sarajevo", "Zmaja od Bosne 7a", "71000", true, nil, true, nil, insertTime, insertTime, nil)

	// mock Select query
	query := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY id ASC LIMIT 20 OFFSET 0"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	// prepare paging query result row
	countRow := sqlmock.NewRows([]string{"total_entries_size"}).
		AddRow(3)

	// mock count query
	query = "SELECT count(*) FROM `users` WHERE `users`.`deleted_at` IS NULL"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(countRow)

	paginator := paging.NewWithDefaults()
	users, err := s.repository.GetAll(paginator)
	if err != nil {
		s.Errorf(err, "Unable to fetch users")
		return
	}

	// assert paginator
	assert.Equal(s.T(), 3, paginator.CurrentEntriesSize, "Expected apginator.CurrentEntriesSize to be 3, got %v", paginator.CurrentEntriesSize)
	assert.Equal(s.T(), 3, paginator.TotalEntriesSize, "Expected apginator.TotalEntriesSize to be 3, got %v", paginator.TotalEntriesSize)

	user := users[0]

	assert.NotNil(s.T(), user, "Expected model to be an object, got %v", user)
	assert.Equal(s.T(), uint64(1), user.ID, "Expected user.ID to be 1, got `%v`", user.ID)
	assert.Equal(s.T(), "Sedin", user.FirstName, "Expected user.FirstName to be `Sedin`, got `%v`", user.FirstName)
	assert.Equal(s.T(), "Dugum", user.LastName, "Expected user.LastName to be `Dugum`, got `%v`", user.LastName)
	assert.Equal(s.T(), "/path/to/image", user.ProfileImage, "Expected user.ProfileImage to be `/path/to/image`, got `%v`", user.ProfileImage)
	assert.Equal(s.T(), "sedin@mop.ba", user.Email, "Expected user.Email to be `sedin@mop.ba`, got `%v`", user.Email)
	assert.Equal(s.T(), true, user.IsEmailVerified, "Expected user.IsEmailVerified to be `true`, got `%v`", user.IsEmailVerified)
	assert.Equal(s.T(), "bio", user.Bio, "Expected user.Bio to be `bio`, got `%v`", user.Bio)
	assert.Equal(s.T(), "+123", user.PhoneNumber, "Expected user.PhoneNumber to be `+123`, got `%v`", user.PhoneNumber)
	assert.Equal(s.T(), true, user.IsPhoneVerified, "Expected user.IsPhoneVerified to be `true`, got `%v`", user.IsPhoneVerified)
	assert.Equal(s.T(), "BA", user.Country, "Expected user.Country to be `BA`, got `%v`", user.Country)
	assert.Equal(s.T(), "FBiH", user.State, "Expected user.State to be `FBiH`, got `%v`", user.State)
	assert.Equal(s.T(), "Sarajevo Canton", user.Area, "Expected user.Area to be `Sarajevo Canton`, got `%v`", user.Area)
	assert.Equal(s.T(), "Sarajevo", user.City, "Expected user.City to be `Sarajevo`, got `%v`", user.City)
	assert.Equal(s.T(), "Zmaja od Bosne 7a", user.Address, "Expected user.Address to be `Zmaja od Bosne 7a`, got `%v`", user.Address)
	assert.Equal(s.T(), "71000", user.PostCode, "Expected user.State to be `71000`, got `%v`", user.PostCode)
	assert.Equal(s.T(), true, user.IsActive, "Expected user.IsActive to be `true`, got `%v`", user.IsActive)
	assert.Nil(s.T(), user.BirthDate, "Expected user.BirthDate to be `nil` got `%v`", user.BirthDate)
	assert.Equal(s.T(), true, user.TosAccepted, "Expected user.TosAccepted to be `true` got `%v`", user.TosAccepted)
	assert.Nil(s.T(), user.InvitedByUserID, "Expected user.InvitedByUserID to be `nil` got `%v`", user.InvitedByUserID)
	assert.Equal(s.T(), user.CreatedAt, user.UpdatedAt, "Expected user.CreatedAt to be equal to user.UpdatedAt, got `%v`!= `%v`", user.CreatedAt, user.UpdatedAt)
	assert.Nil(s.T(), user.DeletedAt, "Expected user.DeletedAt to be `nil` got %v", user.DeletedAt)
}

func (s *UserRepositorySuite) Test_Create() {
	insertTime := time.Now()
	// prepare model
	user := &models.User{
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
		CreatedAt: insertTime,
		UpdatedAt: insertTime,
	}

	// mock query
	s.mock.ExpectBegin()
	query := "INSERT INTO `users` (`first_name`,`last_name`,`profile_image`,`email`,`is_email_verified`,`bio`,`phone_number`,`is_phone_verified`,`country`,`state`,`area`,`city`,`address`,`post_code`,`birth_date`,`tos_accepted`,`invited_by_user_id`,`is_active`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(
			user.FirstName,
			user.LastName,
			user.ProfileImage,
			user.Email,
			user.IsEmailVerified,
			user.Bio,
			user.PhoneNumber,
			user.IsPhoneVerified,
			user.Country,
			user.State,
			user.Area,
			user.City,
			user.Address,
			user.PostCode,
			user.BirthDate,
			user.TosAccepted,
			user.InvitedByUserID,
			user.IsActive,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.Save(user)
	if err != nil {
		s.Errorf(err, "unable to create user")
	}

	assert.NotNil(s.T(), user, "Expected model to be an object, got %v", user)
	assert.Equal(s.T(), uint64(1), user.ID, "Expected user.ID to be 1, got `%v`", user.ID)
	assert.Equal(s.T(), "Sedin", user.FirstName, "Expected user.FirstName to be `Sedin`, got `%v`", user.FirstName)
	assert.Equal(s.T(), "Dugum", user.LastName, "Expected user.LastName to be `Dugum`, got `%v`", user.LastName)
	assert.Equal(s.T(), "", user.ProfileImage, "Expected user.ProfileImage to be ``, got `%v`", user.ProfileImage)
	assert.Equal(s.T(), "sedin@mop.ba", user.Email, "Expected user.Email to be `sedin@mop.ba`, got `%v`", user.Email)
	assert.Equal(s.T(), false, user.IsEmailVerified, "Expected user.IsEmailVerified to be `false`, got `%v`", user.IsEmailVerified)
	assert.Equal(s.T(), "", user.Bio, "Expected user.Bio to be ``, got `%v`", user.Bio)
	assert.Equal(s.T(), "", user.PhoneNumber, "Expected user.PhoneNumber to be ``, got `%v`", user.PhoneNumber)
	assert.Equal(s.T(), false, user.IsPhoneVerified, "Expected user.IsPhoneVerified to be `false`, got `%v`", user.IsPhoneVerified)
	assert.Equal(s.T(), "", user.Country, "Expected user.Country to be ``, got `%v`", user.Country)
	assert.Equal(s.T(), "", user.State, "Expected user.State to be ``, got `%v`", user.State)
	assert.Equal(s.T(), "", user.Area, "Expected user.Area to be ``, got `%v`", user.Area)
	assert.Equal(s.T(), "", user.City, "Expected user.City to be ``, got `%v`", user.City)
	assert.Equal(s.T(), "", user.Address, "Expected user.Address to be ``, got `%v`", user.Address)
	assert.Equal(s.T(), "", user.PostCode, "Expected user.State to be ``, got `%v`", user.PostCode)
	assert.Equal(s.T(), false, user.IsActive, "Expected user.IsActive to be `false`, got `%v`", user.IsActive)
	assert.Nil(s.T(), user.BirthDate, "Expected user.BirthDate to be `nil` got `%v`", user.BirthDate)
	assert.Equal(s.T(), false, user.TosAccepted, "Expected user.TosAccepted to be `false` got `%v`", user.TosAccepted)
	assert.Nil(s.T(), user.InvitedByUserID, "Expected user.InvitedByUserID to be `nil` got `%v`", user.InvitedByUserID)
	assert.Equal(s.T(), user.CreatedAt, user.UpdatedAt, "Expected user.CreatedAt to be equal to user.UpdatedAt, got `%v`!= `%v`", user.CreatedAt, user.UpdatedAt)
	assert.Nil(s.T(), user.DeletedAt, "Expected user.DeletedAt to be `nil` got %v", user.DeletedAt)
}

func (s *UserRepositorySuite) Test_Update() {
	// prepare model
	user := &models.User{
		ID:        uint64(1),
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	// mock query
	s.mock.ExpectBegin()
	query := "UPDATE `users` SET `first_name` = ?, `last_name` = ?, `profile_image` = ?, `email` = ?, `is_email_verified` = ?, `bio` = ?, `phone_number` = ?, `is_phone_verified` = ?, `country` = ?, `state` = ?, `area` = ?, `city` = ?, `address` = ?, `post_code` = ?, `birth_date` = ?, `tos_accepted` = ?, `invited_by_user_id` = ?, `is_active` = ?, `created_at` = ?, `updated_at` = ?, `deleted_at` = ? WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ?"
	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(
			user.FirstName,
			user.LastName,
			user.ProfileImage,
			user.Email,
			user.IsEmailVerified,
			user.Bio,
			user.PhoneNumber,
			user.IsPhoneVerified,
			user.Country,
			user.State,
			user.Area,
			user.City,
			user.Address,
			user.PostCode,
			user.BirthDate,
			user.TosAccepted,
			user.InvitedByUserID,
			user.IsActive,
			AnyTime{},
			AnyTime{},
			user.DeletedAt,
			user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := s.repository.Save(user)
	if err != nil {
		s.Errorf(err, "unable to update user")
	}

	assert.NotNil(s.T(), user, "Expected model to be an object, got %v", user)
	assert.Equal(s.T(), uint64(1), user.ID, "Expected user.ID to be 1, got `%v`", user.ID)
	assert.Equal(s.T(), "Sedin", user.FirstName, "Expected user.FirstName to be `Sedin`, got `%v`", user.FirstName)
	assert.Equal(s.T(), "Dugum", user.LastName, "Expected user.LastName to be `Dugum`, got `%v`", user.LastName)
	assert.Equal(s.T(), "", user.ProfileImage, "Expected user.ProfileImage to be ``, got `%v`", user.ProfileImage)
	assert.Equal(s.T(), "sedin@mop.ba", user.Email, "Expected user.Email to be `sedin@mop.ba`, got `%v`", user.Email)
	assert.Equal(s.T(), false, user.IsEmailVerified, "Expected user.IsEmailVerified to be `false`, got `%v`", user.IsEmailVerified)
	assert.Equal(s.T(), "", user.Bio, "Expected user.Bio to be ``, got `%v`", user.Bio)
	assert.Equal(s.T(), "", user.PhoneNumber, "Expected user.PhoneNumber to be ``, got `%v`", user.PhoneNumber)
	assert.Equal(s.T(), false, user.IsPhoneVerified, "Expected user.IsPhoneVerified to be `false`, got `%v`", user.IsPhoneVerified)
	assert.Equal(s.T(), "", user.Country, "Expected user.Country to be ``, got `%v`", user.Country)
	assert.Equal(s.T(), "", user.State, "Expected user.State to be ``, got `%v`", user.State)
	assert.Equal(s.T(), "", user.Area, "Expected user.Area to be ``, got `%v`", user.Area)
	assert.Equal(s.T(), "", user.City, "Expected user.City to be ``, got `%v`", user.City)
	assert.Equal(s.T(), "", user.Address, "Expected user.Address to be ``, got `%v`", user.Address)
	assert.Equal(s.T(), "", user.PostCode, "Expected user.State to be ``, got `%v`", user.PostCode)
	assert.Equal(s.T(), false, user.IsActive, "Expected user.IsActive to be `false`, got `%v`", user.IsActive)
	assert.Nil(s.T(), user.BirthDate, "Expected user.BirthDate to be `nil` got `%v`", user.BirthDate)
	assert.Equal(s.T(), false, user.TosAccepted, "Expected user.TosAccepted to be `false` got `%v`", user.TosAccepted)
	assert.Nil(s.T(), user.InvitedByUserID, "Expected user.InvitedByUserID to be `nil` got `%v`", user.InvitedByUserID)
	//assert.Equal(s.T(), user.CreatedAt, user.UpdatedAt, "Expected user.CreatedAt to be equal to user.UpdatedAt, got `%v`!= `%v`", user.CreatedAt, user.UpdatedAt)
	assert.Nil(s.T(), user.DeletedAt, "Expected user.DeletedAt to be `nil` got %v", user.DeletedAt)
}

func (s *UserRepositorySuite) Test_Delete() {
	// prepare model
	user := &models.User{
		ID:        uint64(1),
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	// mock query
	s.mock.ExpectBegin()
	query := "DELETE FROM `users` WHERE `users`.`id` = ?"
	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(
			user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := s.repository.Delete(user)
	if err != nil {
		s.Errorf(err, "unable to create user")
	}
}
