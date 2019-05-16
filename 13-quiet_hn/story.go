package main

type Story struct {
	// The item's unique id.
	Id int64 `json:"id"`
	// true if the item is deleted.
	Deleted bool `json:"deleted,omitempty"`
	// The type of item. One of "job", "story", "comment", "poll", or "pollopt".
	Type string `json:"type"`
	// The username of the item's author.
	By string `json:"by"`
	// Creation date of the item, in Unix Time.
	Time int64 `json:"time"`
	// The comment, story or poll text. HTML.
	Text string `json:"text,omitempty"`
	// true if the item is dead.
	Dead bool `json:"dead,omitempty"`
	// The comment's parent: either another comment or the relevant story.
	Parent int64 `json:"parent,omitempty"`
	// The pollopt's associated poll.
	Poll int64 `json:"poll,omitempty"`
	// The ids of the item's comments, in ranked display order.
	Kids []int64 `json:"kids"`
	// The URL of the story.
	Url string `json:"url"`
	// The story's score, or the votes for a pollopt.
	Score int64 `json:"score"`
	// The title of the story, poll or job.
	Title string `json:"title"`
	// A list of related pollopts, in display order.
	Parts []int64 `json:"parts,omitempty"`
	// In the case of stories or polls, the total comment count.
	Descendants int64 `json:"descendants"`
}

type StoryResponse struct {
	Story
	Err error
}

type ByIdsDescendant []StoryResponse

func (a ByIdsDescendant) Len() int {
	return len(a)
}

func (a ByIdsDescendant) Less(i, j int) bool {
	return a[i].Id > a[j].Id
}

func (a ByIdsDescendant) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
