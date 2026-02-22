package api

type UnrestrictLinkResponse struct {
	Id         string `json:"id"`
	Filename   string `json:"filename"`
	MimeType   string `json:"mimetype"`
	Filesize   int    `json:"filesize"`
	Link       string `json:"link"`
	Host       string `json:"host"`
	Chunks     int    `json:"chunks"`
	Crc        int    `json:"crc"`
	Download   string `json:"download"`
	Streamable int    `json:"streamable"`
}
