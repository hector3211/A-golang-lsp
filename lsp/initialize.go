package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`

	// Tons more for a proper lsp
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities SeverCapabilites `json:"Capabilities"`
	ServerInfo   ServerInfo       `json:"serverInfo"`
}

type SeverCapabilites struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: SeverCapabilites{},
			ServerInfo: ServerInfo{
				Name:    "educationallsp",
				Version: "0.0.0.0-beta1.final",
			},
		},
	}
}
