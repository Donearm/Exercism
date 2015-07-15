var words = function(w) {
	// trimming too to remove whitespaces at the beginning
	var splitw = w.trim().split(/\s+/);
	var result = Object.create(null);
	
	splitw.forEach(function (word) {
		if (result[word] === undefined) { // undefined = not there
			result[word] = 1;
		} else {
			result[word] += 1;
		}
	})
	return result;
};

module.exports = words;
