import notie, {AlertType} from "notie";
import Swal, {SweetAlertIcon} from "sweetalert2";
import {SweetAlertResult} from "sweetalert2";

export const notify = (msg:string, msgType:string) => {
    notie.alert({
        type: msgType as AlertType|undefined, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,
    })
}
export const notify_modal = (icon:string, title:string, text:string, footer:string): Promise<SweetAlertResult<any>> => {
    return Swal.fire({
        icon: icon as SweetAlertIcon,
        title: title,
        text: text,
        footer: footer
    })
}
