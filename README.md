# gLog

Your silver bullet logging solution. Because that is definitely what
the world needed: another logging library.

![Coverage](https://img.shields.io/static/v1?label=coverage&message=91%25&color=brightgreen&style=flat-square)
![current Version: v1.0.0](https://img.shields.io/static/v1?label=version&message=1.0.0&color=green&style=flat-square)
![License](https://img.shields.io/static/v1?label=license&message=MPL-2&color=blue&style=flat-square)
[![godocs.io](https://img.shields.io/badge/godoc-reference-blue?style=flat-square)](https://godocs.io/git.sr.ht/~poldi1405/glog)
![no non-stdlib deps that must be credited](https://img.shields.io/badge/external_dependencies-0-green?style=flat-square)[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmpldr%2Fglog.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmpldr%2Fglog?ref=badge_shield)
\*

## Features

gLog supports the following features:

- log-levels 
- smart-colors
- customization
- caller-insertion
- multiple outputs
- log rotation (coming soon™)
- panic-handling
- conditionals:
	- output per logging level
	- caller from a certain level
- sane defaults

## Output

The output follows this structure:

``` %level:\t%time – %caller – … ```

`%level` is coloured by  default when logging  to stdout  or stderr.
It  can be enabled and disabled at will.

`%time` is in the ISO-8601 format working with nanosecond precision
(if supported)

`%caller` is the  name  (format:  `package.function`)  of  the
function that is logging the message.  By default, this does not
happen for `INFO` and `WARNING`.

### Why `–` and not `-`

I  decided  to  use  [U+2013](https://codepoints.net/U+2013)  (EN
DASH)  over a normal [U+002D](https://codepoints.net/U+002D)
(HYPHEN-MINUS)  due  to multiple reasons; some of which are:

- `2D` is not at all unlikely to appear in a log-message
- they look very similar while they can still be distinguished
- it is also in most fonts used in a terminal
	- if it is not in your's think about changing fonts
- it allows for easy splitting and processing of logs with tools like
  `cut` or `awk`
- you don't have to type it by hand anyway

## Levels

| Level   | Description                                                       |
|---------|-------------------------------------------------------------------|
| TRACE   | Print *everything*. This is usually not wanted unless debugging.  |
| DEBUG   | Print *every* function call.                                      |
| INFO    | Print general status messages like HTTP-Requests (a good default) |
| WARNING | Handled errors. (the better default¹)                             |
| ERROR   | Non-Critical Errors like access denied                            |
| FATAL   | Errors that do not allow the Program to continue                  |

¹) "no news is good news" or so they say

## Notes

- Obviously, setting a low loglevel will slow down your program, as
  the writing is not buffered and getting the caller is relatively
  expensive
	- using StdOut/StdErr for logging decreases this time because
	  it does not need to write to disk

## Learn more

You can find more information in [this project's
wiki](https://man.sr.ht/~poldi1405/glog/)

### `NO_COLOR`

glog respects [`NO_COLOR`](https://no-color.org)

## Contribute

Contributions are welcome from anyone. Just send a patchset to
[~poldi1405/patches@lists.sr.ht](mailto:~poldi1405/patches@lists.sr.ht)
and wait for feedback. For general questions or other communications
feel free to drop a message to
[~poldi1405/discussion@lists.sr.ht](mailto:~poldi1405/discussion@lists.sr.ht)

Updates will be announced
[here](https://lists.sr.ht/~poldi1405/updates)

The changelog can be found
[here](https://lists.sr.ht/~poldi1405/updates?search=%5Bglog%5D)

## License

\* the claim 0 external dependencies refers to 0 dependencies that
   have to be credited. By crediting this package, the crediting for the
   only external dependency is also fulfilled

&copy; Moritz Poldrack and [contributors](CONTRIBUTORS.md)

gLog is a gun unter the MPL-2.0

To learn more you may:
- read the license-text [here](https://www.mozilla.org/en-US/MPL/2.0/)
- take a look at the official
  [FAQs](https://www.mozilla.org/en-US/MPL/2.0/FAQ/)
- take a look at a summary at
  [TLDRLegal](https://www.tldrlegal.com/l/mpl-2.0)


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmpldr%2Fglog.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmpldr%2Fglog?ref=badge_large)