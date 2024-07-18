package handlerIMPL

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
	"io"
	"net/http"
)

type Handler struct {
	log *logger.Logger
	u   *usecase.UseCase
}

// valid Валидация полученных данных
func valid(v interface{}) error {
	validate := validator.New()
	// Checking whether JSON has all the attributes
	if err := validate.Struct(v); err != nil {
		// Create custom error
		var notFoundAttr string
		for idx, attr := range err.(validator.ValidationErrors) {
			if idx == 0 {
				notFoundAttr = attr.Field()
			} else {
				notFoundAttr = notFoundAttr + ", " + attr.Field()
			}
		}
		return errors.New(fmt.Sprintf("The JSON file has no attributes: %s", notFoundAttr))
	}

	return nil
}

func (h *Handler) decodeIO(bodyIo io.ReadCloser) ([]byte, error) {
	// Декодирование Body в json формат
	bodyJSON, err := io.ReadAll(bodyIo)
	if err != nil {
		h.log.Log.Error().Err(err)
		return nil, err
	}
	// Закрытие Body
	defer bodyIo.Close()

	return bodyJSON, err
}

func (h *Handler) convertBodyToStruct(bodyIo io.ReadCloser, iStruct *interface{}) error {
	// Декодирование BodyIO в json формат
	bodyJSON, err := h.decodeIO(bodyIo)
	if err != nil {
		return err
	}
	// Записываем JSON тело в структуру
	if err := json.Unmarshal(bodyJSON, &iStruct); err != nil {
		return err
	}
	// Валидация данных
	if err := valid(&iStruct); err != nil {
		return err
	}
	return nil
}

func New(log *logger.Logger, u *usecase.UseCase) *Handler {
	return &Handler{
		log: log,
		u:   u,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	req := map[string]int{"status": http.StatusCreated}

	jsonReq, _ := json.Marshal(req)
	w.Write(jsonReq)
}
