package director

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type Config struct {
	Content string
}

func (d DirectorImpl) LatestConfig() (Config, error) {
	resps, err := d.client.Configs()

	if err != nil {
		return Config{}, err
	}

	if len(resps) == 0 {
		return Config{}, bosherr.Error("No config")
	}

	return resps[0], nil
}

func (c Client) Configs() ([]Config, error) {
	var resps []Config

	err := c.clientRequest.Get("/configs?limit=1", &resps)
	if err != nil {
		return resps, bosherr.WrapErrorf(err, "Finding configs")
	}

	return resps, nil
}
