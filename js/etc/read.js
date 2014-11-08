// Bookmark url:
// javascript:(function(){var d=document,s=d.createElement('script');s.src='http://sadasant.com/js/read.js';d.body.appendChild(s);}())
if (window.READING) {
    alert("READ is already available :)");
} else {
    (function() {
        var READING = window.READING = {};
        READING.delay   = 150;
        READING.started = false;
        READING.stop    = false;
        READING.next    = 0;

        // Geting the slected text
        function getSelected() {
            var text = "";
            var selection;
            if (window.getSelection) {
                selection = window.getSelection();
                text = selection.toString();
            } else if (document.selection && document.selection.type != "Control") {
                selection = document.selection.createRange();
                text = selection.text;
            }
            return {
                text: text,
                selection: selection
            };
        }

        // Scrolling to the given element
        function scrollTo(element) {
            element.scrollIntoView();
            window.scrollTo(window.scrollX, window.scrollY - 100);
        }

        // Selecting the paragraph where the current selection is located
        function selectTextIn(element) {
            var range;
            if (document.body.createTextRange) {
                range = document.body.createTextRange();
                range.moveToElementText(element);
                range.select();
            } else {
                range = document.createRange();
                range.selectNodeContents(element);
                var selection = window.getSelection();
                selection.removeAllRanges();
                selection.addRange(range);
            }
        }

        // Creating the reading widget
        function createWidget() {
            var div = document.createElement('div');
            var txt = document.createElement('div');

            div.style.position   = "fixed";
            div.style.textAlign  = "center";
            div.style.top        = "0px";
            div.style.left       = "0px";
            div.style.width      = "100%";
            div.style.height     = "100%";
            div.style.background = "rgba(0,0,0,.9)";
            div.style.zIndex     = "999999999";

            txt.style.color      = "white";
            txt.style.marginTop  = "100px";
            txt.style.fontSize   = "32px";
            txt.style.fontFamily = "arial";

            div.appendChild(txt);
            document.body.appendChild(div);

            return {
                div: div,
                txt: txt
            };
        }

        // Read creates the widget and starts reading the selected text
        function read() {
            if (READING.started) {
                READING.stop = true;
                return;
            }

            var text = getSelected().text;
            if (!text) return;

            function prepare() {
                var words = [""];
                var valid = /[\(\)\[\]\{\}¿?¡!""''“”‘’0-9a-zA-Z\u00C0-\u017F]/; // The range \u00C0-\u017F is for accents.
                text.split("").forEach(function(char, i, chars) {
                    words[words.length - 1] += char;
                    if (!char.match(valid)) {
                        words[words.length - 1].trim();
                        words.push("");
                    }
                    if (i === chars.length - 1) {
                        start(words);
                    }
                });
            }

            function start(words) {
                READING.started = true;

                var w = createWidget();

                console.log("READING", words);

                function next() {
                    if (READING.stop || !words.length) {
                        READING.stop = false;
                        READING.started = false;
                        w.div.remove();
                        console.log("READING STOPPED");
                        return;
                    }

                    var word = words.shift();
                    w.txt.innerText = word;

                    // Delaying the next call
                    // because of certain characters
                    // or word length
                    var PLUS = 0;
                    if (word) {
                        if (word[word.length-1].match(/[,.:?!]/)) {
                            PLUS = 25;
                        }
                        if (word.length > 10) {
                            PLUS = Math.floor(word.length/10) * 100;
                        }
                    }

                    setTimeout(next, READING.delay + PLUS);
                }

                setTimeout(next, READING.delay);
            }

            console.log("READING", text);
            prepare();
        }

        // Catching keys
        function catchKeys() {
            var keydown;

            if (window.onkeydown) {
                keydown = window.onkeydown;
            }

            window.onkeydown = function(e) {
                if (e.ctrlKey) return keydown && keydown();
                // shift+r: start reading
                if (e.shiftKey && e.keyCode === 82) return read();
                // ESC: stop reading
                if (READING.started && e.keyCode === 27) {
                    READING.stop = true;
                }
                var p = getSelected().selection.anchorNode.parentNode;
                // shift+s: Select the whole paragraph.
                if (e.shiftKey && e.keyCode === 83) {
                    selectTextIn(p);
                }
                // shift+n: Select the next paragraph.
                if (e.shiftKey && e.keyCode === 78) {
                    p = p.nextSibling;
                    while(p.nodeType != 1) {
                        p = p.nextSibling;
                    }
                    scrollTo(p);
                    selectTextIn(p);
                }
                // shift+p: Select the previous paragraph.
                if (e.shiftKey && e.keyCode === 80) {
                    p = p.previousSibling;
                    while(p.nodeType != 1) {
                        p = p.previousSibling;
                    }
                    scrollTo(p);
                    selectTextIn(p);
                }
                // shift+d: Change the delay until the next word.
                if (e.shiftKey && e.keyCode === 68) {
                    READING.delay = parseInt(prompt("Delay until the next word:", READING.delay), 10);
                }
                if (keydown) keydown();
            };
        }

        catchKeys();
        alert([
            "READ: Now reading!",
            "Available commands:",
            "shift+r: Start reading.",
            "ESC: Stop reading.",
            "shift+s: Select the whole paragraph.",
            "shift+n: Select the next paragraph.",
            "shift+p: Select the previous paragraph.",
            "shift+d: Change the delay until the next word."
        ].join("\n"));
    })();
}
