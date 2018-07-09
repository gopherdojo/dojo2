package omikuji

import (
	"math/rand"
	"time"
)

// Service provides omikuji.
type Service interface {
	Hiku() Omikuji
}

// New returns a new Service.
func New() Service {
	return &DefaultService{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// DefaultService is a default implementation of Service.
type DefaultService struct {
	random *rand.Rand       // Random generator
	time   func() time.Time // Time provider (optional)
}

// Hiku peforms おみくじを引く.
// If the current date is shogatsu, it always returns the DaiKichi.
func (s *DefaultService) Hiku() Omikuji {
	if s.isShogatsu() {
		return DaiKichi
	}
	n := s.random.Intn(6)
	switch n {
	default:
		return Kyo
	case 2, 3:
		return ShoKichi
	case 4:
		return ChuKichi
	case 5:
		return DaiKichi
	}
}

func (s *DefaultService) isShogatsu() bool {
	now := time.Now()
	if s.time != nil {
		now = s.time()
	}
	if now.Month() == time.January {
		switch now.Day() {
		case 1, 2, 3:
			return true
		}
	}
	return false
}
