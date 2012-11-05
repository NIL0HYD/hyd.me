/**
 * 学习jQuery
 * Window 对象表示一个浏览器窗口或一个框架。
 在客户端 JavaScript 中，Window 对象是全局对象，
 所有的表达式都在当前的环境中计算。也就是说，
 要引用当前窗口根本不需要特殊的语法，可以把那个窗口的属性作为全局变量来使用。
 例如，可以只写 document，而不必写 window.document。
 http://www.w3school.com.cn/htmldom/dom_obj_window.asp
 */
(function (window) {
// a lot of variables
var _H = window.H,
	_Y = window.Y,
	_D = window.D,
	H = function() {
		return new H.fn.init();
	};

H.fn = H.prototype = {
	init: function() {
		return function() {};
	}
};

H.fn.init.prototype = H.fn;

H.extend = H.fn.extend = function (map) {
	alert(map.key);
};

H.extend({key: "Wow!"});

H.fn.extend({key: "beard me!"});


// Expose H to the global object
window.H = window.Y = window.D = H;

//return H;

}(window));