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
