package cache

import (
	"github.com/coocood/freecache"
	"go-blog/src/constant"
)

var CacheUtil *freecache.Cache

func init()  {
	CacheUtil = freecache.NewCache(constant.SESSION_EXPORE_SECONDS)
}
