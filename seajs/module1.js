// module1.js
define(function(require, exports, module) {
    var $ = require('jquery');
    var m2 = require('module2');
    var m3 = require('module3');
     
    exports.run = function() {
        return $.merge(['module1'], $.merge(m2.run(), m3.run()));    
    }
});