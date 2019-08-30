package parser

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"testing"
)

var memRe2 = regexp.MustCompile(
	`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top">([a-zA-Z0-9 :&;]*<a class="el" href="../../[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html[#a-z0-9]*">[a-zA-Z0-9 _]+</a>[a-zA-Z0-9 :&;*]*)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)[#a-z0-9]+">([a-zA-Z0-9_]+)</a>([a-zA-Z0-9 ()_]+)</td></tr>`)
var content1 = `<tr  name="cpp" class="memitem:a0234c1s"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>`

func TestParseMemitem(t *testing.T) {
	submatch := memRe2.FindAllSubmatch([]byte(content1), -1)

	for _, m := range submatch {
		log.Printf("%s", m[2])
		log.Printf("%s", m[4])
		log.Printf("%s", m[5])

		returnSubmatch := returnRe.FindAllSubmatch(m[2], -1)
		var buffer bytes.Buffer
		for _, n := range returnSubmatch {
			//url := string(n[1])
			buffer.Write(n[1]) //返回值前缀如const等
			buffer.Write(n[4]) //返回值名
			buffer.Write(n[5]) //返回值后缀如* &等
		}
		log.Printf("\t--%s", buffer.String())

		buffer.Reset()
		items := strings.Split(string(m[5]), ",")
		for idx, item := range items {
			if idx != 0 {
				buffer.Write([]byte(","))
			}
			paramSubmatch := paramRe.FindAllSubmatch([]byte(item), -1)
			for _, n := range paramSubmatch {
				//url := string(n[1])
				buffer.Write(n[2])
				buffer.Write(n[3])
			}
		}
		log.Printf("\t--%s", buffer.String())

	}
}
