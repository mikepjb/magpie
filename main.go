package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	twitterAccounts   = []string{
		"edlatimore",
		"AJA_Cortes",
		"ROGUEWEALTH",
		"TellYourSonThis",
		"AndraZaharia",
		"nntaleb",
		"joserosado",
	}
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	tweets, _ := api.GetUserTimeline(url.Values{
		"screen_name":     []string{"mikepjb"},
		"exclude_replies": []string{"true"},
	})

	for _, tweet := range tweets {
		fmt.Println(tweet.RetweetCount, tweet.FavoriteCount)
		fmt.Println(tweet.Text)
	}

	// searchResult, _ := api.GetSearch("@mikepjb", nil)
	// for _, tweet := range searchResult.Statuses {
	// 	fmt.Println(tweet.Text)
	// }
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
