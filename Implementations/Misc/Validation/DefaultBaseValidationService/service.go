package DefaultBaseValidationService

import (
	gov "github.com/asaskevich/govalidator"
	"strings"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Validation"
)

type service struct{}

func (s *service) extendedIsEmail(str string) bool {
	return gov.IsEmail(str) && len(strings.TrimSpace(str)) >= 3
}

//Public methods
func (s *service) InRange(value, left, right float64) bool   { return gov.InRange(value, left, right) }
func (s *service) IsASCII(str string) bool                   { return gov.IsASCII(str) }
func (s *service) IsAlpha(str string) bool                   { return gov.IsAlpha(str) }
func (s *service) IsAlphanumeric(str string) bool            { return gov.IsAlphanumeric(str) }
func (s *service) IsBase64(str string) bool                  { return gov.IsBase64(str) }
func (s *service) IsCreditCard(str string) bool              { return gov.IsCreditCard(str) }
func (s *service) IsEmail(str string) bool                   { return s.extendedIsEmail(str) }
func (s *service) IsEmptyStringOrWhitespace(str string) bool { return strings.TrimSpace(str) == "" }
func (s *service) IsFilePath(str string) (bool, int)         { return gov.IsFilePath(str) }
func (s *service) IsFloat(str string) bool                   { return gov.IsFloat(str) }
func (s *service) IsHexadecimal(str string) bool             { return gov.IsHexadecimal(str) }
func (s *service) IsHexcolor(str string) bool                { return gov.IsHexcolor(str) }
func (s *service) IsIP(str string) bool                      { return gov.IsIP(str) }
func (s *service) IsIPv4(str string) bool                    { return gov.IsIPv4(str) }
func (s *service) IsIPv6(str string) bool                    { return gov.IsIPv6(str) }
func (s *service) IsInt(str string) bool                     { return gov.IsInt(str) }
func (s *service) IsJSON(str string) bool                    { return gov.IsJSON(str) }
func (s *service) IsLatitude(str string) bool                { return gov.IsLatitude(str) }
func (s *service) IsLongitude(str string) bool               { return gov.IsLongitude(str) }
func (s *service) IsNatural(value float64) bool              { return gov.IsNatural(value) }
func (s *service) IsNegative(value float64) bool             { return gov.IsNegative(value) }
func (s *service) IsNonNegative(value float64) bool          { return gov.IsNonNegative(value) }
func (s *service) IsNonPositive(value float64) bool          { return gov.IsNonPositive(value) }
func (s *service) IsNumeric(str string) bool                 { return gov.IsNumeric(str) }
func (s *service) IsPositive(value float64) bool             { return gov.IsPositive(value) }
func (s *service) IsWhole(value float64) bool                { return gov.IsWhole(value) }

func New() BaseValidationService {
	return &service{}
}
