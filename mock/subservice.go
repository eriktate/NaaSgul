package mock

import "github.com/eriktate/naasgul"

// SubService mocks the SubService interface.
type SubService struct {
	CreateFn      func(*naasgul.Sub) (*naasgul.Sub, error)
	CreateInvoked bool

	GetFn      func(naasgul.SubID) (*naasgul.Sub, error)
	GetInvoked bool

	PassThru bool
}

// Create mocked
func (m *SubService) Create(sub *naasgul.Sub) (*naasgul.Sub, error) {
	m.CreateInvoked = true
	if m.PassThru {
		return m.CreateFn(sub)
	}
	return nil, nil
}

// Get mocked
func (m *SubService) Get(subID naasgul.SubID) (*naasgul.Sub, error) {
	m.GetInvoked = true
	if m.PassThru {
		return m.GetFn(subID)
	}
	return nil, nil
}
