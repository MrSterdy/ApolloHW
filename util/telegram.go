package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/MrSterdy/ApolloHW/bot"
	"github.com/MrSterdy/ApolloHW/model/student"

	"github.com/buger/jsonparser"
)

var key []byte

func init() {
	hKey := hmac.New(sha256.New, []byte("WebAppData"))
	hKey.Write([]byte(os.Getenv("BOT_TOKEN")))

	key = hKey.Sum(nil)
}

func ParseInitData(initData string) (user int64, ok bool) {
	if initData == "" {
		return
	}

	query, err := url.QueryUnescape(initData)
	if err != nil {
		return
	}

	var hash string
	var sortedParams []string
	for _, param := range strings.Split(query, "&") {
		k, v, _ := strings.Cut(param, "=")

		switch {
		case hash == "" && k == "hash":
			hash = v
		case user == 0 && k == "user":
			user, err = jsonparser.GetInt([]byte(v), "id")
			if err != nil {
				return
			}

			fallthrough
		default:
			sortedParams = append(sortedParams, param)
		}
	}

	sort.Strings(sortedParams)

	checkHash := hmac.New(sha256.New, key)
	checkHash.Write([]byte(strings.Join(sortedParams, "\n")))

	ok = hex.EncodeToString(checkHash.Sum(nil)) == hash

	return
}

func SendNotification(message string, minRole student.Role) {
	for _, s := range student.GetAll() {
		if s.Role >= minRole {
			if _, err := bot.Bot.SendMessage(s.ID, message, nil); err != nil {
				panic(err)
			}
		}
	}
}
