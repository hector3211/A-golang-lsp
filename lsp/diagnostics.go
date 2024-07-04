package lsp

type PublishDiagnosticsNotification struct {
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Rang    Range  `json:"range"`
	Message string `json:"message"`
}
