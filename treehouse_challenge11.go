package channels

func readFromChannel() string {
  // CREATE A CHANNEL FOR string VALUES HERE
  channel := make(chan string)
  // HERE, CALL writeToChannel AS A GOROUTINE, AND PASS IT YOUR CHANNEL
  go writeToChannel(channel)
  // HERE, READ A STRING FROM YOUR CHANNEL AND RETURN IT
  return <-channel
}
