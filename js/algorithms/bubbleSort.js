var a = [1, 6, 2, 4, 9, 0, 5, 3, 7, 8];

console.log("Unsorted Array", a);
console.log("Sorted Array  ", bubbleSort(a));

function bubbleSort(a) {
    var swapped = true;
    var ai, ai1;
    while (swapped) {
        swapped = false;
        for (var i = 0; i < a.length; i++) {
            if (a[i+1] < a[i]) {
                ai  = a[i];
                ai1 = a[i+1];
                a[i]   = ai1;
                a[i+1] = ai;
                swapped = true;
            }
        }
    }
    return a;
}
