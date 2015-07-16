var compute = function(seqA, seqB) {
	if (seqA.length !== seqB.length) {
		throw new Error("DNA strands must be of equal length.");
	}
	var difference = 0;
	for (var i = 0; i < seqA.length; i++) {
		if (seqA.charAt(i) !== seqB.charAt(i)) {
			difference++;
		}
	}
	return difference;
};

module.exports.compute = compute;
