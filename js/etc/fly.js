// 1kFly.js ~ sadasant.com
// http://jsbin.com/aluyax/3
function fly(f,mx,my,gx,gy,tx,ty,b,X,Y,x,y,i,r,M,a,w,W,w1){
with(document){
var h=["☼❤☼","&nbsp;❤"],
P=M.PI*i,A=M.abs,
D=.7+Math.random()*.7,
p=createElement('p'),s,
o=onmousemove;(s=p.style).fontFamily="monospace";s.position="fixed"
body.appendChild(p)
onmousemove=function(e){mx=e.pageX;my=e.pageY;o(e)}
onmouseup=function(){fly(-f,mx,my,gx,gy,tx,ty,b,X,Y,x,y,i,r*D,M,a,w,W,w1)}}
setInterval(function(){
p.innerHTML=h[~~(b=!b)]
dx=gx-x;dy=gy-y
w=(A(dx/tx)+A(dy/ty))/2
if(!W&w<i){x=X;y=Y}
W=W||w<i
if(W){
a=x*f/P
y+=i;x+=i
w1=2
}else{
w1=1.1-w
a=(w*f*3*w1*D)/P;j=i*w1*D
w1*=r*D
y+=j*dy
x+=j*dx
}s.top=(Y=y+w1*M.cos(a))+'px'
s.left=(X=x+w1*M.sin(a))+'px'
if(W&A((mx-gx)+(my-gy))/2>r)
tx=(gx=mx)-x,ty=(gy=my)-y,f=M.random()<.5?-1:1,W=!W
},30)}fly(1,1,1,1,1,1,1,0,0,0,0,0,.05,30,Math)
