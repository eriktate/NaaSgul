package mock

import "github.com/eriktate/naasgul"

// NoteReader mocks the NoteReader interface.
type NoteReader struct {
	ReadFn      func(naasgul.NoteID, naasgul.SubID) error
	ReadInvoked bool

	GetFn      func(naasgul.NoteID) (*naasgul.Note, error)
	GetInvoked bool

	GetForSubFn      func(naasgul.SubID) ([]*naasgul.Note, error)
	GetForSubInvoked bool

	PassThru bool
}

// Read mocked
func (m *NoteReader) Read(noteID naasgul.NoteID, subID naasgul.SubID) error {
	m.ReadInvoked = true
	if m.PassThru {
		return m.ReadFn(noteID, subID)
	}
	return nil
}

// Get mocked
func (m *NoteReader) Get(noteID naasgul.NoteID) (*naasgul.Note, error) {
	m.GetInvoked = true
	if m.PassThru {
		return m.GetFn(noteID)
	}
	return nil, nil
}

// GetForSub mocked
func (m *NoteReader) GetForSub(subID naasgul.SubID) ([]*naasgul.Note, error) {
	m.GetForSubInvoked = true
	if m.PassThru {
		return m.GetForSubFn(subID)
	}
	return nil, nil
}
