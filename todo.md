# ToDo

This is a Todo-List of what steps are still required to reach the next milestone.

Next Milestone: *v1.0*

## v1.0

- [x] log-levels
- [x] color
	- [x] detect whether output is on a terminal
- [x] caller
	- [x] limit by level
- [x] output
	- [x] pipe certain levels into certains outputs (`io.Writer` must be implemented)
- [ ] panic handler
- [x] the actual logging
	- [x] limit by log-level
	- [x] pipe certain levels into certain outputs

## v1.x

- [ ] allow writing to syslog/eventlog
- [ ] logrotation
