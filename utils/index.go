package utils

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

const SpotifyStateKey = "spotify_auth_state"

var Scope = strings.Join([]string{
	"user-read-private",
	"user-read-email",
	"user-follow-read",
	"user-read-currently-playing",
	"user-read-playback-state",
	"user-top-read",
	"user-read-recently-played",
}, " ")

func CopyString(s string) *string {
	copy := new(string)
	*copy = s
	return copy
}
