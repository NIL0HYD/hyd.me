/**
 * To write a jQuery plugin
 * Note:
 * 1. In the immediate scope of the plugin function, 
 * the this keyword refers to the jQuery object the plugin was invoked on.
 * 金额小写-->大写转换函数
 */
(function($) {
	var methods = {
		Uppercase: function(num) {
			var strOutput = "";
			var strUnit = '仟佰拾亿仟佰拾万仟佰拾元角分';
			var numUnit = '零壹贰叁肆伍陆柒捌玖';
			num += "00";
			var intPos = num.indexOf('.');
			if (intPos >= 0) {
				num = num.substring(0, intPos) + num.substr(intPos + 1, 2);
			}
			strUnit = strUnit.substr(strUnit.length - num.length);
		
			for ( var i = 0; i < num.length; i++) {
				strOutput += numUnit.substr(num.substr(i, 1), 1) + strUnit.substr(i, 1);
			}
			return strOutput.replace(/零角零分$/, '整') // 匹配任何以"零角零分"为结尾的字符，替换为"整"
							.replace(/零[仟佰拾]/g, '零') // 查找给定集合内的任何字符[仟佰拾]，全部替换为"零"
							.replace(/零{2,}/g, '零') // 匹配包含至少 2 个"零"的序列的字符串，全部替换为"零"
							.replace(/零([亿|万])/g, '$1') // 小括号允许重复多个字符，替换为第一个匹配的字符
							.replace(/零+元/, '元') // 匹配任何包含至少一个"零"的字符串，替换为"元"
							.replace(/亿零{0,3}万/, '亿') // 匹配包含 0 或3 个"零"的序列的字符串，替换为"亿"
							.replace(/^元/, "零元"); // 匹配任何开头为"元"的字符串
		}
	};
	
	$.fn.convertRMBUppercase = function(elOutPut) {
		var numValue = this.val();
		$("#" + elOutPut).val(methods.Uppercase(numValue));
		return this;
	};
})(jQuery);