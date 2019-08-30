package parser

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"testing"
)

var memRe2 = regexp.MustCompile(
	`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top">(<a class="anchor" id="[a-z0-9]+"></a>[\s]+)*([\S _]+)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)#[a-z0-9]+">([\S]+)</a>([\S _]+)</td></tr>`)
var content1 = `<tr  name="cpp" class="memitem:a0234c1f2c02129634151a60625c8d13e"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>
<tr name="cpp"  class="memdesc:a0234c1f2c02129634151a60625c8d13e"><td class="mdescLeft">&#160;</td><td class="mdescRight">Return a SpriteFrameName to be used.  <a href="#a0234c1f2c02129634151a60625c8d13e">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a0234c1f2c02129634151a60625c8d13e"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:acab7cd7323521a7d6a92de55e4a57bca"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acab7cd7323521a7d6a92de55e4a57bca">setSpriteFrame</a> (<a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *frame)</td></tr>
<tr name="cpp"  class="memdesc:acab7cd7323521a7d6a92de55e4a57bca"><td class="mdescLeft">&#160;</td><td class="mdescRight">Set the <a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html" title="A SpriteFrame has: ">SpriteFrame</a>.  <a href="#acab7cd7323521a7d6a92de55e4a57bca">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:acab7cd7323521a7d6a92de55e4a57bca"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="js" class="memitem:acab7cd7323521a7d6a92de55e4a57bca"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acab7cd7323521a7d6a92de55e4a57bca">setSpriteFrame</a> ( var frame)</td></tr>
<tr name="js"  class="memdesc:acab7cd7323521a7d6a92de55e4a57bca"><td class="mdescLeft">&#160;</td><td class="mdescRight">Set the <a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html" title="A SpriteFrame has: ">SpriteFrame</a>.  <a href="#acab7cd7323521a7d6a92de55e4a57bca">More...</a><br /></td></tr>
<tr name = "js" class="separator:acab7cd7323521a7d6a92de55e4a57bca"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:acab7cd7323521a7d6a92de55e4a57bca"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acab7cd7323521a7d6a92de55e4a57bca">setSpriteFrame</a> ( local frame)</td></tr>
<tr name="lua"  class="memdesc:acab7cd7323521a7d6a92de55e4a57bca"><td class="mdescLeft">&#160;</td><td class="mdescRight">Set the <a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html" title="A SpriteFrame has: ">SpriteFrame</a>.  <a href="#acab7cd7323521a7d6a92de55e4a57bca">More...</a><br /></td></tr>
<tr name="lua" class="separator:acab7cd7323521a7d6a92de55e4a57bca"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a46ec4ad1fef24093354d2509d578ef5c"><td class="memItemLeft" align="right" valign="top">float&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a46ec4ad1fef24093354d2509d578ef5c">getDelayUnits</a> () const </td></tr>
<tr name="cpp"  class="memdesc:a46ec4ad1fef24093354d2509d578ef5c"><td class="mdescLeft">&#160;</td><td class="mdescRight">Gets the units of time the frame takes.  <a href="#a46ec4ad1fef24093354d2509d578ef5c">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a46ec4ad1fef24093354d2509d578ef5c"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="js" class="memitem:a46ec4ad1fef24093354d2509d578ef5c"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a46ec4ad1fef24093354d2509d578ef5c">getDelayUnits</a> ()</td></tr>
<tr name="js"  class="memdesc:a46ec4ad1fef24093354d2509d578ef5c"><td class="mdescLeft">&#160;</td><td class="mdescRight">Gets the units of time the frame takes.  <a href="#a46ec4ad1fef24093354d2509d578ef5c">More...</a><br /></td></tr>
<tr name = "js" class="separator:a46ec4ad1fef24093354d2509d578ef5c"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a46ec4ad1fef24093354d2509d578ef5c"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a46ec4ad1fef24093354d2509d578ef5c">getDelayUnits</a> ()</td></tr>
<tr name="lua"  class="memdesc:a46ec4ad1fef24093354d2509d578ef5c"><td class="mdescLeft">&#160;</td><td class="mdescRight">Gets the units of time the frame takes.  <a href="#a46ec4ad1fef24093354d2509d578ef5c">More...</a><br /></td></tr>
<tr name="lua" class="separator:a46ec4ad1fef24093354d2509d578ef5c"><td class="memSeparator" colspan="2">&#160;</td></tr>`

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
