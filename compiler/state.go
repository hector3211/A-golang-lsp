package compiler

import (
	"fmt"
	"golang-lsp/lsp"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func PublishDiagnostic(text string) []lsp.Diagnostic {
	document := text
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(document, "\n") {
		hasSemiColon := strings.Contains(line, ";")
		if !hasSemiColon {
			idx := strings.Index(line, ";")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Rang:     lineRange(row, idx, idx+len(";")),
				Severity: 2,
				Source:   "",
				Message:  "Forgot to add semicolon!",
			})
		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return PublishDiagnostic(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return PublishDiagnostic(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, characters: %d", uri, len(document)),
		},
	}
}
func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
func (s *State) CodeAction(id int, uri string, position lsp.Position) lsp.CodeActionResponse {
	text := s.Documents[uri]
	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "goodbye")
		if idx > 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   lineRange(row, idx, idx+len("goodbye")),
					NewText: "bye",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Relace Hello with a alternitive in another language",
				Edit:  &lsp.WorkSpaceEdit{Changes: replaceChange},
			})

			upperCaseChange := map[string][]lsp.TextEdit{}
			upperCaseChange[uri] = []lsp.TextEdit{
				{
					Range:   lineRange(row, idx, idx+len("goodbye")),
					NewText: "BYE",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Make bye upper case.",
				Edit:  &lsp.WorkSpaceEdit{Changes: upperCaseChange},
			})
		}
	}
	return lsp.CodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
}

func lineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}

func (s *State) CodeCompletion(id int, uri string) lsp.CompletionResponse {
	completions := []lsp.CompletionItem{
		{
			Label:         "Nvim (FTW)",
			Detail:        "Fastest editor out there",
			Documentation: "Longer version of detail here.",
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: completions,
	}
}
