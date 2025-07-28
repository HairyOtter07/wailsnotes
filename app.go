package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// App struct
type App struct {
	ctx    context.Context
	notes  []Note
	nextID int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		notes:  []Note{},
		nextID: 1,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateNote(title, content string) {
	id := fmt.Sprintf("%d", a.nextID)
	a.nextID++
	a.notes = append(a.notes, Note{
		ID:      id,
		Title:   title,
		Content: content,
	})
	runtime.EventsEmit(a.ctx, "updateNotes", a.notes)
}

func (a *App) GetNotes() []Note {
	return a.notes
}

func (a *App) GetNote(id string) (Note, error) {
	for _, note := range a.notes {
		if note.ID == id {
			return note, nil
		}
	}
	return Note{}, fmt.Errorf("note not found")
}

func (a *App) UpdateNote(id string, title, content string) error {
	for i, note := range a.notes {
		if note.ID == id {
			a.notes[i] = Note{
				ID:      id,
				Title:   title,
				Content: content,
			}
			runtime.EventsEmit(a.ctx, "updateNotes", a.notes)
			return nil
		}
	}
	return fmt.Errorf("note not found")
}

func (a *App) DeleteNote(id string) error {
	for i, note := range a.notes {
		if note.ID == id {
			a.notes = append(a.notes[:i], a.notes[i+1:]...)
			runtime.EventsEmit(a.ctx, "updateNotes", a.notes)
			return nil
		}
	}
	return fmt.Errorf("note not found")
}
