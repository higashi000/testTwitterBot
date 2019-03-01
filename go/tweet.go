package main

import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
  "os"
)
func main() {
  api := GetTwitterApi()

  text := "Go言語からツイートしています"
  tweet, err := api.PostTweet(text, nil)
  if err != nil {
    panic(err)
  }

  fmt.Print(tweet.Text)
}

func GetTwitterApi() *anaconda.TwitterApi {
  anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
  anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
  api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"),
                                os.Getenv("ACCESS_TOKEN_SECRET"))
  return api
}
