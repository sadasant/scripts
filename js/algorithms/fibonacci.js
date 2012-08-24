/* By Daniel R. (sadasant.com) */
(function f(l,a,b){
  a=a||0
  b=b||1
  console.log(b)
  --l&&f(l,b,a+b)
})(8)
