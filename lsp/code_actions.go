package lsp

type CodeActionRequest struct {
	Request
	Params CodeActionParams `json:"params"`
}

type CodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"context"`
}

type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkSpaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

type CodeActionContext struct {
	// Diagnostics []Diagnostic `json:"diagnostics"`
	// Only        *[]string    `json:"only"`
	// TiggerKind  *int         `json:"triggerKind"`
}

// type Diagnostic struct {
// 	Range   Range  `json:"range"`
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// }
