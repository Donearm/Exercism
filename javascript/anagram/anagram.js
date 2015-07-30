var anagram = function(str) {
	// sort letters in a string (lowercase)
	var sorted = function(w) {
		return w.toLowerCase().split("").sort().join('');
	};

	this.word = str;
	this.matches = function(words) {
		// make an array if arguments are string(s)
		if (typeof words === 'string') {
			words = [].slice.apply(arguments);
		}

		var sortedWord = sorted(this.word);
		var sameWord = [];

		for (var i = 0; i < words.length; i++) {
			// same exact word, case insensitively? Discard
			if (this.word.toLowerCase() === words[i].toLowerCase()) {
				break
			} else if (sortedWord === sorted(words[i])) {
					sameWord.push(words[i]);
			}
		}
		return sameWord;
	}
	return this;
};

module.exports = anagram;
