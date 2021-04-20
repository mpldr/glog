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
- panic-handling
- conditionals:
	- output per logging level
	- caller from a certain level
- sane defaults

## Output

The output follows this structure:

```
%level:\t%time – %caller – …
```

`%level` is colored by default when logging to stdout or stderr. It can be
enabled and disabled at will.

`%time` is in the ISO-8601 format working with nanosecond precision (if supported)

`%caller` is the name (format: `package.function`) of the function that is logging
the message. By default this does not happen for `INFO` and `WARNING`.

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

## License

&copy; Moritz Poldrack

gLog is a gun unter the MPL-2.0

To learn more you may:
- read the license-text [here](https://www.mozilla.org/en-US/MPL/2.0/)
- take a look at the official [FAQs](https://www.mozilla.org/en-US/MPL/2.0/FAQ/)
- take a look at a summary at [TLDRLegal](https://www.tldrlegal.com/l/mpl-2.0)
