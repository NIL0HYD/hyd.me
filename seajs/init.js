// init.js
define(function(require, exports, module) {
    var $ = require('jquery');
    var m1 = require('module1');
     
    exports.initPage = function() {
        $('.content').html(m1.run());    
    }
});