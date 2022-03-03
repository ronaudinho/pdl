package pdl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Enrich wraps request to /person/enrich API.
func (p *Person) Enrich(params map[string]string) (map[string]interface{}, error) {
	if params == nil {
		return nil, ErrMissingParams
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/enrich", p.baseURL), nil)
	if err != nil {
		return nil, err
	}

	// this, up to response reading, is repetitive, maybe refactor later
	q := req.URL.Query()
	q.Add("api_key", p.PDL.apiKey)
	// params are not validated
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	// need to map this or simply return status code and leave it to consumer to decide?
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s status code: %d", req.URL.String(), res.StatusCode)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// better define a schema
	// it currently returns a map, which I am personally not a fan of since they are arbitrary,
	// but at this stage, map should be better than defining a concrete struct.
	// I mean https://docs.peopledatalabs.com/docs/fields is rather long.
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	// return as is
	// required params are not validated (if they appear in the response)
	return m, nil
}
