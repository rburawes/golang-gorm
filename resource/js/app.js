/**
 * Get all the authors and populate author drop down menu.
 */
$(document).ready(function () {
    $('#book').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget) // Button that triggered the modal
        var authorIds = button.data('authorid')
        $.ajax({
            type: "get",
            url: "/authors",
            dataType: "json",
            cache: false
        }).done(function (response) {
            if (response != "") {
                var len = response.length;
                $("#author").empty();
                var authors = '';
                for (var i = 0; i < len; i++) {
                    var id = response[i]['id'];
                    var fname = response[i]['firstname'];
                    var lname = response[i]['lastname'];
                    authors += "<option value='" + id + "'>" + fname + " " + lname + "</option>";
                }
                $("#author").append(authors);
                if (authorIds) {
                    for (var i = 0; i < authorIds.length; i++) {
                        $("#author option[value=" + authorIds[i] + "]").attr('selected', 'selected');
                    }
                }
                $("#author").multiselect('rebuild');
            }
            else {
                console.log("error");
            }
        });
    });
});

/**
 * If add book button is clicked.
 */
$(document).ready(function () {
    $('#addBtn').on('click', function (event) {
        $('#bookForm').attr('action', '/books/create/process');
        $("#isbn").val('');
        $("#title").val('');
        $("#price").val('');
    });
});

/**
 * When update book is clicked.
 * @param isbn
 */
function updateBook(isbn) {
    event.preventDefault();
    $('#bookForm').attr('action', '/books/update/process');
    $.ajax({
        type: "get",
        url: "/books/update?isbn=" + isbn,
        dataType: "json",
        cache: false
    }).done(function (data) {
        if (data != "") {
            $("#isbn").val(data.isbn);
            $("#title").val(data.title);
            $("#price").val(data.price);
        }
        else {
            console.log("error");
        }
    });
}

/**
 * Confirms delete event from a button or an anchor.
 * TODO: Not the best practice to delete using GET method, should be ended in POST/DELETE.
 */
$(document).ready(function () {
    $('#confirm-delete').on('show.bs.modal', function (event) {
        $(this).find('.btn-ok').attr('href', $(event.relatedTarget).data('href'));
        $('.debug-url').html('Delete URL: <strong>' + $(this).find('.btn-ok').attr('href') + '</strong>');
    });
});

/**
 * Submits a form.
 *
 */
$(document).ready(function () {
    $("form").submit(function (event) {
        event.preventDefault();
        var post_url = $(this).attr("action");
        var form_data = $(this).serialize();
        $.post(post_url, form_data).done(function (data) {
            if ((post_url.indexOf("/login") === -1) && (post_url.indexOf("/signup") === -1)) {
                $("#book").prop('disabled', true);
                $('#successContent').html(data.title);
                $('#successPopup').modal('show')
            } else {
                window.location.href = "/";
            }
        }).fail(function (data) {
            $("#error-msg").html("<span><strong>ERROR:</strong> " + data.responseText + "</span>");
            $("#error-msg").show();
        }).always(function () {
            console.log("Form submission ended");
        });
    });

    $("#successPopup").on("hidden.bs.modal", function () {
        $("#book").prop('disabled', false);
        window.location = "/books";
    });
});

/**
 * Validates signup form.
 */
$(document).ready(function () {
    $flag = 1;
    $("#firstname").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_firstname").text("* First name is required.");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_lastname").text("");

        }
    });
    $("#lastname").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_lastname").text("* Last name is required!");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_lastname").text("");
        }
    });
    $("#email").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_email").text("* Email address is required!");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_email").text("");
        }
    });
    $("#password").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_password").text("* Please provide a valid password!");
        }
        else {
            $(this).css({"border-color": "#2eb82e"});
            $('#submit').attr('disabled', false);
            $("#error_password").text("");

        }
    });
    $("#cpassword").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_confirm_password").text("* Confirm your password please!");
        } else {
            $(this).css({"border-color": "#2eb82e"});
            $('#submit').attr('disabled', false);
            $("#error_confirm_password").text("");
        }

    });

    $("#submit").click(function () {
        if ($("#firstname").val() == '') {
            $("#firstname").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_firstname").text("* First name is required!");
        }
        if ($("#lastname").val() == '') {
            $("#lastname").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_lastname").text("* Last name is required!");
        }
        if ($("#email").val() == '') {
            $("#email").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_email").text("* Email address is required!");
        }
        if ($("#password").val() == '') {
            $("#password").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_password").text("* Please provide a valid password!");
        }
        if ($("#cpassword").val() == '') {
            $("#cpassword").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_confirm_password").text("* Confirm your password please!");
        }
    });
});