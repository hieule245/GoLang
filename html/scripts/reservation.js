let base = Promt();

// let Change = document.getElementById("click")
// if(Change.classList.contains("text-danger")){
//     Change.classList.remove("text-danger")
// } else {
//     Change.classList.add("text-danger")
// }

document.getElementById("click").addEventListener("click", function () {
    // notify("This is my message", "warning")
    // swalFire("You do it successfully!", "Do you want to continue?", "success", "Next")
    let html = `
    <form id="check-availbility-form" class="need-validation" action="" method="post" novalidate>
                    <div class="row" id="date-modal">
                        <div class="col">
                                <input type="text" class="form-control" id="start" name="start" required
                                    placeholder="YYYY-MM-DD" disabled>
                        </div>
                        <div class="col">
                                <input type="text" class="form-control" id="end" name="end" required
                                    placeholder="YYYY-MM-DD" disabled>
                        </div>
                    </div>
                </form>
    `
    base.custom({ title: "Input the timeline", msg: html })
})


//     (function () {
//         'use strict';
//         window.addEventListener('load', function () {
//             // Fetch all the form
//             var forms = this.document.getElementsByClassName("need-validation");
//             var validation = Array.prototype.filter.call(forms, function (form) {
//                 form.addEventListener('submit', function (event) {
//                     if (form.checkValidity() === false) {
//                         event.preventDefault();
//                         event.stopPropagation();
//                     }
//                     form.classList.add('was-validated');
//                 }, false);
//             });
//         }, false);
//     })();

// const elem = document.getElementById('foo');
// const rangepicker = new DateRangePicker(elem, {
//     format: "yyyy-mm-dd",
// });

// function notify(noti, notiType) {
//     notie.alert({
//         type: notiType,
//         text: noti,
//     })
// }

// function swalFire(title, text, icon, confirmText) {
//     Swal.fire({
//         title: title,
//         html: text,
//         icon: icon,
//         confirmButtonText: confirmText
//     })
// }

function Promt() {

    // let error = function (c) {
    //     const {
    //         msg = "",
    //         icon = "",
    //         position = ""
    //     } = c
    //     Swal.fire({
    //         position: position,
    //         icon: icon,
    //         title: msg,
    //         showConfirmButton: false,
    //         timer: 1500
    //     });
    // }

    // let success = function (c) {
    //     const {
    //         msg = "",
    //         icon = "",
    //         position = ""
    //     } = c
    //     Swal.fire({
    //         position: position,
    //         icon: icon,
    //         title: msg,
    //         showConfirmButton: false,
    //         timer: 1500
    //     });
    // }

    // let toast = function (c) {
    //     const {
    //         msg = "",
    //         icon = "success",
    //         position = "top-end",
    //     } = c
    //     const Toast = Swal.mixin({
    //         toast: true,
    //         position: position,
    //         showConfirmButton: false,
    //         timer: 3000,
    //         timerProgressBar: true,
    //         didOpen: (toast) => {
    //             toast.onmouseenter = Swal.stopTimer;
    //             toast.onmouseleave = Swal.resumeTimer;
    //         }
    //     });
    //     Toast.fire({
    //         icon: icon,
    //         title: msg
    //     });
    // }

    async function custom(c) {
        const {
            title = "",
            msg = "",
            idElement1 = "",
            idElement2 = ""
        } = c
        const { value: formValues } = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                const elem = document.getElementById("date-modal")
                const rp = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                    showOnFocus: true,
                    autoHide: false,
                })
            },
            preConfirm: () => {
                return [
                    document.getElementById("start").value,
                    document.getElementById("end").value
                ]
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
                // Đưa DatePicker vào modal của Swal để tránh bị che
                setTimeout(() => {
                    document.querySelectorAll('.datepicker-dropdown').forEach(dp => {
                        document.querySelector('.swal2-popup').appendChild(dp);
                    });
                }, 0);
            }
        }
        );
        if (formValues) {
            Swal.fire(JSON.stringify(formValues));
        }
    }

    return {
        // success: success,
        // toast: toast,
        // error: error,
        custom: custom
    }
}

