package response

//"github.com/johnlonganecker/web-server/parser"

type Response struct {
	Version      string
	Status       string
	ReasonPhrase string
	Headers      string
	Body         string
}

func (r *Response) AddHeader(key string, value string) {
	r.Headers = r.Headers + key + ": " + value + "\n"
}

func (r *Response) String() string {
	return r.Version + " " + r.Status + " " + r.ReasonPhrase + "\n" + r.Headers + "\n" + r.Body
}
