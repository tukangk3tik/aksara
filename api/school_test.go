package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	mockdb "github.com/tukangk3tik/aksara/db/mock"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/security"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/mock/gomock"
)

const userID = 1

func TestListSchoolAPI(t *testing.T) {
	schools := []db.ListAllSchoolsRow{}
	for i := 1; i <= 10; i++ {
		schools = append(schools, randomSchool(i))
	}

	testCase := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListAllSchoolsParams{
					Limit:  int32(10),
					Offset: 0,
				}

				store.EXPECT().
					ListAllSchools(gomock.Any(), gomock.Eq(&arg)).
					Times(1).
					Return(schools, nil)

				store.EXPECT().
					TotalListAllSchools(gomock.Any()).
					Times(1).
					Return(int64(len(schools)), nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name: "NoAuthorization",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAllSchools(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					TotalListAllSchools(gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListAllSchoolsParams{
					Limit:  int32(10),
					Offset: 0,
				}

				store.EXPECT().
					ListAllSchools(gomock.Any(), gomock.Eq(&arg)).
					Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/schools", nil)
			require.NoError(t, err)

			tc.setupAuth(t, req, server.tokenMaker)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetSchoolAPI(t *testing.T) {
	school := parseSchoolRowToGetSchoolByIdRowModel(randomSchool(1))

	testCase := []struct {
		name          string
		schoolId      int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker)
	}{
		{
			name:     "OK",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(school, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchSchoolDetail(t, recorder.Body, school)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NotFound",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(db.GetSchoolByIdRow{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "InternalError",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(db.GetSchoolByIdRow{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NoAuthorization",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(school.ID)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/schools/%d", tc.schoolId), nil)
			require.NoError(t, err)

			tc.setupAuth(t, req, server.tokenMaker)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestCreateSchoolAPI(t *testing.T) {
	school := parseSchoolRowToSchoolModel(randomSchool(1))
	postBody := gin.H{
		"code":             school.Code,
		"name":             school.Name,
		"is_public_school": school.IsPublicSchool,
		"office_id":        school.OfficeID.Int64,
		"province_id":      school.ProvinceID,
		"regency_id":       school.RegencyID,
		"district_id":      school.DistrictID,
		"email":            school.Email.String,
		"phone":            school.Phone.String,
		"address":          school.Address.String,
		"logo_url":         school.LogoUrl.String,
	}

	testCase := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker)
	}{
		{
			name: "OK",
			body: postBody,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateSchoolParams{
					Code:           school.Code,
					Name:           school.Name,
					IsPublicSchool: school.IsPublicSchool,
					OfficeID:       school.OfficeID,
					ProvinceID:     school.ProvinceID,
					RegencyID:      school.RegencyID,
					DistrictID:     school.DistrictID,
					Email:          school.Email,
					Phone:          school.Phone,
					Address:        school.Address,
					LogoUrl:        school.LogoUrl,
					CreatedBy:      userID,
				}

				store.EXPECT().
					CreateSchool(gomock.Any(), gomock.Eq(&arg)).
					Times(1).
					Return(school, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchSchool(t, recorder.Body, school)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"name": school.Name,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSchool(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name: "DuplicateCode",
			body: postBody,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSchool(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Schools{}, &pq.Error{Code: db.UniqueViolation, Constraint: "schools_code_key"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name: "InternalError",
			body: postBody,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSchool(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Schools{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name: "NoAuthorization",
			body: postBody,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSchool(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/schools", bytes.NewReader(data))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			tc.setupAuth(t, req, server.tokenMaker)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdateSchoolAPI(t *testing.T) {
	randomSchool := randomSchool(1)
	oldSchool := parseSchoolRowToGetSchoolByIdRowModel(randomSchool)

	updateSchool := parseSchoolRowToSchoolModel(randomSchool)
	updateSchool.Name = "Updated School Name"
	updateSchool.Email = sql.NullString{String: "updated@school.com", Valid: true}
	updateSchool.Phone = sql.NullString{String: "08123456789", Valid: true}
	updateSchool.Address = sql.NullString{String: "Updated Address", Valid: true}
	updateSchool.LogoUrl = sql.NullString{String: "updated_logo.png", Valid: true}
	updateSchool.IsPublicSchool = true

	testCase := []struct {
		name          string
		schoolId      int64
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker)
	}{
		{
			name:     "OK",
			schoolId: oldSchool.ID,
			body: gin.H{
				"name":     "Updated School Name",
				"email":    "updated@school.com",
				"phone":    "08123456789",
				"address":  "Updated Address",
				"logo_url": "updated_logo.png",
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateSchoolParams{
					ID:         oldSchool.ID,
					Name:       "Updated School Name",
					OfficeID:   oldSchool.OfficeID,
					ProvinceID: oldSchool.ProvinceID,
					RegencyID:  oldSchool.RegencyID,
					DistrictID: oldSchool.DistrictID,
					Email:      sql.NullString{String: "updated@school.com", Valid: true},
					Phone:      sql.NullString{String: "08123456789", Valid: true},
					Address:    sql.NullString{String: "Updated Address", Valid: true},
					LogoUrl:    sql.NullString{String: "updated_logo.png", Valid: true},
				}

				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(oldSchool.ID)).
					Times(1).
					Return(oldSchool, nil)

				store.EXPECT().
					UpdateSchool(gomock.Any(), gomock.Eq(&arg)).
					Times(1).
					Return(updateSchool, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				updateSchoolRow := db.GetSchoolByIdRow{
					ID:             updateSchool.ID,
					Code:           updateSchool.Code,
					Name:           updateSchool.Name,
					IsPublicSchool: updateSchool.IsPublicSchool,
					OfficeID:       updateSchool.OfficeID,
					ProvinceID:     updateSchool.ProvinceID,
					RegencyID:      updateSchool.RegencyID,
					DistrictID:     updateSchool.DistrictID,
					Email:          updateSchool.Email,
					Phone:          updateSchool.Phone,
					Address:        updateSchool.Address,
					LogoUrl:        updateSchool.LogoUrl,
					CreatedAt:      updateSchool.CreatedAt,
					UpdatedAt:      updateSchool.UpdatedAt,
					CreatedBy:      updateSchool.CreatedBy,
				}

				requireBodyMatchSchoolDetail(t, recorder.Body, updateSchoolRow)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NotFound",
			schoolId: oldSchool.ID,
			body: gin.H{
				"name":     "Updated School Name",
				"email":    "updated@school.com",
				"phone":    "08123456789",
				"address":  "Updated Address",
				"logo_url": "updated_logo.png",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(oldSchool.ID)).
					Times(1).
					Return(db.GetSchoolByIdRow{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "BadRequest",
			schoolId: oldSchool.ID,
			body: gin.H{
				"email":    "updated@school.com",
				"phone":    "08123456789",
				"address":  "Updated Address",
				"logo_url": "updated_logo.png",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					UpdateSchool(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "InternalError",
			schoolId: oldSchool.ID,
			body: gin.H{
				"name":     "Updated School Name",
				"email":    "updated@school.com",
				"phone":    "08123456789",
				"address":  "Updated Address",
				"logo_url": "updated_logo.png",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Eq(oldSchool.ID)).
					Times(1).
					Return(oldSchool, nil)

				store.EXPECT().
					UpdateSchool(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Schools{}, sql.ErrConnDone)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NoAuthorization",
			schoolId: oldSchool.ID,
			body: gin.H{
				"name":     "Updated School Name",
				"email":    "updated@school.com",
				"phone":    "08123456789",
				"address":  "Updated Address",
				"logo_url": "updated_logo.png",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSchoolById(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					UpdateSchool(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/schools/%d", tc.schoolId)
			req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			tc.setupAuth(t, req, server.tokenMaker)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeleteSchoolAPI(t *testing.T) {
	school := parseSchoolRowToGetSchoolByIdRowModel(randomSchool(1))
	deleteResult := mockSQLResult{}
	deleteNotFoundResult := mockNotFoundSQLResult{}

	testCase := []struct {
		name          string
		schoolId      int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker)
	}{
		{
			name:     "OK",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSchool(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(deleteResult, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NotFound",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSchool(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(deleteNotFoundResult, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "InternalError",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSchool(gomock.Any(), gomock.Eq(school.ID)).
					Times(1).
					Return(deleteResult, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
				addAuthorization(t, request, tokenMaker, security.AuthorizationTypeBearer, userID, time.Minute)
			},
		},
		{
			name:     "NoAuthorization",
			schoolId: school.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSchool(gomock.Any(), gomock.Eq(school.ID)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker security.TokenMaker) {
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/schools/%d", tc.schoolId), nil)
			require.NoError(t, err)

			tc.setupAuth(t, req, server.tokenMaker)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchSchoolDetail(t *testing.T, body *bytes.Buffer, school db.GetSchoolByIdRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotSuccessResponse response.SuccessResponse
	err = json.Unmarshal(data, &gotSuccessResponse)
	require.NoError(t, err)

	require.NotNil(t, gotSuccessResponse.Data, "response data should not be nil")
	dataMap, ok := gotSuccessResponse.Data.(map[string]interface{})
	require.True(t, ok, "response data should be a map")

	gotSchool := parseMapToSchoolModelDetail(dataMap)
	require.Equal(t, school, gotSchool)
}

func requireBodyMatchSchool(t *testing.T, body *bytes.Buffer, school db.Schools) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotSuccessResponse response.SuccessResponse
	err = json.Unmarshal(data, &gotSuccessResponse)
	require.NoError(t, err)

	require.NotNil(t, gotSuccessResponse.Data, "response data should not be nil")
	dataMap, ok := gotSuccessResponse.Data.(map[string]interface{})
	require.True(t, ok, "response data should be a map")

	gotSchool := parseMapToSchoolModel(dataMap)
	require.Equal(t, school, gotSchool)
}

func randomSchool(ID int) db.ListAllSchoolsRow {
	return db.ListAllSchoolsRow{
		ID:             int64(ID),
		Code:           fmt.Sprintf("TESTSCHOOL-%d", ID),
		Name:           utils.RandomString(10),
		IsPublicSchool: true,
		Office:         utils.RandomString(10),
		OfficeID:       sql.NullInt64{Int64: int64(1), Valid: true},
		Province:       "Test Province",
		Regency:        "Test Regency",
		District:       sql.NullString{String: "Test District", Valid: true},
		ProvinceID:     provinceID,
		RegencyID:      regencyID,
		DistrictID:     districtID,
		Email:          sql.NullString{String: fmt.Sprintf("test-%d@aksara.com", ID), Valid: true},
		Phone:          sql.NullString{String: fmt.Sprintf("0812345678%d", ID), Valid: true},
		Address:        sql.NullString{String: "Test Address", Valid: true},
		LogoUrl:        sql.NullString{String: "logo_url", Valid: true},
		CreatedBy:      userID,
	}
}

func parseSchoolRowToGetSchoolByIdRowModel(model db.ListAllSchoolsRow) db.GetSchoolByIdRow {
	return db.GetSchoolByIdRow{
		ID:             int64(model.ID),
		Code:           model.Code,
		Name:           model.Name,
		IsPublicSchool: model.IsPublicSchool,
		OfficeID:       sql.NullInt64{Int64: int64(model.OfficeID.Int64), Valid: true},
		ProvinceID:     int32(model.ProvinceID),
		RegencyID:      int32(model.RegencyID),
		DistrictID:     int32(model.DistrictID),
		Email:          sql.NullString{String: model.Email.String, Valid: true},
		Phone:          sql.NullString{String: model.Phone.String, Valid: true},
		Address:        sql.NullString{String: model.Address.String, Valid: true},
		LogoUrl:        sql.NullString{String: model.LogoUrl.String, Valid: true},
		CreatedBy:      model.CreatedBy,
	}
}

func parseSchoolRowToSchoolModel(model db.ListAllSchoolsRow) db.Schools {
	return db.Schools{
		ID:             int64(model.ID),
		Code:           model.Code,
		Name:           model.Name,
		IsPublicSchool: model.IsPublicSchool,
		OfficeID:       sql.NullInt64{Int64: int64(model.OfficeID.Int64), Valid: true},
		ProvinceID:     int32(model.ProvinceID),
		RegencyID:      int32(model.RegencyID),
		DistrictID:     int32(model.DistrictID),
		Email:          sql.NullString{String: model.Email.String, Valid: true},
		Phone:          sql.NullString{String: model.Phone.String, Valid: true},
		Address:        sql.NullString{String: model.Address.String, Valid: true},
		LogoUrl:        sql.NullString{String: model.LogoUrl.String, Valid: true},
		CreatedBy:      model.CreatedBy,
	}
}

func parseMapToSchoolModelDetail(data map[string]any) db.GetSchoolByIdRow {
	school := db.GetSchoolByIdRow{
		IsPublicSchool: data["is_public_school"].(bool),
	}

	if id, ok := data["id"].(float64); ok {
		school.ID = int64(id)
	}
	if code, ok := data["code"].(string); ok {
		school.Code = code
	}
	if name, ok := data["name"].(string); ok {
		school.Name = name
	}

	fmt.Println("office_id", data["office_id"])
	if officeID, ok := data["office_id"].(float64); ok {
		school.OfficeID = sql.NullInt64{Int64: int64(officeID), Valid: true}
	}
	if provinceID, ok := data["province_id"].(float64); ok {
		school.ProvinceID = int32(provinceID)
	}
	if regencyID, ok := data["regency_id"].(float64); ok {
		school.RegencyID = int32(regencyID)
	}
	if districtID, ok := data["district_id"].(float64); ok {
		school.DistrictID = int32(districtID)
	}
	if email, ok := data["email"].(string); ok {
		school.Email = sql.NullString{String: email, Valid: true}
	}
	if phone, ok := data["phone"].(string); ok {
		school.Phone = sql.NullString{String: phone, Valid: true}
	}
	if address, ok := data["address"].(string); ok {
		school.Address = sql.NullString{String: address, Valid: true}
	}
	if logoURL, ok := data["logo_url"].(string); ok {
		school.LogoUrl = sql.NullString{String: logoURL, Valid: true}
	}
	if createdBy, ok := data["created_by"].(float64); ok {
		school.CreatedBy = int64(createdBy)
	}

	return school
}

func parseMapToSchoolModel(data map[string]any) db.Schools {
	school := db.Schools{
		IsPublicSchool: data["is_public_school"].(bool),
	}

	if id, ok := data["id"].(float64); ok {
		school.ID = int64(id)
	}
	if code, ok := data["code"].(string); ok {
		school.Code = code
	}
	if name, ok := data["name"].(string); ok {
		school.Name = name
	}

	fmt.Println("office_id", data["office_id"])
	if officeID, ok := data["office_id"].(float64); ok {
		school.OfficeID = sql.NullInt64{Int64: int64(officeID), Valid: true}
	}
	if provinceID, ok := data["province_id"].(float64); ok {
		school.ProvinceID = int32(provinceID)
	}
	if regencyID, ok := data["regency_id"].(float64); ok {
		school.RegencyID = int32(regencyID)
	}
	if districtID, ok := data["district_id"].(float64); ok {
		school.DistrictID = int32(districtID)
	}
	if email, ok := data["email"].(string); ok {
		school.Email = sql.NullString{String: email, Valid: true}
	}
	if phone, ok := data["phone"].(string); ok {
		school.Phone = sql.NullString{String: phone, Valid: true}
	}
	if address, ok := data["address"].(string); ok {
		school.Address = sql.NullString{String: address, Valid: true}
	}
	if logoURL, ok := data["logo_url"].(string); ok {
		school.LogoUrl = sql.NullString{String: logoURL, Valid: true}
	}
	if createdBy, ok := data["created_by"].(float64); ok {
		school.CreatedBy = int64(createdBy)
	}

	return school
}
