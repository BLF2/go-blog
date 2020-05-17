package cache

import (
	"github.com/coocood/freecache"
	"go-blog/constant"
	"log"
)

var CacheUtil *freecache.Cache

func init()  {
	log.Println("初始化缓存")
	CacheUtil = freecache.NewCache(constant.SESSION_EXPORE_SECONDS)
	if CacheUtil != nil{

		log.Println("初始化缓存成功")
	}else{
		log.Fatalln("初始化缓存失败")
	}
}
