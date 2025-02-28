let Change = document.getElementById("changePrint").classList

document.getElementById("click").addEventListener("click", function () {
    if (Change.contains("text-danger")) {
        Change.remove("text-danger")
    } else {
        Change.add("text-danger")
    }
})