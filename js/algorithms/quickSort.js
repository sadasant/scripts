var a = [1, 6, 2, 4, 9, 0, 5, 3, 7, 8];

console.log("Unsorted Array", a);
console.log("Sorted Array  ", quickSort(a));

function quickSort(a) {
    if (a.length < 2) return a;

    var left  = 0;
    var right = a.length - 1;
    var pivot = Math.floor(Math.random()*a.length);

    var a_pivot = a[pivot];
    a[pivot]    = a[right];
    a[right]    = a_pivot;

    for (var i = 0; i < a.length; i++) {
        if (a[i] < a[right]) {
            var a_i = a[i];
            a[i]    = a[left];
            a[left] = a_i;
            left++;
        }
    }

    a_left   = a[left];
    a[left]  = a[right];
    a[right] = a_left;

    return [].concat(quickSort(a.slice(0, left+1)), quickSort(a.slice(left+1)));
}
