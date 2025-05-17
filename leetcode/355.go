package leetcode

import "container/heap"

type Tweet struct {
	tweetId int
	time    int
	userId  int
}

// Min-Heap to keep most recent tweets
type TweetHeap []Tweet

func (th TweetHeap) Len() int           { return len(th) }
func (th TweetHeap) Less(i, j int) bool { return th[i].time > th[j].time }
func (th TweetHeap) Swap(i, j int)      { th[i], th[j] = th[j], th[i] }

func (th *TweetHeap) Pop() interface{} {
	old := *th
	n := len(old)
	x := old[n-1]
	*th = old[:n-1]
	return x
}

func (th *TweetHeap) Push(x interface{}) {
	*th = append(*th, x.(Tweet))
}

type Twitter struct {
	timestamp int
	followers map[int]map[int]struct{}
	tweets    map[int][]Tweet
}

func Constructor_355() Twitter {
	return Twitter{
		0,
		make(map[int]map[int]struct{}),
		make(map[int][]Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets[userId] = append(this.tweets[userId], Tweet{tweetId, this.timestamp, userId})
	this.timestamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	th := &TweetHeap{}
	heap.Init(th)

	// Include itself
	if _, ok := this.followers[userId]; !ok {
		followees := make(map[int]struct{})
		this.followers[userId] = followees
	}
	this.followers[userId][userId] = struct{}{}

	for followeeId := range this.followers[userId] {
		tweets := this.tweets[followeeId]
		for i := len(tweets) - 1; i >= 0 && len(tweets)-i <= 10; i-- {
			heap.Push(th, tweets[i])
		}
	}

	result := make([]int, 0, 10)
	for i := 0; i < 10 && th.Len() > 0; i++ {
		tweet := heap.Pop(th).(Tweet)
		result = append(result, tweet.tweetId)
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	followees, ok := this.followers[followerId]
	if !ok {
		followees = make(map[int]struct{})
		this.followers[followerId] = followees
	}

	followees[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followees, ok := this.followers[followerId]
	if !ok {
		return
	}
	delete(followees, followeeId)
}

// 355. Design Twitter
// https://leetcode.com/problems/design-twitter
////
// Recap - Heap
// 1. Create a struct
// 2. Make a heap as a list of the struct you just created
// 3. Implement the heap methods with PPLLS -> Pop, Push, Len, Less, Swap
// Note that Pop/Push should prefix with a pointer to struct, where the others can prefix with
// a reference to struct
//
// We first implement a struct `Tweet` which means a post or a message created and shared
// by a user on the platform Twitter
