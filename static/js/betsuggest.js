window.chartColors = {
	red: 'rgb(255, 99, 132)',
	orange: 'rgb(255, 159, 64)',
	yellow: 'rgb(255, 205, 86)',
	green: 'rgb(75, 192, 192)',
	blue: 'rgb(54, 162, 235)',
	purple: 'rgb(153, 102, 255)',
	grey: 'rgb(201, 203, 207)'
};

window.randomScalingFactor = function() {
	return (Math.random() > 0.5 ? 1.0 : -1.0) * Math.round(Math.random() * 100);
};

var MONTHS = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
var config = {
    type: 'line',
    data: {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [{
            label: "帐户金额",
            backgroundColor: window.chartColors.red,
            borderColor: window.chartColors.red,
            data: [
                randomScalingFactor(),
                randomScalingFactor(),
                randomScalingFactor(),
                randomScalingFactor(),
                randomScalingFactor(),
                randomScalingFactor(),
                randomScalingFactor()
            ],
            fill: false,
        }]
    },
    options: {
        responsive: true,
        title:{
            display:true,
            text:'帐户金额历史波动'
        },
        tooltips: {
            mode: 'index',
            intersect: false,
        },
        hover: {
            mode: 'nearest',
            intersect: true
        },
        scales: {
            xAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Month'
                }
            }],
            yAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Value'
                }
            }]
        }
    }
};

function TransBetTypeToStr(betType) {
	switch (betType) {
	case 0:
		return "庄"
	case 1:
		return "闲"
	case 2:
		return "和"
	}
	return ""
}
var socket;

$(document).ready(function () {
    var ctx = document.getElementById("betamountChart").getContext("2d");
    window.myLine = new Chart(ctx, config);
    // Create a socket
    socket = new WebSocket('ws://localhost:8080/ws/join?uname=Test');
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        console.log(data);
        var ContentObj,msgStr;
        switch (data.Type) {
        case 3: // EVENT_SUGGESTION
            ContentObj=JSON.parse(data.Content)
            //SuggestionBetStr=TransBetTypeToStr(ContentObj.SuggestionBet)
            //msgStr="第 " +ContentObj.TableNo + " 桌 " + ContentObj.GameIDDisplay + " 下一局建議買 " + SuggestionBetStr + " (" + ContentObj.TrendName + ")"
            //$("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + msgStr + "</li>");
            createSuggestion(ContentObj)
            break;
        case 4: // RESULT
            ContentObj=JSON.parse(data.Content)
            var guessResultStr
			if (ContentObj.TieReturn) {
				guessResultStr = "平"

			} else {

				if (ContentObj.GuessResult ) {
					guessResultStr = "勝"
				} else {
					guessResultStr = "負"
				}

			}
			if (ContentObj.FirstHand) {
				guessResultStr = "第一局預測不記結果"
			}

			msgStr = "第 " + ContentObj.TableNo + " 桌 " + ContentObj.GameIDDisplay + " 開 " + TransBetTypeToStr(ContentObj.Result) + " 建議結果:" + guessResultStr
            //$("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + msgStr + "</li>");
            break;
        case 5: // EVENT_ACCOUNT
            //$("#ubalance").html(data.Content);
            break;
        case 6: // EVENT_BET
        var BetType="閒"
            
            ContentObj=JSON.parse(data.Content)
            /*
            if(ContentObj.Settled){
                $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + ContentObj.BetTime + " 第 "+ContentObj.TableNo+" 桌"+ContentObj.GameIDDisplay+" 開 "+ ContentObj.GameResultTypeStr+" "+ContentObj.WinAmmount+" 帳戶餘額:"+ContentObj.CurrentBalance+"</li>");
            }else{
                $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + ContentObj.BetTime + " 第 "+ContentObj.TableNo+" 桌"+ContentObj.GameIDDisplay+" 買 "+ ContentObj.BetTypeStr+" "+ContentObj.BetAmmount+" 帳戶餘額:"+ContentObj.CurrentBalance+"</li>");
            }
            */
            break;
        }
    };


});


function createSuggestion(ContentObj){
    var SuggestionBetStr=TransBetTypeToStr(ContentObj.SuggestionBet)
    var SuggestionBetColorStr="green-bg"
    if(ContentObj.SuggestionBet==0){
        SuggestionBetColorStr="red-bg"
    }
    if(ContentObj.SuggestionBet==1){
        SuggestionBetColorStr="blue-bg"
    }


    var domStr="<div class='col-lg-6 col-md-6 col-sm-12 col-xs-12'>";
    domStr=domStr+"<div class='info-box "+SuggestionBetColorStr+"'>";
    domStr=domStr+"<i class='fa'>"+SuggestionBetStr+"</i>";
    domStr=domStr+"<div class='title'>第<span>"+ContentObj.TableNo+"</span>桌建议押</div>";
    domStr=domStr+"<div class='count'>正确率：<span>40.36</span>%</div>";
    domStr=domStr+"<div class='count'>状态：<span>未跟注</span> (点击改变状态)</div>";
    domStr=domStr+"</div>";
    domStr=domStr+"</div>";
    $("#suggestionList div").first().before(domStr)

}


/*
        document.getElementById('randomizeData').addEventListener('click', function() {
            config.data.datasets.forEach(function(dataset) {
                dataset.data = dataset.data.map(function() {
                    return randomScalingFactor();
                });

            });

            window.myLine.update();
        });

        var colorNames = Object.keys(window.chartColors);
        document.getElementById('addDataset').addEventListener('click', function() {
            var colorName = colorNames[config.data.datasets.length % colorNames.length];
            var newColor = window.chartColors[colorName];
            var newDataset = {
                label: 'Dataset ' + config.data.datasets.length,
                backgroundColor: newColor,
                borderColor: newColor,
                data: [],
                fill: false
            };

            for (var index = 0; index < config.data.labels.length; ++index) {
                newDataset.data.push(randomScalingFactor());
            }

            config.data.datasets.push(newDataset);
            window.myLine.update();
        });

        document.getElementById('addData').addEventListener('click', function() {
            if (config.data.datasets.length > 0) {
                var month = MONTHS[config.data.labels.length % MONTHS.length];
                config.data.labels.push(month);

                config.data.datasets.forEach(function(dataset) {
                    dataset.data.push(randomScalingFactor());
                });

                window.myLine.update();
            }
        });

        document.getElementById('removeDataset').addEventListener('click', function() {
            config.data.datasets.splice(0, 1);
            window.myLine.update();
        });

        document.getElementById('removeData').addEventListener('click', function() {
            config.data.labels.splice(-1, 1); // remove the label first

            config.data.datasets.forEach(function(dataset, datasetIndex) {
                dataset.data.pop();
            });

            window.myLine.update();
        });
        */