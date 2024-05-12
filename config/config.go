package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var Default Config

type Config struct {
	Addr       string `yaml:"addr"`
	ServiceURL string `yaml:"service_url"`
	Mode       string `yaml:"mode"`
	SiteURL    string `yaml:"site_url"`
	DouYinApi  string `yaml:"douyin_api"`

	Authorization struct {
		Enable   bool   `yaml:"enable"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"authorization"`
	DouYin struct {
		Cookie string `yaml:"cookie"`
		Domain string `yaml:"domain"`
		Aweme  struct {
			Base  string `yaml:"base"`
			Video string `yaml:"video"`
			Image string `yaml:"image"`
		} `yaml:"aweme"`
	} `yaml:"douyin"`
	XiaoHongShu struct {
		Cookie string `yaml:"cookie"`
		Domain string `yaml:"domain"`
	} `yaml:"xiaohongshu"`
}

func (c Config) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// GetAssetLocation searches for `file` in certain locations
func GetAssetLocation(file string) string {
	defPath := filepath.Join(os.Getenv("HOME"), ".douyinbot", file)
	for _, p := range []string{
		defPath,
		filepath.Join("/usr/local/share/douyinbot/", file),
		filepath.Join("/usr/share/douyinbot/", file),
		filepath.Join("/opt/share/douyinbot/", file),
	} {
		log.Println(p)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}
		return p
	}
	return defPath
}
