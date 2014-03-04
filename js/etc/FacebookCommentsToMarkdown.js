// Facebook Comments to Markdown.
// ===
// _By Daniel Rodríguez (@sadasant)_  
// To run it, copy the following code in the javascript console of the conversation's page.
// To go to the conversation's page, click on the little gray date below the original post.
// Remember to load all the previous comments (they appear hidden if there are too many).
Array.prototype.filter.call(document.getElementsByClassName("UFICommentActorName"), function(e) {
    return e.parentNode.nextSibling.firstChild.firstChild.firstChild;
}).map(function(e) {
    var date = e.parentNode.nextSibling.firstChild.firstChild.firstChild.attributes.title.value;
    var name = e.text;
    var content = e.nextSibling.nextSibling.firstChild.innerHTML.replace(/<[^>]*>/g,"");
    var links = e.nextSibling.nextSibling.firstChild.getElementsByTagName("a");
    for (var i = 0; i < links.length; i++) {
        content = content.replace(/https{0,1}:\/\/[^\s]*.../, "<"+links[i].attributes.href.value+">");
    }
    return [
        "_"+date+"_  ",
        "**"+name+":**  ",
        content
    ].join("\n");
}).join("\n\n---\n\n");
