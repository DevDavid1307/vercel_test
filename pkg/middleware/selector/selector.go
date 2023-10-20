package selector

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Builder struct {
	prefix []string
	path   []string

	ms []gin.HandlerFunc
}

func (s *Builder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !s.match(c.Request.URL.Path) {
			c.Next()
			return
		}
		for _, m := range s.ms {
			m(c)
		}
		c.Next()
	}
}

func (s *Builder) match(url string) bool {
	for _, p := range s.path {
		if url == p {
			return true
		}
	}
	for _, p := range s.prefix {
		if strings.HasPrefix(url, p) {
			return true
		}
	}
	return false
}

func (s *Builder) Prefix(prefix ...string) *Builder {
	s.prefix = prefix
	return s
}

func (s *Builder) Path(path ...string) *Builder {
	s.path = path
	return s
}

func Selector(ms ...gin.HandlerFunc) *Builder {
	return &Builder{
		ms: ms,
	}
}
