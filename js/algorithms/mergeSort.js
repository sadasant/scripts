// by Daniel R. http://sadasant.com/

function mergeSort(M) {
  var l = M.length;
  if (l <= 1) return M;
  var a = [],
      b = [],
      h = Math.round(l/2),
      i;
  for (i = 0; i < h; i++)
      a[a.length] = M[i];
  for (i = h; i < l; i++)
      b[b.length] = M[i];
  a = mergeSort(a);
  b = mergeSort(b);
  return merge(a, b);
}

function merge(a, b) {
  var i = 0, j = 0,
      l = a.length + b.length,
      r = [];
  while (r.length !== l) {
    var _a = a[i],
        _b = b[j];
    if (!_b && _a || _a <= _b) {
      r[r.length] = _a;
      i++;
    } else
    if (!_a && _b || _b <= _a) {
      r[r.length] = _b;
      j++;
    }
  }
  return r;
}

console.log(mergeSort([1,0.5,6,8,123,3,4,2]));
