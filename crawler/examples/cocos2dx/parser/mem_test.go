package parser

import "testing"

func TestParseMemitem(t *testing.T) {
	const content = `<tr  name="cpp" class="memitem:ac83ef78efb9d88c6b9217c1ccf7e4cac"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d2/d0c/classcocos2d_1_1_vec3.html">Vec3</a>&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#ac83ef78efb9d88c6b9217c1ccf7e4cac">vertex</a> (const <a class="el" href="../../d1/d9c/classcocos2d_1_1_vec2.html">Vec2</a> &amp;pos) const </td></tr>
	<tr  name="js" class="memitem:ac83ef78efb9d88c6b9217c1ccf7e4cac"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#ac83ef78efb9d88c6b9217c1ccf7e4cac">vertex</a> ( var pos)</td></tr>
	<tr  name="cpp" class="memitem:a6689779177452cdadf408d57cdeae9a8"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a6689779177452cdadf408d57cdeae9a8"></a>
	<a class="el" href="../../d2/d0c/classcocos2d_1_1_vec3.html">Vec3</a>&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a6689779177452cdadf408d57cdeae9a8">getOriginalVertex</a> (const <a class="el" href="../../d1/d9c/classcocos2d_1_1_vec2.html">Vec2</a> &amp;pos) const </td></tr>
	<tr name="cpp"  class="memdesc:a6689779177452cdadf408d57cdeae9a8"><td class="mdescLeft">&#160;</td><td class="mdescRight">Returns the original (non-transformed) vertex at a given position. <br /></td></tr>
	<tr name = "cpp" class="separator:a6689779177452cdadf408d57cdeae9a8"><td class="memSeparator" colspan="2">&#160;</td></tr>
	<tr  name="cpp" class="memitem:a60247649d9c66382b6c3136d961f6c75"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d2/d0c/classcocos2d_1_1_vec3.html">Vec3</a>&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a60247649d9c66382b6c3136d961f6c75">originalVertex</a> (const <a class="el" href="../../d1/d9c/classcocos2d_1_1_vec2.html">Vec2</a> &amp;pos) const </td></tr>
	<tr  name="cpp" class="memitem:a0175824c0a8a7484c25fdc4487ddff99"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a0175824c0a8a7484c25fdc4487ddff99"></a>
		void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a0175824c0a8a7484c25fdc4487ddff99">setVertex</a> (const <a class="el" href="../../d1/d9c/classcocos2d_1_1_vec2.html">Vec2</a> &amp;pos, const <a class="el" href="../../d2/d0c/classcocos2d_1_1_vec3.html">Vec3</a> &amp;<a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#ac83ef78efb9d88c6b9217c1ccf7e4cac">vertex</a>)</td></tr>
	<tr name="cpp"  class="memdesc:a0175824c0a8a7484c25fdc4487ddff99"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets a new vertex at a given position. <br /></td></tr>
	<tr name = "cpp" class="separator:a0175824c0a8a7484c25fdc4487ddff99"><td class="memSeparator" colspan="2">&#160;</td></tr>
	<tr><td colspan="2"><div class="groupHeader"></div></td></tr>
	<tr  name="cpp" class="memitem:a78acafe1de2c03799c530d083a3cac4f"><td class="memItemLeft" align="right" valign="top">virtual void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a78acafe1de2c03799c530d083a3cac4f">beforeBlit</a> () override</td></tr>
	<tr  name="js" class="memitem:a78acafe1de2c03799c530d083a3cac4f"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a78acafe1de2c03799c530d083a3cac4f">beforeBlit</a> ()</td></tr>
	<tr  name="lua" class="memitem:a78acafe1de2c03799c530d083a3cac4f"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a78acafe1de2c03799c530d083a3cac4f">beforeBlit</a> ()</td></tr>
	<tr  name="cpp" class="memitem:aec55b077a80c8321729f42a6cda622e3"><td class="memItemLeft" align="right" valign="top">virtual void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#aec55b077a80c8321729f42a6cda622e3">blit</a> () override</td></tr>
	<tr name="cpp"  class="memdesc:aec55b077a80c8321729f42a6cda622e3"><td class="mdescLeft">&#160;</td><td class="mdescRight">Interface used to blit the texture with grid to screen.  <a href="#aec55b077a80c8321729f42a6cda622e3">More...</a><br /></td></tr>
	<tr name = "cpp" class="separator:aec55b077a80c8321729f42a6cda622e3"><td class="memSeparator" colspan="2">&#160;</td></tr>
	<tr  name="lua" class="memitem:aec55b077a80c8321729f42a6cda622e3"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#aec55b077a80c8321729f42a6cda622e3">blit</a> ()</td></tr>
	<tr name="lua"  class="memdesc:aec55b077a80c8321729f42a6cda622e3"><td class="mdescLeft">&#160;</td><td class="mdescRight">Interface used to blit the texture with grid to screen.  <a href="#aec55b077a80c8321729f42a6cda622e3">More...</a><br /></td></tr>
	<tr name="lua" class="separator:aec55b077a80c8321729f42a6cda622e3"><td class="memSeparator" colspan="2">&#160;</td></tr>
	<tr  name="cpp" class="memitem:a912cd1d7f037753dfe2c50121d7d9eb8"><td class="memItemLeft" align="right" valign="top">virtual void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d59/classcocos2d_1_1_grid3_d.html#a912cd1d7f037753dfe2c50121d7d9eb8">reuse</a> () override</td></tr>
	<tr name="cpp"  class="memdesc:a912cd1d7f037753dfe2c50121d7d9eb8"><td class="mdescLeft">&#160;</td><td class="mdescRight">Interface, Reuse the grid vertices.  <a href="#a912cd1d7f037753dfe2c50121d7d9eb8">More...</a><br /></td></tr>
	<tr name = "cpp" class="separator:a912cd1d7f037753dfe2c50121d7d9eb8"><td class="memSeparator" colspan="2">&#160;</td></tr>`

	ParseMemitem([]byte(content), "")
}
