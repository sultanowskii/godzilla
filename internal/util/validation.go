package util

import (
	"errors"
	"net/url"
	"regexp"

	"github.com/sultanowskii/godzilla/pkg/storage"
)

var suffixRegex = regexp.MustCompile(`^[\w\-.\\:]+$`)

func IsSuffixValid(s string) bool {
	return suffixRegex.MatchString(s)
}

func ValidateCustomSuffix(s string) error {
	if !IsSuffixValid(s) {
		return errors.New("suffix contains invalid characters")
	}

	client := storage.RedisClient

	suffixExists := client.Exists(storage.RedisContext, s).Val()

	if suffixExists == 1 {
		return errors.New("suffix already exists")
	}

	return nil
}

func ValidateUrl(s string) error {
	u, err := url.Parse(s)

	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("please provide absolute url")
	}

	return nil
}
