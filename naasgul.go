package naasgul

import "time"

// NoteID is the type for a Note's primary identifier.
type NoteID int

// SubID is the type for a Sub's primary identifier.
type SubID int

// GroupID is the type for a Group's primary identifier.
type GroupID int

// Note is the core entity for NaaSgul. It represents an individual notification that could be going to any
// number of subscribers.
type Note struct {
	ID        NoteID
	Type      string // Shoud be an enum of possible note types.
	Payload   []byte
	SendAt    *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Sub is an individual subscriber of notes.
type Sub struct {
	ID         SubID
	ExternalID string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

// Group is a grouping of subscribers for easy note blasting.
type Group struct {
	ID        GroupID
	Name      string
	Subs      []*Sub
	UpdatedAt *time.Time
	CreatedAt *time.Time
	DeletedAt *time.Time
}

// A NoteWriter knows how to bring new Notes into the world and share them with subscribers.
type NoteWriter interface {
	Create(*Note) (*Note, error)       // Create a new note.
	Send(*Note, []*Sub) (*Note, error) // Create and send a note.
	Forward(NoteID, []*Sub) error      // Forward an existing note.
	Cancel(NoteID) error               // Prevents note from being passed if it hasn't yet.
}

// A NoteReader can receive and retrieve Notes.
type NoteReader interface {
	Read(NoteID, SubID) error         // Mark a note as read for a sub.
	Get(NoteID) (*Note, error)        // Get an existing note.
	GetForSub(SubID) ([]*Note, error) // Get all notes for a given Sub.
}

// NewNote constructs a new Note entity and returns its pointer.
func NewNote(noteType string, payload []byte, sendAt *time.Time) *Note {
	return &Note{
		Type:    noteType,
		Payload: payload,
		SendAt:  sendAt,
	}
}

// NewSub constructs a new Sub entity and returns its pointer.
func NewSub(externalID string) *Sub {
	return &Sub{
		ExternalID: externalID,
	}
}
