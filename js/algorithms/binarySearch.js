function binarySearch(items, value){
  var startIndex = 0
    , stopIndex = items.length - 1
    , middle = Math.floor((stopIndex + startIndex)/2)

  while(items[middle] != value && startIndex < stopIndex){
    //adjust search area
    if (value < items[middle]){
      stopIndex = middle - 1
    } else if (value > items[middle]){
      startIndex = middle + 1
    }
    //recalculate middle
    middle = Math.floor((stopIndex + startIndex)/2)
    console.debug(middle)
  }

  //make sure it's the right value
  return (items[middle] != value) ? -1 : middle
}

var vector = [1,2,3,4,5,6,7]
  , num = 2

document.write(binarySearch(vector, num))
