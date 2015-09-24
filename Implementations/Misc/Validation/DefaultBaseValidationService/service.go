package DefaultBaseValidationService

import (
	"github.com/asaskevich/govalidator"
	"strings"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Validation"
)

type service struct{}

func (s *service) IsEmail(str string) bool {
	return govalidator.IsEmail(str) && len(strings.TrimSpace(str)) >= 3
}

func New() BaseValidationService {
	return &service{}
}
