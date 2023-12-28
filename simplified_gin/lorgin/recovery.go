package lorgin

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				c.String(500, "Internal Server Error :%s", err)
			}
		}()
		c.Next()
	}
}
