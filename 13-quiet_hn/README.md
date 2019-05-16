# Exercise 13 - "Quiet HN"

View on [gophercises.com](https://gophercises.com/exercises/quiet_hn).

## Description
One of the most common approaches to learning a new language is to rebuild things you know. As a result, you will often
find tons of clones for websites like Twitter, Hacker News (HN), Pinterest, and countless others. The clones are rarely
better than the original, but that isn't the point. The point is to build something you know so that you can eliminate
a lot of the guesswork and uncertainty that comes with building something new.

In this exercise we aren't going to be building a clone from scratch, but we are going to take a relatively simple
[Hacker News](https://news.ycombinator.com) clone (called
[Quiet Hacker News](https://github.com/tomspeak/quiet-hacker-news))
and use it to explore concurrency and caching in Go. That said, you are welcome to build your own HN clone before
moving forward with the exercise by reading what the current one does below and writing a similar server.

The application then renders all of those stories, along with some footer text that logs how long it took
to render the web page.

![example rendering of the Quiet HN page](https://www.dropbox.com/s/nexh2oql60a25df/Screenshot%202018-04-02%2017.34.01.png?dl=0&raw=1)

### Concurrency

Rather than focusing on how to build this application we are going to look at ways to add both concurrency and caching
in order to speed up the application. The first - concurrency - will be explored because it is a common reason to want
to check out Go, and it is nice to get a feel for how it works in the language. The second - caching - is important
because this is actually one of the easiest and most effective ways to speed up our application, and even outperforms
our concurrency changes.

1. Stories MUST retain their original order
2. Make sure you ALWAYS print out 30, and only 30, stories

### Caching
In addition to adding concurrency, add caching to the application. Your cache should store the results of the top
`numStories` stories so that subsequent web requests don't require additional API calls, but that cache should expire
at some point after which time more API calls will be needed to update the cache.

How you implement this is up to you, but you should definitely consider the fact that many web requests can be
processed at the same time, so you may need to take race conditions into consideration. A great way to test this
is the [-race](https://blog.golang.org/race-detector) flag.
