// updateCSS.js
// ------------

// A JavaScript _plugin_
// to constantly update any style link
// on pages with jQuery

// Get your link in the line below,
// then copy all this code without comments
// to your JavaScript console

var link = $('link[href^="/stylesheets/style"]')

;(function (l,h,s){
h=l[0].href;l.remove()
s=$('<style/>')
$('head').append(s)
function get(){$.get(h,got)}
function got(D){s.html(D);setTimeout(get,3000)}
get()})(link)