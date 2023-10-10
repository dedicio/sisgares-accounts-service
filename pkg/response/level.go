package response

import (
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
)

type LevelResponse struct {
	*dto.LevelResponseDto
}

func (pr *LevelResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewLevelResponse(product *dto.LevelResponseDto) *LevelResponse {
	return &LevelResponse{product}
}

type LevelsResponse struct {
	Levels []*dto.LevelResponseDto `json:"items"`
}

func (pr *LevelsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewLevelsResponse(products []*dto.LevelResponseDto) *LevelsResponse {
	return &LevelsResponse{products}
}
