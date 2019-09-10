window.onload = function() {

    document.ontouchmove = function(e){ e.preventDefault(); }

    var canvas  = document.getElementById('kanjiDrawId');
    var context = canvas.getContext("2d");
    canvas.width = $('#kanaTextId').outerWidth(); //todo: make responsive to orientation change

    var lastx;
    var lasty;

    context.strokeStyle = "#000000";
    context.lineCap = 'round';
    context.lineJoin = 'round';
    context.lineWidth = 4;

    function clear() {
        context.fillStyle = "#eeeeee";
        context.rect(0, 0, canvas.width, canvas.height);
        context.fill();
    }

    function dot(x,y) {
        context.beginPath();
        context.fillStyle = "#000000";
        context.arc(x,y,1,0,Math.PI*2,true);
        context.fill();
        context.stroke();
        context.closePath();
    }

    function line(fromx,fromy, tox,toy) {
        context.beginPath();
        context.moveTo(fromx, fromy);
        context.lineTo(tox, toy);
        context.stroke();
        context.closePath();
    }

    canvas.ontouchstart = function(event){
        event.preventDefault();

        var rect = canvas.getBoundingClientRect();
        lastx = event.touches[0].pageX - rect.left;
        lasty = event.touches[0].pageY - rect.top;

        dot(lastx,lasty);
    }

    canvas.ontouchmove = function(event){
        event.preventDefault();

        var rect = canvas.getBoundingClientRect();
        var newx = event.touches[0].pageX - rect.left;
        var newy = event.touches[0].pageY - rect.top;

        line(lastx,lasty, newx,newy);

        lastx = newx;
        lasty = newy;
    }


    var clearButton = document.getElementById('kanjiDrawClear')
    clearButton.onclick = clear

    clear()
}

// // Canvas Code
// context = document.getElementById('kanjiDrawId').getContext("2d");
// canvas = $('#kanjiDrawId');
//
// // Set canvas dimensions
// // canvas.width(window.innerWidth);
// // canvas.height(220);
//
// var clickX = new Array();
// var clickY = new Array();
// var clickDrag = new Array();
// var paint;
//
// // canvas.mousedown(mouseDownOrTouchStart);
// canvas.on('vmousedown', function(e){
//     var mouseX = e.pageX - this.offsetLeft;
//     var mouseY = e.pageY - this.offsetTop;
//
//     paint = true;
//     addClick(e.pageX - this.offsetLeft, e.pageY - this.offsetTop);
//     redraw();
// }
// );
//
// canvas.on('vmousemove', function(e){
//     if(paint){
//         addClick(e.pageX - this.offsetLeft, e.pageY - this.offsetTop, true);
//         redraw();
//     }
// });
//
// canvas.on('vmouseup', function(e){
//     paint = false;
// });
//
// canvas.on('vmouseleave', function(e){
//     paint = false;
// });
//
// function addClick(x, y, dragging)
// {
//     clickX.push(x);
//     clickY.push(y);
//     clickDrag.push(dragging);
// }
//
// function redraw(){
//     context.clearRect(0, 0, context.canvas.width, context.canvas.height); // Clears the canvas
//
//     context.strokeStyle = "#808080";
//     context.lineJoin = "round";
//     context.lineWidth = 3;
//
//     for(var i=0; i < clickX.length; i++) {
//         context.beginPath();
//         if(clickDrag[i] && i){
//             context.moveTo(clickX[i-1], clickY[i-1]);
//         }else{
//             context.moveTo(clickX[i]-1, clickY[i]);
//         }
//         context.lineTo(clickX[i], clickY[i]);
//         context.closePath();
//         context.stroke();
//     }
// }
//
// $('#kanjiDrawClear').mousedown(function(e)
// {
//     clickX = new Array();
//     clickY = new Array();
//     clickDrag = new Array();
//     redraw();
// });