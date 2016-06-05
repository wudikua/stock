var stock = {
	trade : function(code, name, days, trades) {
		option = {
		    title : {
		        text: code + ' [' + name + '] '+ '日线图'
		    },
		    tooltip : {
		        trigger: 'axis',
		        showDelay: 2, 
		        formatter: function (params) {
		            var res = params[0].seriesName + ' ' + params[0].name;
		            res += '<br/>  开盘 : ' + params[0].value[0] + '  最高 : ' + params[0].value[3];
		            res += '<br/>  收盘 : ' + params[0].value[1] + '  最低 : ' + params[0].value[2];
		            return res;
		        }
		    },
		    legend: {
		        data:['日线图', '成交量']
		    },
		    toolbox: {
		        show : true,
		        feature : {
		            mark : {show: true},
		            dataZoom : {show: true},
		            dataView : {show: true, readOnly: false},
		            magicType: {show: true, type: ['line', 'bar']},
		            restore : {show: true},
		            saveAsImage : {show: true}
		        }
		    },
		    dataZoom : {
		        show : true,
		        realtime: true,
		        start : 50,
		        end : 100
		    },
		    xAxis : [
		        {
		            type : 'category',
		            boundaryGap : true,
		            axisTick: {onGap:false},
		            splitLine: {show:false},
		            data : days
		        }
		    ],
		    yAxis : [
		        {
		            type : 'value',
		            scale:true,
		            boundaryGap: [0.01, 0.01]
		        }
		    ],
		    series : [
		        {
		            name:name,
		            type:'k',
		            data:trades
		        },
		        
		    ]
		};
		return option                   
	},

	volume : function(code, name, days, volume) {
		option = {
		    tooltip : {
		        trigger: 'axis',
		        showDelay: 2,            // 显示延迟，添加显示延迟可以避免频繁切换，单位ms
		        formatter: function (v) {
		        	return Math.round(v[0].value/10000) + ' 万'
		        },
		    },
		    legend: {
		        y : -30,
		        data:['日线图','成交量']
		    },
		    toolbox: {
		        show : false,
		        feature : {
		            mark : {show: true},
		            dataZoom : {show: true},
		            dataView : {show: true, readOnly: false},
		            magicType : {show: true, type: ['line', 'bar']},
		            restore : {show: true},
		            saveAsImage : {show: true}
		        }
		    },
		    dataZoom : {
		        show : false,
		        realtime: true,
		        start : 50,
		        end : 100
		    },
		    xAxis : [
		        {
		            type : 'category',
		            position:'bottom',
		            boundaryGap : true,
		            axisTick: {onGap:false},
		            splitLine: {show:false},
		            data : days
		        }
		    ],
		    yAxis : [
		        {
		            type : 'value',
		            scale:true,
		            splitNumber:3,
		            boundaryGap: [0.05, 0.05],
		            axisLabel: {
		                formatter: function (v) {
		                    return Math.round(v/10000) + ' 万'
		                }
		            },
		            splitArea : {show : true}
		        }
		    ],
		    series : [
		        {
		            name:'成交量',
		            type:'bar',
		            symbol: 'none',
		            data: volume,
		        }
		    ]
		};
		return option
	}
}