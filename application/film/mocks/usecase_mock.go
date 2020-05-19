// Code generated by MockGen. DO NOT EDIT.
// Source: ./application/film/usecase.go

// Package mock_film is a generated GoMock package.
package mock_film

import (
	models "github.com/go-park-mail-ru/2020_1_k-on/application/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetFilmsList mocks base method
func (m *MockUsecase) GetFilmsList(begin, end uint) (models.Films, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsList", begin, end)
	ret0, _ := ret[0].(models.Films)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFilmsList indicates an expected call of GetFilmsList
func (mr *MockUsecaseMockRecorder) GetFilmsList(begin, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsList", reflect.TypeOf((*MockUsecase)(nil).GetFilmsList), begin, end)
}

// GetFilm mocks base method
func (m *MockUsecase) GetFilm(id uint) (models.Film, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilm", id)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFilm indicates an expected call of GetFilm
func (mr *MockUsecaseMockRecorder) GetFilm(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilm", reflect.TypeOf((*MockUsecase)(nil).GetFilm), id)
}

// CreateFilm mocks base method
func (m *MockUsecase) CreateFilm(f models.Film) (models.Film, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilm", f)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CreateFilm indicates an expected call of CreateFilm
func (mr *MockUsecaseMockRecorder) CreateFilm(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilm", reflect.TypeOf((*MockUsecase)(nil).CreateFilm), f)
}

// GetFilmGenres mocks base method
func (m *MockUsecase) GetFilmGenres(fid uint) (models.Genres, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmGenres", fid)
	ret0, _ := ret[0].(models.Genres)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFilmGenres indicates an expected call of GetFilmGenres
func (mr *MockUsecaseMockRecorder) GetFilmGenres(fid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmGenres", reflect.TypeOf((*MockUsecase)(nil).GetFilmGenres), fid)
}

// Search mocks base method
func (m *MockUsecase) Search(word string, query map[string][]string) (models.Films, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", word, query)
	ret0, _ := ret[0].(models.Films)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockUsecaseMockRecorder) Search(word, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUsecase)(nil).Search), word, query)
}

// GetSimilarFilms mocks base method
func (m *MockUsecase) GetSimilarFilms(fid uint) (models.Films, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarFilms", fid)
	ret0, _ := ret[0].(models.Films)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSimilarFilms indicates an expected call of GetSimilarFilms
func (mr *MockUsecaseMockRecorder) GetSimilarFilms(fid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarFilms", reflect.TypeOf((*MockUsecase)(nil).GetSimilarFilms), fid)
}

// GetSimilarSeries mocks base method
func (m *MockUsecase) GetSimilarSeries(fid uint) (models.SeriesArr, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarSeries", fid)
	ret0, _ := ret[0].(models.SeriesArr)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSimilarSeries indicates an expected call of GetSimilarSeries
func (mr *MockUsecaseMockRecorder) GetSimilarSeries(fid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarSeries", reflect.TypeOf((*MockUsecase)(nil).GetSimilarSeries), fid)
}
