package naasgul

//The only imports allowed in this package are for specific data types.
import (
	"time"

	"github.com/satori/go.uuid"
)

// NoteID aliases the type used to uniquely identify Notes.
type NoteID uuid.UUID

//SubID aliases the type used to uniquely identify Subs.
type SubID uuid.UUID

// Note is the notification model for NaaSgul. It represents the core data points that make up
// an individual notification and acts as a receiver for helper functions pertaining to specific
// notifications. In order to use the helper functions, a Note should be given a noteService.
type Note struct {
	ID        NoteID
	Subject   string
	Body      string
	From      *Sub
	CreatedAt time.Time

	service NoteService
}

// Sub is the subscriber model for NaaSgul. It represents the core data points that make up an
// individual subscriber and acts as a receiver for helper functions pertaining to specific
// subscribers.
type Sub struct {
	ID         SubID
	ExternalID string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time

	service SubService
}

// NoteService defines the things a NoteService should be able to do. This should be representative
// of all actions that can be taken on a Note.
type NoteService interface {
	Note(id NoteID) (*Note, error)
	NotesBySub(id SubID) ([]*Note, error)
	Send(note *Note, subs []SubID) error
	AddSubs(id NoteID, subs []SubID) error
	GetReceived(id NoteID) ([]*Sub, error)
	GetRecipients(id NoteID) ([]*Sub, error)
}

// SubService defines the things a SubService should be able to do. This should be representative
// of all actions that can be taken on a Sub.
type SubService interface {
	Sub(id SubID) (*Sub, error)
	Receive(id NoteID) error
	Create(sub *Sub) error
	Update(sub *Sub) error
	Delete(id SubID) error
}

// Dispatcher defines the things a Note dispatcher should do. A dispatcher should be able to be queued with notes
// and be told to force a dispatch.
type Dispatcher interface {
	Queue(note *Note, rcpts []SubID)
	Dispatch()
}

// Subs returns a slice of the Subs this Note was sent to.
func (note *Note) Subs() ([]*Sub, error) {
	return note.service.GetRecipients(note.ID)
}

// Received returns a slice of the Subs that have seen this Note.
func (note *Note) Received() ([]*Sub, error) {
	return note.service.GetReceived(note.ID)
}

// SetNoteService provides a way to change the NoteService being used by the Note.
func (note *Note) SetNoteService(ns NoteService) {
	note.service = ns
}

// Receive signals that the Sub has received the specified Note.
func (sub *Sub) Receive(id NoteID) error {
	return sub.service.Receive(id)
}

// SetSubService provides a way to change the SubService being used by the Note.
func (sub *Sub) SetSubService(ss SubService) {
	sub.service = ss
}
