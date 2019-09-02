package parser

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

var animationFrame = `<body><tr  name="cpp" class="memitem:a0234c1f2c02129634151a60625c8d13e"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>
<tr name="cpp"  class="memdesc:a0234c1f2c02129634151a60625c8d13e"><td class="mdescLeft">&#160;</td><td class="mdescRight">Return a SpriteFrameName to be used.  <a href="#a0234c1f2c02129634151a60625c8d13e">More...</a><br /></td></tr>
`
var content1 = `
<tr  name="cpp" class="memitem:a0234c1f2c02129634151a60625c8d13e"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>
<tr  name="cpp" class="memitem:ab1604dce45043d57edb79c1d48034270"><td class="memItemLeft" align="right" valign="top">virtual void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/de4/classcocos2d_1_1_layer.html#ab1604dce45043d57edb79c1d48034270">onKeyPressed</a> (<a class="el" href="../../dd/df2/group__base.html#ga7885f47644a0388f981f416fa20389b2">EventKeyboard::KeyCode</a> keyCode, <a class="el" href="../../d4/d7a/classcocos2d_1_1_event.html">Event</a> *event)</td></tr>
<tr  name="cpp" class="memitem:a30d06f9e4cd396f0311efa9b773d0cb3"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../d2/db1/classcocos2d_1_1ui_1_1_button.html#a30d06f9e4cd396f0311efa9b773d0cb3">loadTextures</a> (const std::string &amp;normal, const std::string &amp;selected, const std::string &amp;disabled=&quot;&quot;, <a class="el" href="../../dd/df1/group__ui.html#gada57b5cb7e28956b3793954a578f377c">TextureResType</a> texType=TextureResType::LOCAL)</td></tr>
`

func TestParseMemitem(t *testing.T) {
	all := anchorRe.ReplaceAll([]byte(content1), []byte(""))

	submatch := memRe.FindAllSubmatch(all, -1)

	for _, m := range submatch {
		log.Printf("%s", m[1])
		log.Printf("%s", m[3])
		log.Printf("%s", m[4])

		returnSubmatch := returnRe.FindAllSubmatch(m[1], -1)
		var buffer bytes.Buffer
		for _, n := range returnSubmatch {
			//url := string(n[1])
			buffer.Write(n[1]) //返回值前缀如const等
			buffer.Write(n[4]) //返回值名
			buffer.Write(n[5]) //返回值后缀如* &等
		}
		log.Printf("\t--%s", buffer.String())

		buffer.Reset()
		items := strings.Split(string(m[4]), ",")
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
