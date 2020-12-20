console.log("woot crdt")

// Character Deleted
$('#text_box').on('keydown', function (event) {
    // If Backspace/Delete 
    if (event.which === 8) {
        // Technically returns last character in the string
        console.log("string %s deleted %s", $(this).val(), $(this).val().substring($(this).val().length - 1, $(this).val().length));
    }
});

// Character Added
$('#text_box').on('keypress', function (event) {
    console.log("string %s added %s", $(this).val(), String.fromCharCode(event.which));
});