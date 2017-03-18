package GoGym

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	HTTPMethodNotAllowed = 405
	HTTPOk               = 200
)

const (
	ServiceResponse = "Response"
)

// Response service
type Response struct {
	boss *Gym // Service Container

	rw         http.ResponseWriter
	statusCode int
	respone    interface{}
	header     http.Header
}

// Prepare is a function prepares the Response service
func (r *Response) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
}

// WhoIsYourBoss is a function sets the service container into the Response
func (r *Response) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

// CallYourBoss is a function gets the service container
func (r *Response) CallYourBoss() *Gym {
	return r.boss
}

// JsonResponse is a function prepares the JSON response
func (r *Response) JsonResponse(resp interface{}, statusCode int, header http.Header) {
	r.respone = resp
	r.statusCode = statusCode
	var respHeader http.Header
	if header != nil {
		respHeader = header
		respHeader.Add("Content-Type", "application/json")
	} else {
		respHeader.Add("Content-Type", "application/json")
	}

	r.header = respHeader
}

// wait is a function does preparation for sending response
func (r *Response) wait(rw http.ResponseWriter) {
	r.rw = rw
	r.statusCode = HTTPOk
}

// send is a function sending the http response
func (r *Response) send() {
	for k, v := range r.header {
		for _, h := range v {
			r.rw.Header().Add(k, h)
		}
	}
	r.rw.WriteHeader(r.statusCode)
	rsp, err := json.Marshal(r.respone)
	if err != nil {
		// TODO: logging error
		fmt.Println("JSON err:", err)
	}
	r.rw.Write(rsp)
}
