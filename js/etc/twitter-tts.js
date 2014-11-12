// Bookmark url (http):
// javascript:(function(){var d=document,s=d.createElement('script');s.src='http://sadasant.com/js/twitter-tts.js';d.body.appendChild(s);}())
// Bookmark url (https):
// javascript:(function(){var d=document,s=d.createElement('script');s.src='https://rawgit.com/sadasant/scripts/master/js/etc/twitter-tts.js';d.body.appendChild(s);}())
if (window.TWITTER_TTS) {
    alert("TWITTER TTS is already available :)");
} else {
    (function() {
        var TTS = window.TWITTER_TTS = {};
        TTS.last_tweets = 10;

        function listen(txt, done) {
            var audio = new Audio("http://tts-api.com/tts.mp3?q=" + escape(input));
            audio.play();
            audio.remove();
            if (done) audio.addEventListener("ended", done);
        }

        function tts() {
            var bar = $(".js-new-tweets-bar");
            if (!bar[0]) return tts_done();

            var new_tweets = bar.attr("data-item-count");
            listen("You have "+new_tweets+" new tweets.");

            bar.click();

            var tweets;

            function eachTweet() {
                var tweet = tweets.slice(0);
                var name  = $(tweet.previousSibling.previousSibling).find(".fullname");
                listen(name+" said: "+tweet.innerText, function() {
                    if (!tweets.length) return tts_done();
                    setTimeout(eachTweet, 1000);
                });
            }

            setTimeout(function() {
                listen("Listening to the last "+TTS.last_tweets+" tweets.");
                tweets = $('.js-tweet-text.tweet-text[lang="en"]').slice(0, TTS.last_tweets);
                listenTweets();
            }, 3000);

        }

        function tts_done() {
            setTimeout(tts, 30000); // 30 seconds
        }

        tts();
    })();
}
