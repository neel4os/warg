import { type QVueGlobals } from "quasar";

export function showNotification(q: QVueGlobals,  
    message: string, 
    caption: string,
    type: string
) {
    q.notify({
        message: message,
        caption: caption,
        position: 'top-right',
        timeout: 2000,
        type: type
    });
}
