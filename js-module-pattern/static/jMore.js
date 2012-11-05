/**
 * 博客文章
 * http://www.adequatelygood.com/2010/3/JavaScript-Module-Pattern-In-Depth
 * http://www.cnblogs.com/wbpmrck/archive/2012/02/20/2359591.html
 */
 
var jMore = (function (M) {
	M.count = 0;
	M.YES = "yes";

	M.info = {
		version: "0.0.1"
	};

	return M;
}(jMore || {}));


// Module Export
var MODULE = (function (M) {
	var my = {},
		privateVariable = 1;
	
	function privateMethod() {
		// ...
	}
	
	my.moduleProperty = 1;
	my.moduleMethod = function () {
		alert("moduleMethod");
	};
	
	return my;
}(jMore || {}));

/* It's not work!
   注意加载顺序。加载顺序对的话，可以正常工作。
*/ 
var USE = (function(){
	$('#container').empty().html("<h1>USE jQuery!!</h1>");
}());


// Cross-File Private State
var CMODULE = (function (my) {
	var _private = my._private = my._private || {},
		_seal = my._seal = my._seal || function () {
			delete my._private;
			delete my._seal;
			delete my._unseal;
		},
		_unseal = my._unseal = my._unseal || function () {
			my._private = _private;
			my._seal = _seal;
			my._unseal = _unseal;
		};
	
	// permanent access to _private, _seal, and _unseal
	
	return my;
}(CMODULE || {}));



