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
  -i, --interval int
        The times to check per second (default 20)
  -s, --source string
        The thread to download [required]
  -v, --verbose
        Enable output
  -w, --watcher
        Watch the thread for new files
```

`--destination` supports 3 variables:

`{THREAD}`: The title of the thread. If no title is found, the thread's ID will be used.

`{THREADID}`: The ID of the thread.

`{BOARD}`: The board the thread is on.


for example:

`chandler -s http://boards.4chan.org/wg/thread/7289180 -d "/home/user/pictures/{BOARD}/{THREAD}"`


The thread watcher can be enabled with the `--watcher` flag. By default it will check the current thread every 20 seconds. To add your own interval, use `--interval` (in seconds).

### TODO

- Find a cleaner solution for exiting the watcher when the thread died

### Thanks

Thank you [kori](https://github.com/kori) for making me rename this into the way funnier name.

Thank you [eti](https://github.com/eti0) for the idea with the destination variables.
