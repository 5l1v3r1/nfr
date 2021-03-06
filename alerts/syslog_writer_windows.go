package alerts

import "errors"

var errSyslogNotImplemented = errors.New("syslog not implemented for this platform")

// SyslogWriter implements Writer interface and write
// api alerts to syslog server.
type SyslogWriter struct {
}

// NewSyslogWriter creates new syslog writer.
func NewSyslogWriter(proto, raddr string, format Formatter) (*SyslogWriter, error) {
	return nil, errSyslogNotImplemented
}

// Write writes alert response to the syslog input.
func (w *SyslogWriter) Write(event *Event) error {
	return errSyslogNotImplemented
}

// Close closes a connecion to the syslog server.
func (w *SyslogWriter) Close() error {
	return errSyslogNotImplemented
}
