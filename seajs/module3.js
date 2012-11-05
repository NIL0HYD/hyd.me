// module3.js
define(function(require, exports, module) {
    var $ = require('jquery');
    var m4 = require('module4');
     
    exports.run = function() {
        return $.merge(['module3'], m4.run());    
    }
});