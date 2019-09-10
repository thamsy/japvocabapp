$(document).ready(function() {
    // Logout Button
    $('#logout-submit').mousedown(function (e) {
        document.cookie = "questions_wrong=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        window.location.href = '/logout'
    });

    // Show Wrong Questions
    let cookie = getCookie("questions_wrong")

    if (cookie != '') {
        let questions_wrong = JSON.parse(cookie);

        let html_str_questions_wrong =
            '<p>Questions Wrong (' + questions_wrong.length.toString() + '): </p>\n' +
            '<ul class="list-group">\n';
        questions_wrong.forEach(function (question, idx) {
            html_str_questions_wrong +=
                '<li class="list-group-item list-group-item-danger">' +
                question["Eng"] + ' : ' + question["Kana"] + " / " + question["Kanji"] + " / " + question["Grp"] +
                '</li>\n'
        });
        html_str_questions_wrong += '</ul>\n' +
            '<button type="button" class="btn btn-outline-success btn-block" id="clear">Clear</button>';

        $('.questions-wrong').html(html_str_questions_wrong)

        $('#clear').mousedown(function (e) {
            $('.questions-wrong').html("")
            document.cookie = "questions_wrong=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        })
    }
});

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}