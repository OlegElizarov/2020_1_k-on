package series

import "github.com/go-park-mail-ru/2020_1_k-on/application/models"

type Repository interface {
	GetSeriesByID(id uint) (models.Series, bool)
	GetSeriesSeasons(id uint) (models.Seasons, bool)
	GetSeasonEpisodes(id uint) (models.Episodes, bool)
	FilterSeriesList(map[string][]string) (*models.SeriesArr, bool)
	FilterSeriesData() (map[string]interface{}, bool)
}
