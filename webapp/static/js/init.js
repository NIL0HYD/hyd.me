// index init	
define(function(require, exports, module) {

	var $ = require("/static/js/jquery.sea.js");
	require("/static/js/bootstrap/bootstrap.js")($);
	//require("/static/css/bootstrap.min.css");
	
	$('.carousel').carousel({
        interval: 5000
    });

});