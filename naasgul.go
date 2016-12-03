package naasgul

import "time"

// NoteID is the type for a Note's primary identifier.
type NoteID int

// SubID is the type for a Sub's primary identifier.
type SubID int

// GroupID is the type for a Group's primary identifier.
type GroupID int

// Note is the core entity for NaaSgul. It represents an
// individual notification that could be going to any
// number of subscribers.
type Note struct {
	ID        NoteID
	Type      string // Shoud be an enum of possible note types.
	Payload   string
	SendAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Sub is an individual subscriber of notes.
type Sub struct {
	ID         SubID
	ExternalID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// Group is a grouping of subscribers for easy note blasting.
type Group struct {
	ID        GroupID
	Name      string
	Subs      []*Sub
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// NoteWriter is the API for working directly with
// Note entities.
type NoteWriter interface {
	Write(*Note) NoteID   // Write a new note.
	Peek(NoteID) *Note    // Peek at an existing note.
	Pass(*Note, []*Sub)   // Write and pass new note.
	Share(NoteID, []*Sub) // Share existing note.
	Hide(NoteID)          // Prevents note from being passed if it hasn't yet.

}
