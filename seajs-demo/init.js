//init.js
define(function(require, exports, module) {
    var $ = require('./jquery');
    var data = require('./data');
    var css = require('./style.css');
 
    $('.author').html(data.author);
    $('.blog').attr('href', data.blog);

    exports.tips = function() {
    	alert("Haha, got you!");
    }

});