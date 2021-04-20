# gLog

Your silver bullet logging solution. Because that is definitely what the world
needed: another logging library.

## Features

gLog supports the following features:

- log-levels 
- smart-colors
- customization
- caller-insertion
- mutiple outputs
- log rotation (coming soon™)
- panic-handling <!-- ToDo -->
- conditionals:
	- output per logging level
	- caller from a certain level
- sane defaults

## Output

The output follows this structure:

```
%level:\t%time – %caller – …
```

`%level` is  colored by  default when logging  to stdout  or stderr.  It  can be
enabled and disabled at will.

`%time` is in the ISO-8601 format working with nanosecond precision (if supported)

`%caller` is the  name  (format:  `package.function`)  of  the  function that is
logging the message.  By default this does not happen for `INFO` and `WARNING`.

### Why `–` and not `-`

I  decided to  use [U+2013](https://codepoints.net/U+2013)  (EN DASH)  because a
normal [U+002D](https://codepoints.net/U+002D)  (HYPHEN-MINUS)  due  to multiple
reasons; some of which are:

- `2D` is not at all unlikely to appear in a log-message
	- using `2013` allows for an easy way of splitting log lines without using RegEx
- they look very similar while they can still be distinguished
- it is also in most fonts used in a terminal
	- if it is not in your's think about changing fonts
- it allows for easy splitting and processing of logs with tools like `cut` or `awk`
- you don't have to type it by hand anyway

## Levels

| Level   | Description                                                      |
|---------|------------------------------------------------------------------|
| TRACE   | Print *everything*. This is usually not wanted unless debugging. |
| DEBUG   | Print *every* function call.                                     |
| INFO    | Print general statusmessages like HTTP-Requests (a good default) |
| WARNING | Handled errors. (the better default¹)                            |
| ERROR   | Non-Critical Errors like access denied                           |
| FATAL   | Errors that do not allow the Program to continue                 |

¹) "no news is good news" or so they say

## Environment Variables

All environment-variable *values* (not the variables themselves) are case insensitive.

### `GLOG_COLOR`

- `ALWAYS`, `ON`, `1`
	- automatically sets OverwriteColor to 1, thereby enabling color even 
	on file outputs

- `NEVER`, `OFF`, `-1`
	- automatically sets OverwriteColor to -1, thereby disabling color even
	on terminals

### `GLOG_LEVEL`

Sets the level to the specified value. Special cases:

- `VERBOSE`
	- an alias for `TRACE`

- `MUTE`, `SILENT`
	- disable every and all log-output

### `GLOG_METALOGGER`

- `1`
	- enable the built in meta-logger (the logger, logging activities of
	the logging library). *DO NOT USE THIS UNLESS YOU HAVE GOOD REASON!* It
	look horrible.

## License

&copy; Moritz Poldrack

gLog is a gun unter the MPL-2.0

To learn more you may:
- read the license-text [here](https://www.mozilla.org/en-US/MPL/2.0/)
- take a look at the official [FAQs](https://www.mozilla.org/en-US/MPL/2.0/FAQ/)
- take a look at a summary at [TLDRLegal](https://www.tldrlegal.com/l/mpl-2.0)
