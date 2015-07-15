var Bob = function() {};

Bob.prototype.hey = function(input) {
	if (input.lentgh === 0 || input.match(/^\s*$/)) {
		return 'Fine. Be that way!'
	} else if (input === input.toUpperCase() && /[A-Z]/.test(input)) {
		return 'Whoa, chill out!'
	} else if (input.charAt(input.length - 1) === '?') {
		return 'Sure.'
	} else {
		return 'Whatever.'
	}
};

module.exports = Bob;
