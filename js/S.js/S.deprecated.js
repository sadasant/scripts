// sadasant.com/license

S.deprecated = function() {
  var D = document
  if (!(
     XMLHttpRequest
  && D.querySelectorAll
  )) {
    alert("Browser Deprecated")
  }
}
