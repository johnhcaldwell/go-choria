// generated code; DO NOT EDIT

package scoutclient

func (c *ScoutClient) debugf(msg string, a ...any) {
	c.clientOpts.logger.Debugf(msg, a...)
}

func (c *ScoutClient) infof(msg string, a ...any) {
	c.clientOpts.logger.Infof(msg, a...)
}

func (c *ScoutClient) errorf(msg string, a ...any) {
	c.clientOpts.logger.Errorf(msg, a...)
}
