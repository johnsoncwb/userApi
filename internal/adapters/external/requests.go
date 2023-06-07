package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/johnsoncwb/userApi/internal/config"
	"github.com/johnsoncwb/userApi/internal/core/domain"
	"github.com/johnsoncwb/userApi/internal/utils/mmHttpClient"
)

type Requests struct {
}

func (r *Requests) GetAll(ctx context.Context) ([]domain.User, error) {

	client := mmHttpClient.NewHttpClient(ctx, 10*time.Second)

	url := config.GlobalConfig.UsersUrl

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao criar nova request",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return []domain.User{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao efetuar nova request",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return []domain.User{}, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao interpretar body da requisição",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return []domain.User{}, err
	}

	var data []domain.User

	err = json.Unmarshal(response, &data)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro no unmarshal",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return []domain.User{}, err
	}

	return data, nil

}

func (r *Requests) GetById(ctx context.Context) (domain.User, error) {
	client := mmHttpClient.NewHttpClient(ctx, 10*time.Second)
	id, ok := ctx.Value("id").(string)

	if !ok {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro recuperar id do contexto",
		// })
		return domain.User{}, errors.New("invalid context")
	}

	url := fmt.Sprintf("%s/%s", config.GlobalConfig.UsersUrl, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao criar nova request",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return domain.User{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao efetuar nova request",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return domain.User{}, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao interpretar body da requisição",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return domain.User{}, err
	}

	var data domain.User

	err = json.Unmarshal(response, &data)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro no unmarshal",
		// 	"url":     url,
		// 	"method":  http.MethodGet,
		// }).Err(err)
		return domain.User{}, err
	}

	return data, nil
}

func NewRequests() *Requests {
	return &Requests{}
}
