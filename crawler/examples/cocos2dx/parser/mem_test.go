package parser

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

var animationFrame = `<body><tr  name="cpp" class="memitem:a0234c1f2c02129634151a60625c8d13e"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>
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
<tr name="lua" class="separator:a46ec4ad1fef24093354d2509d578ef5c"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acbdf4eb773ca2a3efdf150e18b4ff37d">setDelayUnits</a> (float delayUnits)</td></tr>
<tr name="cpp"  class="memdesc:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets the units of time the frame takes.  <a href="#acbdf4eb773ca2a3efdf150e18b4ff37d">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="js" class="memitem:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acbdf4eb773ca2a3efdf150e18b4ff37d">setDelayUnits</a> ( var delayUnits)</td></tr>
<tr name="js"  class="memdesc:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets the units of time the frame takes.  <a href="#acbdf4eb773ca2a3efdf150e18b4ff37d">More...</a><br /></td></tr>
<tr name = "js" class="separator:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#acbdf4eb773ca2a3efdf150e18b4ff37d">setDelayUnits</a> ( local delayUnits)</td></tr>
<tr name="lua"  class="memdesc:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets the units of time the frame takes.  <a href="#acbdf4eb773ca2a3efdf150e18b4ff37d">More...</a><br /></td></tr>
<tr name="lua" class="separator:acbdf4eb773ca2a3efdf150e18b4ff37d"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a86211458f67e89ac5bd684376648f742"><td class="memItemLeft" align="right" valign="top">const ValueMap &amp;&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a86211458f67e89ac5bd684376648f742">getUserInfo</a> () const </td></tr>
<tr name="cpp"  class="memdesc:a86211458f67e89ac5bd684376648f742"><td class="mdescLeft">&#160;</td><td class="mdescRight">Gets user information A AnimationFrameDisplayedNotification notification will be broadcast when the frame is displayed with this dictionary as UserInfo.  <a href="#a86211458f67e89ac5bd684376648f742">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a86211458f67e89ac5bd684376648f742"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a2a1d0cb5d911c1ba5411665e0ac6e125">setUserInfo</a> (const ValueMap &amp;userInfo)</td></tr>
<tr name="cpp"  class="memdesc:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets user information.  <a href="#a2a1d0cb5d911c1ba5411665e0ac6e125">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="js" class="memitem:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memItemLeft" align="right" valign="top">var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a2a1d0cb5d911c1ba5411665e0ac6e125">setUserInfo</a> ( var userInfo)</td></tr>
<tr name="js"  class="memdesc:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets user information.  <a href="#a2a1d0cb5d911c1ba5411665e0ac6e125">More...</a><br /></td></tr>
<tr name = "js" class="separator:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a2a1d0cb5d911c1ba5411665e0ac6e125">setUserInfo</a> ( local userInfo)</td></tr>
<tr name="lua"  class="memdesc:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="mdescLeft">&#160;</td><td class="mdescRight">Sets user information.  <a href="#a2a1d0cb5d911c1ba5411665e0ac6e125">More...</a><br /></td></tr>
<tr name="lua" class="separator:a2a1d0cb5d911c1ba5411665e0ac6e125"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a09f6d358b8b706d4cb62b56453d9e978"><td class="memItemLeft" align="right" valign="top">virtual <a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html">AnimationFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a09f6d358b8b706d4cb62b56453d9e978">clone</a> () const override</td></tr>
<tr name="cpp"  class="memdesc:a09f6d358b8b706d4cb62b56453d9e978"><td class="mdescLeft">&#160;</td><td class="mdescRight">Returns a copy of the <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html" title="Ref is used for reference count management. ">Ref</a>.  <a href="#a09f6d358b8b706d4cb62b56453d9e978">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a09f6d358b8b706d4cb62b56453d9e978"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr class="inherit_header pub_methods_classcocos2d_1_1_ref"><td colspan="2" onclick="javascript:toggleInherit('pub_methods_classcocos2d_1_1_ref')"><img src="../../closed.png" alt="-"/>&#160;Public Member Functions inherited from <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html">Ref</a></td></tr>
<tr  name="cpp" class="memitem:a24888ae1fe9df2d329c9b485807cb62b inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a24888ae1fe9df2d329c9b485807cb62b">retain</a> ()</td></tr>
<tr name="cpp"  class="memdesc:a24888ae1fe9df2d329c9b485807cb62b inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Retains the ownership.  <a href="#a24888ae1fe9df2d329c9b485807cb62b">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a24888ae1fe9df2d329c9b485807cb62b inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a23b477d0e2d399f75d585d154c346591">release</a> ()</td></tr>
<tr name="cpp"  class="memdesc:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Releases the ownership immediately.  <a href="#a23b477d0e2d399f75d585d154c346591">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a23b477d0e2d399f75d585d154c346591">release</a> ()</td></tr>
<tr name="lua"  class="memdesc:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Releases the ownership immediately.  <a href="#a23b477d0e2d399f75d585d154c346591">More...</a><br /></td></tr>
<tr name="lua" class="separator:a23b477d0e2d399f75d585d154c346591 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a8a2345cdb9d96629de5dbaf88e36d505 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html">Ref</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a8a2345cdb9d96629de5dbaf88e36d505">autorelease</a> ()</td></tr>
<tr name="cpp"  class="memdesc:a8a2345cdb9d96629de5dbaf88e36d505 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Releases the ownership sometime soon automatically.  <a href="#a8a2345cdb9d96629de5dbaf88e36d505">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a8a2345cdb9d96629de5dbaf88e36d505 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">unsigned int&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#ac0fad74f624629043739b7c06c635a48">getReferenceCount</a> () const </td></tr>
<tr name="cpp"  class="memdesc:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Returns the <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html" title="Ref is used for reference count management. ">Ref</a>'s current reference count.  <a href="#ac0fad74f624629043739b7c06c635a48">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#ac0fad74f624629043739b7c06c635a48">getReferenceCount</a> ()</td></tr>
<tr name="lua"  class="memdesc:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Returns the <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html" title="Ref is used for reference count management. ">Ref</a>'s current reference count.  <a href="#ac0fad74f624629043739b7c06c635a48">More...</a><br /></td></tr>
<tr name="lua" class="separator:ac0fad74f624629043739b7c06c635a48 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a2fca001f1d40ace24096674c50e783d5 inherit pub_methods_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top">virtual&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a2fca001f1d40ace24096674c50e783d5">~Ref</a> ()</td></tr>
<tr name="cpp"  class="memdesc:a2fca001f1d40ace24096674c50e783d5 inherit pub_methods_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Destructor.  <a href="#a2fca001f1d40ace24096674c50e783d5">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a2fca001f1d40ace24096674c50e783d5 inherit pub_methods_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr class="inherit_header pub_methods_classcocos2d_1_1_clonable"><td colspan="2" onclick="javascript:toggleInherit('pub_methods_classcocos2d_1_1_clonable')"><img src="../../closed.png" alt="-"/>&#160;Public Member Functions inherited from <a class="el" href="../../db/d5f/classcocos2d_1_1_clonable.html">Clonable</a></td></tr>
<tr  name="cpp" class="memitem:a38eea1a38ff1112358bd8def337bc1de inherit pub_methods_classcocos2d_1_1_clonable"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html">Ref</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../db/d5f/classcocos2d_1_1_clonable.html#a38eea1a38ff1112358bd8def337bc1de">copy</a> () const </td></tr>
<tr name="cpp"  class="memdesc:a38eea1a38ff1112358bd8def337bc1de inherit pub_methods_classcocos2d_1_1_clonable"><td class="mdescLeft">&#160;</td><td class="mdescRight">Returns a copy of the <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html" title="Ref is used for reference count management. ">Ref</a>.  <a href="#a38eea1a38ff1112358bd8def337bc1de">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:a38eea1a38ff1112358bd8def337bc1de inherit pub_methods_classcocos2d_1_1_clonable"><td class="memSeparator" colspan="2">&#160;</td></tr>
</table><table class="memberdecls">
<tr class="heading"><td colspan="2"><h2 class="groupheader"><a name="pub-static-methods"></a>
Static Public Member Functions</h2></td></tr>
<tr  name="cpp" class="memitem:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="memItemLeft" align="right" valign="top">static <a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html">AnimationFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#ab3abccbf8771c7b7bd4dc6f69f91e652">create</a> (<a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *spriteFrame, float delayUnits, const ValueMap &amp;userInfo)</td></tr>
<tr name="cpp"  class="memdesc:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="mdescLeft">&#160;</td><td class="mdescRight">Creates the animation frame with a spriteframe, number of delay units and a notification user info.  <a href="#ab3abccbf8771c7b7bd4dc6f69f91e652">More...</a><br /></td></tr>
<tr name = "cpp" class="separator:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="memItemLeft" align="right" valign="top">local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#ab3abccbf8771c7b7bd4dc6f69f91e652">create</a> ( local spriteFrame, local delayUnits, local userInfo)</td></tr>
<tr name="lua"  class="memdesc:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="mdescLeft">&#160;</td><td class="mdescRight">Creates the animation frame with a spriteframe, number of delay units and a notification user info.  <a href="#ab3abccbf8771c7b7bd4dc6f69f91e652">More...</a><br /></td></tr>
<tr name="lua" class="separator:ab3abccbf8771c7b7bd4dc6f69f91e652"><td class="memSeparator" colspan="2">&#160;</td></tr>
</table><table class="memberdecls">
<tr class="heading"><td colspan="2"><h2 class="groupheader"><a name="pro-methods"></a>
Protected Member Functions</h2></td></tr>
<tr  name="cpp" class="memitem:a553c51273118a192dd8c5b8ab7afc730"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a553c51273118a192dd8c5b8ab7afc730"></a>
bool&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a553c51273118a192dd8c5b8ab7afc730">initWithSpriteFrame</a> (<a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *spriteFrame, float delayUnits, const ValueMap &amp;userInfo)</td></tr>
<tr name="cpp"  class="memdesc:a553c51273118a192dd8c5b8ab7afc730"><td class="mdescLeft">&#160;</td><td class="mdescRight">initializes the animation frame with a spriteframe, number of delay units and a notification user info <br /></td></tr>
<tr name = "cpp" class="separator:a553c51273118a192dd8c5b8ab7afc730"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="js" class="memitem:a553c51273118a192dd8c5b8ab7afc730"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a553c51273118a192dd8c5b8ab7afc730"></a>
var&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a553c51273118a192dd8c5b8ab7afc730">initWithSpriteFrame</a> ( var spriteFrame, var delayUnits, var userInfo)</td></tr>
<tr name="js"  class="memdesc:a553c51273118a192dd8c5b8ab7afc730"><td class="mdescLeft">&#160;</td><td class="mdescRight">initializes the animation frame with a spriteframe, number of delay units and a notification user info <br /></td></tr>
<tr name = "js" class="separator:a553c51273118a192dd8c5b8ab7afc730"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a553c51273118a192dd8c5b8ab7afc730"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a553c51273118a192dd8c5b8ab7afc730"></a>
local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a553c51273118a192dd8c5b8ab7afc730">initWithSpriteFrame</a> ( local spriteFrame, local delayUnits, local userInfo)</td></tr>
<tr name="lua"  class="memdesc:a553c51273118a192dd8c5b8ab7afc730"><td class="mdescLeft">&#160;</td><td class="mdescRight">initializes the animation frame with a spriteframe, number of delay units and a notification user info <br /></td></tr>
<tr name="lua" class="separator:a553c51273118a192dd8c5b8ab7afc730"><td class="memSeparator" colspan="2">&#160;</td></tr>
</table><table class="memberdecls">
<tr class="heading"><td colspan="2"><h2 class="groupheader"><a name="inherited"></a>
Additional Inherited Members</h2></td></tr>
<tr class="inherit_header pub_attribs_classcocos2d_1_1_ref"><td colspan="2" onclick="javascript:toggleInherit('pub_attribs_classcocos2d_1_1_ref')"><img src="../../closed.png" alt="-"/>&#160;Public Attributes inherited from <a class="el" href="../../df/d28/classcocos2d_1_1_ref.html">Ref</a></td></tr>
<tr  name="cpp" class="memitem:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a87199ccfa9acfd9dbbe4a61be4f3f15d"></a>
unsigned int&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a87199ccfa9acfd9dbbe4a61be4f3f15d">_ID</a></td></tr>
<tr name="cpp"  class="memdesc:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">object id, ScriptSupport need public _ID <br /></td></tr>
<tr name = "cpp" class="separator:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a87199ccfa9acfd9dbbe4a61be4f3f15d"></a>
local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a87199ccfa9acfd9dbbe4a61be4f3f15d">_ID</a></td></tr>
<tr name="lua"  class="memdesc:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">object id, ScriptSupport need public _ID <br /></td></tr>
<tr name="lua" class="separator:a87199ccfa9acfd9dbbe4a61be4f3f15d inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="adb10edfc88a28e6dc62942b7f532edd0"></a>
int&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#adb10edfc88a28e6dc62942b7f532edd0">_luaID</a></td></tr>
<tr name="cpp"  class="memdesc:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Lua reference id. <br /></td></tr>
<tr name = "cpp" class="separator:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="adb10edfc88a28e6dc62942b7f532edd0"></a>
local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#adb10edfc88a28e6dc62942b7f532edd0">_luaID</a></td></tr>
<tr name="lua"  class="memdesc:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">Lua reference id. <br /></td></tr>
<tr name="lua" class="separator:adb10edfc88a28e6dc62942b7f532edd0 inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a4397f03c9723ba32b10cfb2d9f162cfe"></a>
void *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a4397f03c9723ba32b10cfb2d9f162cfe">_scriptObject</a></td></tr>
<tr name="cpp"  class="memdesc:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">scriptObject, support for swift <br /></td></tr>
<tr name = "cpp" class="separator:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="a4397f03c9723ba32b10cfb2d9f162cfe"></a>
local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#a4397f03c9723ba32b10cfb2d9f162cfe">_scriptObject</a></td></tr>
<tr name="lua"  class="memdesc:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">scriptObject, support for swift <br /></td></tr>
<tr name="lua" class="separator:a4397f03c9723ba32b10cfb2d9f162cfe inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="cpp" class="memitem:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="ab95879aeaa9d8b8255396bbc8cf9357b"></a>
bool&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#ab95879aeaa9d8b8255396bbc8cf9357b">_rooted</a></td></tr>
<tr name="cpp"  class="memdesc:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">When true, it means that the object was already rooted. <br /></td></tr>
<tr name = "cpp" class="separator:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
<tr  name="lua" class="memitem:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="ab95879aeaa9d8b8255396bbc8cf9357b"></a>
local&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d28/classcocos2d_1_1_ref.html#ab95879aeaa9d8b8255396bbc8cf9357b">_rooted</a></td></tr>
<tr name="lua"  class="memdesc:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="mdescLeft">&#160;</td><td class="mdescRight">When true, it means that the object was already rooted. <br /></td></tr>
<tr name="lua" class="separator:ab95879aeaa9d8b8255396bbc8cf9357b inherit pub_attribs_classcocos2d_1_1_ref"><td class="memSeparator" colspan="2">&#160;</td></tr>
</table>
</html>
`
var content1 = `
<tr  name="cpp" class="memitem:a0234c1f2c02129634151a60625c8d13e"><td class="memItemLeft" align="right" valign="top"><a class="el" href="../../d3/d35/classcocos2d_1_1_sprite_frame.html">SpriteFrame</a> *&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/d7d/classcocos2d_1_1_animation_frame.html#a0234c1f2c02129634151a60625c8d13e">getSpriteFrame</a> () const </td></tr>
<tr  name="cpp" class="memitem:ab1604dce45043d57edb79c1d48034270"><td class="memItemLeft" align="right" valign="top">virtual void&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../../df/de4/classcocos2d_1_1_layer.html#ab1604dce45043d57edb79c1d48034270">onKeyPressed</a> (<a class="el" href="../../dd/df2/group__base.html#ga7885f47644a0388f981f416fa20389b2">EventKeyboard::KeyCode</a> keyCode, <a class="el" href="../../d4/d7a/classcocos2d_1_1_event.html">Event</a> *event)</td></tr>
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
