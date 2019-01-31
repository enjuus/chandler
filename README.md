# chandler

Downloader for 4chan threads.

## Installing

`go get github.com/enjuus/chandler`


## Usage

```
Usage: chandler [options]
Options:
  -d, --destination string
        The path to save to. See README for more options. [required]
  -h, --help
        Display this help message
  -s, --source string
        The thread to download [required]
```

`--destination` supports 3 variables:

`{THREAD}`: The title of the thread. If no title is found, the thread's ID will be used.

`{THREADID}`: The ID of the thread.

`{BOARD}`: The board the thread is on.


for example:

`chandler -s http://boards.4chan.org/wg/thread/7289180 -d "/home/user/pictures/{BOARD}/{THREAD}"`


### TODO

- create thread watcher
