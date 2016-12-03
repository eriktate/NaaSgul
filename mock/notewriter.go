package mock

import "github.com/eriktate/naasgul"

// NoteWriter mocks the NoteWriter interface.
type NoteWriter struct {
	CreateFn      func(*naasgul.Note) (*naasgul.Note, error)
	CreateInvoked bool

	SendFn      func(*naasgul.Note, []*naasgul.Sub) (*naasgul.Note, error)
	SendInvoked bool

	ForwardFn      func(naasgul.NoteID, []*naasgul.Sub) error
	ForwardInvoked bool

	CancelFn      func(naasgul.NoteID) error
	CancelInvoked bool

	PassThru bool
}

// Create mocked
func (m *NoteWriter) Create(note *naasgul.Note) (*naasgul.Note, error) {
	m.CreateInvoked = true
	if m.PassThru {
		return m.CreateFn(note)
	}
	return nil, nil
}

// Send mocked
func (m *NoteWriter) Send(note *naasgul.Note, subs []*naasgul.Sub) (*naasgul.Note, error) {
	m.SendInvoked = true
	if m.PassThru {
		return m.SendFn(note, subs)
	}
	return nil, nil
}

// Forward mocked
func (m *NoteWriter) Forward(noteID naasgul.NoteID, subs []*naasgul.Sub) error {
	m.ForwardInvoked = true
	if m.PassThru {
		return m.ForwardFn(noteID, subs)
	}
	return nil
}

// Cancel mocked
func (m *NoteWriter) Cancel(noteID naasgul.NoteID) error {
	m.CancelInvoked = true
	if m.PassThru {
		return m.CancelFn(noteID)
	}
	return nil
}
