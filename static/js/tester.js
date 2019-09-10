$(document).ready(function() {
    // Delete wrong questions cookie
    document.cookie = "questions_wrong=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

    // Disable Enter Key
    $('form input').keydown(function(event){
        if(event.keyCode == 13) {
            event.preventDefault();
            return false;
        }
    });

    // Get Questions
    var urlParams = new URLSearchParams(window.location.search);
    $.getJSON("/vocab/" + urlParams.get("num"), function(data){
        questions = data;
        resetFields();
        loadQuestion();
        $(".black-screen").fadeOut('fast')
    });
});

let questions;
let question_num = 0;
let is_kana_corr = false;
let is_kanji_corr = false;
let is_grp_corr = false;
let questions_wrong = [];

function loadQuestion() {
    console.assert(question_num < questions.length)
    let question = questions[question_num]

    $("#eng").text(question["Eng"])

    $('.submit').mousedown(function (e) {
        is_kana_corr = $('.kana-text').val() == question["Kana"]
        if (is_kana_corr) {
            $('.kana-text').addClass("is-valid");
        } else {
            $('.kana-text').addClass("is-invalid");
        }
        $('.kana-text').prop("readonly", true);


        is_grp_corr = $('.grp-select').val() == question["Grp"]
        if (is_grp_corr) {
            $('.grp-select').addClass("is-valid");
        } else {
            $('.grp-select').addClass("is-invalid");
        }
        $('.grp-options').prop("disabled", true);

        $('.tick').prop("disabled", false);
        $('.cross').prop("disabled", false);
        $('.show').prop("disabled", false);
        $('.submit').prop("disabled", true);
        $('.clear').prop("disabled", true);
    });

    $('.show').mousedown(function (e) {
        kanji = question["Kanji"] != "" ? question["Kanji"] : "None";
        $(this).html(question["Kana"] + " / " + kanji + " / " + question["Grp"]).prop("disabled", true);
    });

    $('.tick').mousedown(function (e) {
        is_kanji_corr = true;
        completeQuestion(question, $(this))
    });

    $('.cross').mousedown(function (e) {
        is_kanji_corr = false;
        completeQuestion(question, $(this))
    })
}

function completeQuestion(question, btn) {
    if (is_kana_corr && is_kanji_corr && is_grp_corr) {
        if (question_num < questions.length - 1) {
            incrQuestion(question, is_kana_corr, is_kanji_corr, is_grp_corr);
            question_num++;
            resetFields();
            loadQuestion();
        } else {
            btn.html('<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>').prop("disabled", true);
            incrQuestion(question, is_kana_corr, is_kanji_corr, is_grp_corr, finishAndHome);
        }
    } else {
        incrQuestion(question, is_kana_corr, is_kanji_corr, is_grp_corr);
        questions_wrong.push(question);
        let num_q_remain = questions.length - question_num;
        let swap_q_no = Math.floor(Math.random() * num_q_remain) + question_num;
        [questions[swap_q_no], questions[question_num]] = [questions[question_num], questions[swap_q_no]]
        resetFields();
        loadQuestion();
    }
}

function incrQuestion(question, is_kana_corr, is_kanji_corr, is_grp_corr, callback = function(){}) {
    $.post("vocab/incr/" + question["Row"].toString(),
        JSON.stringify({"IsKanaCorr": is_kana_corr, "IsKanjiCorr": is_kanji_corr, "IsGrpCorr": is_grp_corr}))
        .done(function (data) {
            console.log("incr question: " + data);
            callback()
        })
}

function resetFields() {
    $('.tick').prop("disabled", true);
    $('.cross').prop("disabled", true);
    $('.show').prop("disabled", true);
    $('.show').html("Show");
    $('.submit').prop("disabled", false);
    $('.clear').prop("disabled", false);
    $('.q-left').html("Left (" + (questions.length - question_num).toString() + ")")
    $('.kana-text').removeClass("is-valid");
    $('.kana-text').removeClass("is-invalid");
    $('.kana-text').prop("readonly", false);
    $('.kana-text').val("");
    $('.grp-select').removeClass("is-valid");
    $('.grp-select').removeClass("is-invalid");
    $('.grp-select').val("");
    $('.grp-options').prop("disabled", false);
    $('#kanjiDrawClear').click();

    $('.tick').unbind();
    $('.cross').unbind();
    $('.show').unbind();
    $('.submit').unbind();
}

function finishAndHome() {
    $.post("vocab/check", function (data) {
        document.cookie = "questions_wrong=" + JSON.stringify(questions_wrong) + "; path=/";
        location.href = '/'
    })
}
