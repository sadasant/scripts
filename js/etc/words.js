var str = process.argv[2];

if (!str) {
    console.log("\nGiven a string of concatenated words,");
    console.log("this program separates words with spaces based on a dictionary.");
    console.log("It should work with < O(n log(n))");
    console.log("\nTry it with: node words mycustomword\nThat would be:");
    str = "mycustomword";
}

var words  = {};
var words_array = require("fs").readFileSync('/usr/share/dict/words').toString().split("\n");
var wat = false;
for (var i = 0; i < words_array.length; i++) {
    if (!words_array[i]) continue;
    // Filtering non-vowels words with one character
    if (words_array[i].length === 1 && "aeiou".indexOf(words_array[i]) === -1) {
        continue;
    }
    // Filtering words composed only by two non-vowels.
    if (!words_array[i].match(/[aeiouy]/)) {
        continue;
    }
    words[words_array[i]] = true;
}

function separeWords(str) {
    var separated_words = "";
    function checkWord(str) {
        // console.log("checkWord", str);
        var word = "";
        for (var i = 0; i < str.length; i++) {
            word += str[i];
            // if (words[word]) console.log("is word", word);
            if (words[word] && (word.length == str.length || checkWord(str.substr(i+1)))) {
                if (word.length == str.length -1) {
                    separated_words += word;
                } else {
                    separated_words = word + " " + separated_words;
                }
                // console.log("isword", word);
                return true;
            }
        }
        // separated_words += word;
        return false;
    }
    checkWord(str);
    return separated_words;
}

console.log(separeWords(str));
