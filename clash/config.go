package clash

import (
	"bufio"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/constant"
)

func (c *Clash) ConfigFromFile() *config.Config {
	conf, _ := config.Parse(constant.Path.Config())

	return conf
}

func (c *Clash) UpdateConfigFromRemote() error {
	conf := constant.Path.Config()

	if _, err := os.Stat(conf); err != nil {
		return err
	}

	remote, err := func() ([]byte, error) {
		f, err := os.Open(conf)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		b := bufio.NewReader(f)
		re, _, err := b.ReadLine()
		if err != nil {
			return nil, err
		}

		if !strings.HasPrefix(string(re), "#!") {
			return nil, errors.New("Config: No Remote URL is Given")
		}

		return re, nil
	}()
	if err != nil {
		return err
	}

	body, err := func() ([]byte, error) {
		resp, err := http.Get(string(remote[2:]))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return ioutil.ReadAll(resp.Body)
	}()
	if err != nil {
		return err
	}

	return func() error {
		err = os.Remove(conf)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(conf, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(remote)
		if err != nil {
			return err
		}

		_, err = f.Write([]byte("\n"))
		if err != nil {
			return err
		}

		_, err = f.Write(body)
		if err != nil {
			return err
		}

		return nil
	}()
}
