var a = [1, 6, 2, 4, 9, 0, 5, 3, 7, 8];

console.log("Unsorted Array", a);
console.log("Sorted Array", mergeSort(a));
 
function mergeSort(A) {
    var l = A.length;
    if (l <= 1) return A;
    var a = [];
    var b = [];
    var half = Math.round(l/2);
    var i;
    for (i = 0; i < half; i++) {
        a.push(A[i]);
    }
    for (i = half; i < l; i++) {
        b.push(A[i]);
    }
    return merge(mergeSort(a), mergeSort(b));
}

function merge(a, b) {
    var i = 0;
    var j = 0;
    var l = a.length + b.length;
    var r = [];
    while (r.length !== l) {
        var ai = a[i];
        var bj = b[j];
        if ((bj !== 0 && !bj && ai) || ai <= bj) {
            r.push(ai);
            i++;
        } else
        if ((ai !== 0 && !ai && bj) || bj <= ai) {
            r.push(bj);
            j++;
        }
    }
    return r;
}
