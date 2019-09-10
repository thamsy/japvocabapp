$(document).ready(function () {
    let login_btn = $('#login-submit');
    login_btn.unbind();
    login_btn.mousedown(function (e) {
        $.post("/login", $("#login-form").serialize())
            .done(function (data) {
                window.location.href = '/'
            })
            .fail(function (data) {
                $('#wrong-pw-alert').html(
                    '<div class="alert alert-danger alert-dismissible fade show" role="alert">\n' +
                    '  Invalid Credentials\n' +
                    '  <button type="button" class="close" data-dismiss="alert" aria-label="Close">\n' +
                    '      <span aria-hidden="true">&times;</span>\n' +
                    '  </button>\n' +
                    '</div>'
                )
        })
    })
});