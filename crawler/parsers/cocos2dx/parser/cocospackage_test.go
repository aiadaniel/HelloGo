package parser

import (
	"log"
	"regexp"
	"testing"
)

var content = `<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">class &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../d2/d3e/classcocos2d_1_1_text_field_delegate.html">TextFieldDelegate</a></td></tr>
<tr  class="memdesc:"><td class="mdescLeft">&#160;</td><td class="mdescRight">A input protocol for TextField.  <a href="../../d2/d3e/classcocos2d_1_1_text_field_delegate.html#details">More...</a><br /></td></tr>
<tr class="separator:"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">class &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../dc/d95/classcocos2d_1_1_text_field_t_t_f.html">TextFieldTTF</a></td></tr>
<tr  class="memdesc:"><td class="mdescLeft">&#160;</td><td class="mdescRight">A simple text input field with TTF font.  <a href="../../dc/d95/classcocos2d_1_1_text_field_t_t_f.html#details">More...</a><br /></td></tr>
<tr class="separator:"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">class &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../d0/d2f/classcocos2d_1_1ui_1_1_abstract_check_button.html">AbstractCheckButton</a></td></tr>
<tr  class="memdesc:"><td class="mdescLeft">&#160;</td><td class="mdescRight"><a class="el" href="../../d0/d2f/classcocos2d_1_1ui_1_1_abstract_check_button.html" title="AbstractCheckButton is a specific type of two-states button that can be either checked or unchecked...">AbstractCheckButton</a> is a specific type of two-states button that can be either checked or unchecked.  <a href="../../d0/d2f/classcocos2d_1_1ui_1_1_abstract_check_button.html#details">More...</a><br /></td></tr>
<tr class="separator:"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">class &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../d2/db1/classcocos2d_1_1ui_1_1_button.html">Button</a></td></tr>
<tr  class="memdesc:"><td class="mdescLeft">&#160;</td><td class="mdescRight">Represents a push-button widget.  <a href="../../d2/db1/classcocos2d_1_1ui_1_1_button.html#details">More...</a><br /></td></tr>`

var testRe = regexp.MustCompile(
	`<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">([a-zA-Z]+) &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html)">([a-zA-Z0-9]+)</a></td></tr>`)

func TestRe(t *testing.T) {
	submatch := testRe.FindAllSubmatch([]byte(content), -1)
	for _, m := range submatch {
		log.Printf("%s", m[3])
		log.Printf("%s", m[2])
		log.Printf("%s", m[1])
	}
}
