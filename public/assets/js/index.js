import { Config } from "./config.js";

const toastAlert= (message) => Toastify({
    text: message,
    duration: 3000,
    newWindow: false,
    gravity: "top", // `top` or `bottom`
    position: "center", // `left`, `center` or `right`
    stopOnFocus: true, // Prevents dismissing of toast on hover
    style: {
    background: "#fee2e2",
    color: "#dc2626",
    border: "solid red 1px"
    },
    onClick: function(){} // Callback after click
}).showToast()


// htmx.logger = function(elt, event, data) {
//     if(console) {
//         console.log(event, elt, data);
//     }
// }
document.body.addEventListener('htmx:configRequest', function(evt) {
    evt.detail.path = `${Config.apiUrl}${event.detail.path}`
});

htmx.on("htmx:responseError",( event)=> {
    toastAlert(event.detail.xhr.response)
})