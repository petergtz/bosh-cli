package director

import (
	"fmt"
	"net/http"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	gourl "net/url"
)

type Config struct {
	Content string
}

func (d DirectorImpl) LatestConfig(t string, name string) (Config, error) {
	resps, err := d.client.Configs(t, name)

	if err != nil {
		return Config{}, err
	}

	if len(resps) == 0 {
		return Config{}, bosherr.Error("No config")
	}

	return resps[0], nil
}

func (d DirectorImpl) UpdateConfig(t string, name string, content []byte) error {
	return d.client.UpdateConfig(t, name, content)
}

func (c Client) Configs(configType string, name string) ([]Config, error) {
	var resps []Config

	query := gourl.Values{}
	query.Add("name", name)
	query.Add("limit", "1")
	path := fmt.Sprintf("/configs/%s?%s", configType, query.Encode())

	err := c.clientRequest.Get(path, &resps)
	if err != nil {
		return resps, bosherr.WrapErrorf(err, "Finding configs")
	}

	return resps, nil
}

func (c Client) UpdateConfig(configType string, name string, content []byte) error {
	query := gourl.Values{}
	query.Add("name", name)
	path := fmt.Sprintf("/configs/%s?%s", configType, query.Encode())

	setHeaders := func(req *http.Request) {
		req.Header.Add("Content-Type", "text/yaml")
	}

	_, _, err := c.clientRequest.RawPost(path, content, setHeaders)
	if err != nil {
		return bosherr.WrapErrorf(err, "Updating config")
	}

	return nil
}
