package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"` // DocumentURI
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParems `json:"params"`
}

type DidChangeTextDocumentParems struct {
	TextDocumentItem VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges   []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int `json:"version,omitempty"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
	// Range Range `json:"range"`
	// RangeLength uint `json:"rangeLength"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type WorkSpaceEdit struct {
	Changes map[string][]TextEdit `json:"changes"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
