let base = Promt();

document.getElementById("click").addEventListener("click", function () {
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
    base.custom({ msg: html, title: "Input the timeline" })
})

function Promt() {
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
        custom: custom
    }
}
